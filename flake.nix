{
  description = "go learning";

  inputs = {nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";};

  outputs = {
    self,
    nixpkgs,
  }: {
    devShells = {
      x86_64-linux.default = nixpkgs.legacyPackages.x86_64-linux.mkShellNoCC {
        buildInputs = with nixpkgs.legacyPackages.x86_64-linux; [go_1_22 gopls revive];

        shellHook = ''
          echo "go code coco!"
        '';
      };
    };
  };
}
