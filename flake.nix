{
  description = "File Integrity Manager - library, tool and server to manage files integrity";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs }:
    let
      systems = [ "x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin" ];
      forAllSystems = nixpkgs.lib.genAttrs systems;
    in
    {
      packages = forAllSystems (system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
        in
        {
          default = self.packages.${system}.fim;

          fim = pkgs.buildGoModule {
            pname = "fim";
            version = "0.0.0";

            src = self;

            vendorHash = "sha256-G77lbbKp1zfgcOP9CHqPA0kuh2jdXoVfAe+oEQjsmQc=";

            nativeBuildInputs = [ pkgs.git ];

            buildPhase = ''
              runHook preBuild
              export HOME=$TMPDIR
              git init -q
              git config user.email "nix@localhost"
              git config user.name "nix"
              git add -A
              git commit -q -m "nix build"
              ./gomake build -L debug
              runHook postBuild
            '';

            installPhase = ''
              runHook preInstall
              install -Dm755 dist/fim-*/fim $out/bin/fim
              runHook postInstall
            '';

            doCheck = false;

            meta = {
              description = "Library, tool and server to manage files integrity";
              homepage = "https://github.com/n0rad/file-integrity-manager";
              mainProgram = "fim";
            };
          };
        });

      nixosModules.default = { lib, pkgs, config, ... }:
        let
          cfg = config.services.file-integrity-manager;
        in
        {
          options.services.file-integrity-manager = {
            enable = lib.mkEnableOption "file-integrity-manager server";

            package = lib.mkOption {
              type = lib.types.package;
              default = self.packages.${pkgs.stdenv.hostPlatform.system}.fim;
              description = "The file-integrity-manager package to use.";
            };
          };

          config = lib.mkIf cfg.enable {
            environment.systemPackages = [ cfg.package ];
          };
        };
    };
}
