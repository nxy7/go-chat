{
  description = "Project starter";
  inputs = {
    nixpkgs.url = "nixpkgs/nixpkgs-unstable";
    telepresence-nix.url = "github:nxy7/telepresence-nix";
    flake-parts.url = "github:hercules-ci/flake-parts";
    nix2container.url = "github:nlewo/nix2container";
  };

  outputs = { flake-parts, nixpkgs, ... }@inputs:
    flake-parts.lib.mkFlake { inherit inputs; } {
      systems = [ "x86_64-linux" ];
      perSystem = { config, system, ... }:
        let
          pkgs = import nixpkgs { inherit system; };
          telepresencePkg =
            inputs.telepresence-nix.outputs.defaultPackage.${system};
        in {
          devShells.default = pkgs.mkShell {
            packages = with pkgs; [
              go
              telepresencePkg
              nodejs
              bun
              kubernetes-helm
              nodePackages.pnpm
            ];
          };
        };
    };
}
