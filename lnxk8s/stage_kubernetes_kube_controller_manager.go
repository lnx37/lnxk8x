package lnxk8s

import (
	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

var V_stage_kubernetes_kube_controller_manager string = `
#!/bin/bash
set -e
set -o pipefail
set -u
cd "$(dirname "$0")"
date
mkdir -p artifact
cd artifact
cat <<\EOF >install_kubernetes_kube_controller_manager.sh
#!/bin/bash
set -e
set -o pipefail
set -u
cd "$(dirname "$0")"
date
mkdir -p /opt/kubernetes/kube-controller-manager
cp -a kubernetes/ca-key.pem             /opt/kubernetes/ca-key.pem
cp -a kubernetes/ca.pem                 /opt/kubernetes/ca.pem
cp -a kubernetes/front-proxy-ca-key.pem /opt/kubernetes/front-proxy-ca-key.pem
cp -a kubernetes/front-proxy-ca.pem     /opt/kubernetes/front-proxy-ca.pem
cp -a kubernetes/sa.key                 /opt/kubernetes/sa.key
cp -a kubernetes/sa.pub                 /opt/kubernetes/sa.pub
cp -a kubernetes/kube-controller-manager/kube-controller-manager            /opt/kubernetes/kube-controller-manager/kube-controller-manager
cp -a kubernetes/kube-controller-manager/kube-controller-manager.conf       /opt/kubernetes/kube-controller-manager/kube-controller-manager.conf
cp -a kubernetes/kube-controller-manager/kube-controller-manager.kubeconfig /opt/kubernetes/kube-controller-manager/kube-controller-manager.kubeconfig
cp -a kubernetes/kube-controller-manager/kube-controller-manager.service    /usr/lib/systemd/system/kube-controller-manager.service
systemctl daemon-reload
systemctl enable kube-controller-manager
systemctl restart kube-controller-manager || true
date
EOF
date
`

func X_stage_kubernetes_kube_controller_manager(args ...string) {
	log.Println(util.GetCurrentFuncName())

	log.Printf("args: %+v\n", args)

	var script string
	script = strings.TrimSpace(V_stage_kubernetes_kube_controller_manager)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
