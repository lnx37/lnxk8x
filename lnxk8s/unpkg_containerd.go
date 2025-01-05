package lnxk8s

import (
	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

var V_unpkg_containerd string = `
#!/bin/bash
set -e
set -o pipefail
set -u
cd "$(dirname "$0")"
date
mkdir -p pkg
mkdir -p artifact/containerd
cd pkg
[ -d containerd ] && rm -rf containerd
mkdir -p containerd
tar xzf containerd-1.6.28-linux-amd64.tar.gz -C containerd
chown -R root:root containerd
chmod +x containerd/bin/containerd
chmod +x containerd/bin/containerd-shim
chmod +x containerd/bin/containerd-shim-runc-v1
chmod +x containerd/bin/containerd-shim-runc-v2
chmod +x containerd/bin/containerd-stress
chmod +x containerd/bin/ctr
cp -a containerd/bin/containerd              ../artifact/containerd/containerd
cp -a containerd/bin/containerd-shim         ../artifact/containerd/containerd-shim
cp -a containerd/bin/containerd-shim-runc-v1 ../artifact/containerd/containerd-shim-runc-v1
cp -a containerd/bin/containerd-shim-runc-v2 ../artifact/containerd/containerd-shim-runc-v2
cp -a containerd/bin/containerd-stress       ../artifact/containerd/containerd-stress
cp -a containerd/bin/ctr                     ../artifact/containerd/ctr
date
`

func X_unpkg_containerd(args ...string) {
	log.Println(util.GetCurrentFuncName())

	log.Printf("args: %+v\n", args)

	var script string
	script = strings.TrimSpace(V_unpkg_containerd)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
