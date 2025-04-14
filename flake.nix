{
  description = "Development flake for Cocopasty";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-24.11";
  };

  outputs =
    { nixpkgs, ... }:
    let
      x86 = "x86_64-linux";
      pkgs = nixpkgs.legacyPackages."${x86}";
    in
    {
      devShells."${x86}".default = pkgs.mkShellNoCC {
        packages = with pkgs; [
          # Golang
          go
          golangci-lint

          # Formatters
          treefmt
          beautysh
          mdformat
          yamlfmt
          deadnix
          jsonfmt
          hadolint
          nixfmt-rfc-style

          # Tools
          curl
        ];

        shellHook = ''
          git config --local core.hooksPath .githooks/
        '';

        # Environment Variables
        COCOPASTY_DATA_LOCATION = "./value";
        COCOPASTY_PERSIST_TO_DISK = "true";
        COCOPASTY_PORT = ":8080";
      };
    };
}
