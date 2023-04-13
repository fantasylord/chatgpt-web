{ pkgs }: {
    deps = [
        pkgs.sudo
        pkgs.systemd
        pkgs.lsof
        pkgs.unixtools.nettools
        pkgs.nodejs-16_x
        pkgs.go
        pkgs.gopls
    ];
}