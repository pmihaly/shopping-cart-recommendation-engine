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
        scriptDeps = with pkgs; [
          nushellFull
          (pkgs.python312.withPackages
            (p: [ p.httpx p.icecream p.tqdm p.uuid ]))
        ];
        devDeps = with pkgs;
          [
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
        packages.default = pkgs.buildGoModule {
          name = "cart-recommendation-engine";
          src = ./.;
          vendorHash = null;
          doCheck = true;
          preBuild = "${pkgs.sqlc}/bin/sqlc generate";
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
          POSTGRES_URL = "postgresql://${PGUSER}:${PGPASSWORD}@${PGHOST}:${
              builtins.toString PGPORT
            }/${PGDATABASE}";

          buildInputs = deps ++ devDeps;
        };

        formatter = treefmtEval.config.build.wrapper;
      });
}
