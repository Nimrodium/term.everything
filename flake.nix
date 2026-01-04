{
  description = "Run any GUI app in the terminal";
  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  outputs =
    {
      self,
      nixpkgs,
      ...
    }:
    let
      allSystems = [
        "x86_64-linux" # 64-bit Intel/AMD Linux
        "aarch64-linux" # 64-bit ARM Linux
        # "x86_64-darwin" # 64-bit Intel macOS
        # "aarch64-darwin" # 64-bit ARM macOS
      ];
      forAllSystems =
        f:
        nixpkgs.lib.genAttrs allSystems (
          system:
          f {
            pkgs = import nixpkgs { inherit system; };
          }
        );
    in
    {
      packages = forAllSystems (
        { pkgs }:
        {
          default = pkgs.buildGoModule rec {
            pname = "term-everything";
            name = pname;
            version = "0.7.8";
            subPackages = [ "." ];
            src = ./.;
            vendorHash = null;
            nativeBuildInputs = with pkgs; [
              pkg-config
            ];
            buildInputs = with pkgs; [
              glib
              chafa
            ];
            preBuild = ''
              go generate ./wayland
            '';
            postInstall = ''
              # rm $out/bin/generate
              mv $out/bin/term.everything $out/bin/${name}
            '';
          };
        }
      );
    };
}
