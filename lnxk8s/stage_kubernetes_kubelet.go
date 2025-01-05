package lnxk8s

import (
	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

var V_stage_kubernetes_kubelet string = `
#!/bin/bash
set -e
set -o pipefail
set -u
cd "$(dirname "$0")"
date
mkdir -p artifact
cd artifact
cat <<\EOF >install_kubernetes_kubelet.sh
#!/bin/bash
set -e
set -o pipefail
set -u
cd "$(dirname "$0")"
date
mkdir -p /opt/kubernetes/kubelet
mkdir -p /opt/kubernetes/manifests
cp -a kubernetes/ca.pem /opt/kubernetes/ca.pem
cp -a kubernetes/kubelet/kubelet                      /opt/kubernetes/kubelet/kubelet
cp -a kubernetes/kubelet/kubelet-bootstrap.kubeconfig /opt/kubernetes/kubelet/kubelet-bootstrap.kubeconfig
cp -a kubernetes/kubelet/kubelet-config.yml           /opt/kubernetes/kubelet/kubelet-config.yml
cp -a kubernetes/kubelet/kubelet.conf                 /opt/kubernetes/kubelet/kubelet.conf
cp -a kubernetes/kubelet/kubelet.service              /usr/lib/systemd/system/kubelet.service
systemctl daemon-reload
systemctl enable kubelet
systemctl restart kubelet || true
date
EOF
date
`

func X_stage_kubernetes_kubelet(args ...string) {
	log.Println(util.GetCurrentFuncName())

	log.Printf("args: %+v\n", args)

	var script string
	script = strings.TrimSpace(V_stage_kubernetes_kubelet)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
