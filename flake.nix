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
        devDeps = with pkgs; [ postgresql nushell sqlc ];

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
          settings.formatter.sql-formatter = {
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
        });
      in {
        packages.default = pkgs.buildGoModule {
          name = "shopping-cart-recommendation-engine";
          src = ./.;
          vendorHash = null;
          doCheck = true;
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
          PGDATABASE = "shopping-cart-recommendation-engine";
          PGUSER = "postgres";
          PGPASSWORD = "hunter2";

          POSTGRES_URL = "postgresql://${PGUSER}:${PGPASSWORD}@${PGHOST}:${
              builtins.toString PGPORT
            }/${PGDATABASE}";

          inherit (self.checks.${system}.pre-commit-check) shellHook;
          buildInputs = deps ++ devDeps;
        };

        formatter = treefmtEval.config.build.wrapper;
      });
}
