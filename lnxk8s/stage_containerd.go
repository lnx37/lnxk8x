package lnxk8s

import (
	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

var V_stage_containerd string = `
#!/bin/bash
set -e
set -o pipefail
set -u
cd "$(dirname "$0")"
date
mkdir -p artifact
cd artifact
cat <<\EOF >install_containerd.sh
#!/bin/bash
set -e
set -o pipefail
set -u
cd "$(dirname "$0")"
date
mkdir -p /etc/containerd
cp -a containerd/config.toml             /etc/containerd/config.toml
cp -a containerd/containerd              /usr/local/bin/containerd
cp -a containerd/containerd-shim         /usr/local/bin/containerd-shim
cp -a containerd/containerd-shim-runc-v1 /usr/local/bin/containerd-shim-runc-v1
cp -a containerd/containerd-shim-runc-v2 /usr/local/bin/containerd-shim-runc-v2
cp -a containerd/containerd-stress       /usr/local/bin/containerd-stress
cp -a containerd/containerd.service      /usr/lib/systemd/system/containerd.service
cp -a containerd/ctr                     /usr/local/bin/ctr
systemctl daemon-reload
systemctl enable containerd
systemctl restart containerd
date
EOF
date
`

func X_stage_containerd(args ...string) {
	log.Println(util.GetCurrentFuncName())

	log.Printf("args: %+v\n", args)

	var script string
	script = strings.TrimSpace(V_stage_containerd)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
