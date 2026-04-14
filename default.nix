{pkgs ? import <nixpkgs> {}, ...}:
let
  version = "0.23.0";
in
pkgs.buildGoModule {
  pname = "golazo";
  inherit version;
  vendorHash = "sha256-M2gfqU5rOfuiVSZnH/Dr8OVmDhyU2jYkgW7RuIUTd+E=";

  ldflags = [
    "-s" "-w"
    "-X github.com/0xjuanma/golazo/cmd.Version=v${version}"
  ];

  subPackages = ["."];

  src = builtins.path {
    path = ./.;
    name = "source";
  };
}
