{
  description = "Nix flake for gf-cli";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  };

  outputs =
    { self, nixpkgs, ... }:
    let
      forAllSystems = nixpkgs.lib.genAttrs nixpkgs.lib.systems.flakeExposed;
      pkgsFor = system: import nixpkgs { inherit system; };
    in
    {
      packages = forAllSystems (
        system:
        let
          pkgs = pkgsFor system;
          lib = pkgs.lib;
        in
        {
          default = pkgs.buildGoModule rec {
            name = "gf-cli";
            version = "${self.rev or "(devel)"}";
            subPackages = [ "cmd/gf" ];
            src = ./.;
            vendorHash = "sha256-xOoujvNhBYLgTkpdML6bxQhNsRVOINbpNXIAEAqLsl0=";

            nativeBuildInputs = [ pkgs.installShellFiles ];
            ldflags = [
              "-s"
              "-w"
              "-X main.version=${version}"
            ];
            postInstall = ''
              installManPage man/gf*
              installShellCompletion --cmd gf \
                --bash <($out/bin/gf completion bash) \
                --zsh <($out/bin/gf completion zsh) \
                --fish <($out/bin/gf completion fish) \
            '';

            nativeInstallCheckInputs = [ pkgs.versionCheckHook ];
            doInstallCheck = true;

            meta = {
              description = "Grafana API Client for command line operations";
              homepage = "https://github.com/johejo/gf-cli";
              license = lib.licenses.apsl20;
              mainProgram = "gf";
            };
          };
        }
      );
      devShells = forAllSystems (
        system:
        let
          pkgs = pkgsFor system;
        in
        {
          default = pkgs.mkShellNoCC {
            packages = with pkgs; [
              go
              git
              podman
              podman-compose
              nixfmt
            ];
          };
        }
      );
    };
}
