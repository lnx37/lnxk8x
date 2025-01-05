package lnxk8s

import (
	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

var V_stage_kubernetes_kubectl string = `
#!/bin/bash
set -e
set -o pipefail
set -u
cd "$(dirname "$0")"
date
mkdir -p artifact
cd artifact
cat <<\EOF >install_kubernetes_kubectl.sh
#!/bin/bash
set -e
set -o pipefail
set -u
cd "$(dirname "$0")"
date
mkdir -p /opt/kubernetes/kubectl
mkdir -p /root/.kube
cp -a kubernetes/kubectl/kubectl            /opt/kubernetes/kubectl/kubectl
cp -a kubernetes/kubectl/kubectl.kubeconfig /opt/kubernetes/kubectl/kubectl.kubeconfig
cp -a kubernetes/kubectl/kubectl            /usr/local/bin/kubectl
cp -a kubernetes/kubectl/kubectl.kubeconfig /root/.kube/config
cp -a kubernetes/kubectl/kubectl.kubeconfig /opt/kubernetes/admin.kubeconfig
date
EOF
date
`

func X_stage_kubernetes_kubectl(args ...string) {
	log.Println(util.GetCurrentFuncName())

	log.Printf("args: %+v\n", args)

	var script string
	script = strings.TrimSpace(V_stage_kubernetes_kubectl)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
