{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs";
    flake-utils.url = "github:numtide/flake-utils";
    treefmt-nix.url = "github:numtide/treefmt-nix";
  };

  outputs = { self, nixpkgs, flake-utils, treefmt-nix }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        deps = with pkgs; [ go ];
        scriptDeps = with pkgs; [ nushellFull nodejs_20 ];
        devDeps = with pkgs;
          [
            (pkgs.python312.withPackages (p: [
              p.httpx
              p.icecream
              p.tqdm
              p.uuid
              p.psycopg2
              p.types-psycopg2
              p.boto3
              p.mypy-boto3-secretsmanager
            ]))
            postgresql
            sqlc
            entr
            awscli2
            (pkgs.nodePackages.aws-cdk.overrideAttrs (_: {
              preRebuild = ''
                substituteInPlace lib/index.js \
                --replace 'await fs27.copy(fromFile,toFile)' 'await fs27.copy(fromFile, toFile); await fs27.chmod(toFile, 0o644);'
                tar --to-stdout -xf $src package/package.json \
                | ${pkgs.jq}/bin/jq '{"devDependencies"}' > /build/devDependencies.json
              '';
              postInstall = ''
                FIXED_PACKAGE_JSON="$(${pkgs.jq}/bin/jq -s '.[0] * .[1]' package.json /build/devDependencies.json)"
                printf "%s\n" "$FIXED_PACKAGE_JSON" > package.json
              '';
            }))
          ] ++ scriptDeps;

        sqlformat = {
          language = "postgresql";
          dialect = "postgresql";
          keywordCase = "lower";
          functionCase = "lower";
          dataTypeCase = "lower";
          identifierCase = "lower";
        };

        treefmtEval = treefmt-nix.lib.evalModule pkgs ({ pkgs, ... }: {
          projectRootFile = "flake.nix";
          programs.nixfmt.enable = true;
          programs.gofmt.enable = true;
          programs.yamlfmt.enable = true;
          programs.black.enable = true;
          programs.isort.enable = true;
          programs.prettier.enable = true;
          settings = {
            global.excludes = [ "vendor/*" ];
            formatter.sql-formatter = {

              command = pkgs.lib.getExe pkgs.bash;
              options = [
                "-euc"
                ''
                  for file in "$@"; do
                  ${pkgs.nodePackages.sql-formatter}/bin/sql-formatter --fix --config=${
                    pkgs.writeText "config.json" (builtins.toJSON sqlformat)
                  } $file
                  done
                ''
              ];
              includes = [ "*.sql" ];
            };
          };
        });
      in {
        packages = rec {
          default = pkgs.buildGoModule {
            name = "cart-recommendation-engine";
            src = ./.;
            vendorHash = null;
            doCheck = false;
            preBuild = "${pkgs.sqlc}/bin/sqlc generate";
            tags = [ "ignore-iac" "lambda.norpc" ];
            CGO_ENABLED = 0;
          };

          docker = pkgs.dockerTools.buildImage {
            name = "cart-recommendation-engine-docker";
            tag = "latest";
            copyToRoot = pkgs.buildEnv {
              name = "image-root";
              paths = [ default ];
              pathsToLink = [ "/bin" ];
            };
            config = {
              Cmd = [ "/bin/cart-recommendation-engine" ];
              ExposedPorts = { "8090/tcp" = { }; };
            };
          };

          lambdaZip = pkgs.stdenv.mkDerivation {
            name = "cart-recommendation-engine-lambda";
            src = default;
            phases = [ "installPhase" ];
            installPhase = ''
              mkdir -p $out
              cp $src/bin/cart-recommendation-engine bootstrap
              chmod +x bootstrap
              ${pkgs.zip}/bin/zip -r $out/lambda.zip bootstrap
            '';
          };

          initDB = pkgs.dockerTools.buildImage {
            name = "cart-recommendation-engine-initdb";
            tag = "latest";
            copyToRoot = pkgs.buildEnv {
              name = "image-root";
              paths = scriptDeps ++ [ ./seed ./scripts ];
              pathsToLink = [ "/bin" "/" ];
            };
            config = {
              Cmd = [ "${pkgs.nushell}/bin/nu" ];
              ExposedPorts = { "8090/tcp" = { }; };
            };
          };
        };

        devShell = nixpkgs.legacyPackages.${system}.mkShell rec {
          PGHOST = "127.0.0.1";
          PGPORT = 5432;
          PGDATABASE = "cart-recommendation-engine";
          PGUSER = "postgres";
          PGPASSWORD = "hunter2";

          NEO4J_USER = "neo4j";
          NEO4J_PASSWORD = "supersecretpassword";
          NEO4J_URI = "neo4j://localhost";

          NEO4J_AUTH = "${NEO4J_USER}/${NEO4J_PASSWORD}";

          PYTHONPATH = "${pkgs.python312}/bin/python";

          buildInputs = deps ++ devDeps;
        };

        formatter = treefmtEval.config.build.wrapper;
      });
}
