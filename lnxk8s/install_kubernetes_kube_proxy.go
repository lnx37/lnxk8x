package lnxk8s

import (
	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

var V_install_kubernetes_kube_proxy string = `
#!/bin/bash
set -e
set -o pipefail
set -u
cd "$(dirname "$0")"
date
[ "$#" -ne 1 ] && echo "invalid argument, need an ip" && exit 1
WORKER_IP="$1"
echo "${WORKER_IP}"
ssh root@"${WORKER_IP}" "bash /opt/artifact/install_kubernetes_kube_proxy.sh"
date
`

func X_install_kubernetes_kube_proxy(args ...string) {
	log.Println(util.GetCurrentFuncName())

	log.Printf("args: %+v\n", args)

	var script string
	script = strings.TrimSpace(V_install_kubernetes_kube_proxy)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
