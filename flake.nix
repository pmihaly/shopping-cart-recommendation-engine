{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs";
    flake-utils.url = "github:numtide/flake-utils";
    pre-commit-hooks.url = "github:cachix/pre-commit-hooks.nix";
    treefmt-nix.url = "github:numtide/treefmt-nix";
  };

  outputs = { self, nixpkgs, flake-utils, pre-commit-hooks, treefmt-nix }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        deps = with pkgs; [ go ];
        scriptDeps = with pkgs; [
          nushellFull
          (pkgs.python312.withPackages
            (p: [ p.httpx p.icecream p.tqdm p.uuid ]))
        ];
        devDeps = with pkgs; [ postgresql sqlc entr ] ++ scriptDeps;

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

        checks = {
          pre-commit-check = pre-commit-hooks.lib.${system}.run {
            src = ./.;
            hooks = {
              nixfmt.enable = true;
              gofmt.enable = true;
            };
          };
          formatting = treefmtEval.config.build.check self;
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

          inherit (self.checks.${system}.pre-commit-check) shellHook;
          buildInputs = deps ++ devDeps;
        };

        formatter = treefmtEval.config.build.wrapper;
      });
}
