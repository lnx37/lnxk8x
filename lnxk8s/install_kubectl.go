package lnxk8s

import (
	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

var V_install_kubectl string = `
#!/bin/bash
set -e
set -o pipefail
set -u
cd "$(dirname "$0")"
date
[ -f /root/.kube/config ] && echo "/root/.kube/config already exists" && echo "override anyway"
[ -d /root/.kube ] && rm -rf /root/.kube
mkdir -p /root/.kube
cp -a artifact/kubernetes/kubectl/kubectl.kubeconfig /root/.kube/config
date
`

func X_install_kubectl(args ...string) {
	log.Println(util.GetCurrentFuncName())

	log.Printf("args: %+v\n", args)

	var script string
	script = strings.TrimSpace(V_install_kubectl)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
