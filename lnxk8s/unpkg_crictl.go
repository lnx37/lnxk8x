package lnxk8s

import (
	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

var V_unpkg_crictl string = `
#!/bin/bash
set -e
set -o pipefail
set -u
cd "$(dirname "$0")"
date
mkdir -p pkg
mkdir -p artifact/crictl
cd pkg
[ -d crictl ] && rm -rf crictl
mkdir -p crictl
tar xzf crictl-v1.28.0-linux-amd64.tar.gz -C crictl
chown -R root:root crictl
chmod +x crictl/crictl
cp -a crictl/crictl ../artifact/crictl/crictl
date
`

func X_unpkg_crictl(args ...string) {
	log.Println(util.GetCurrentFuncName())

	log.Printf("args: %+v\n", args)

	var script string
	script = strings.TrimSpace(V_unpkg_crictl)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
