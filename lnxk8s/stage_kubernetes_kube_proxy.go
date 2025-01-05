package lnxk8s

import (
	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

var V_stage_kubernetes_kube_proxy string = `
#!/bin/bash
set -e
set -o pipefail
set -u
cd "$(dirname "$0")"
date
mkdir -p artifact
cd artifact
cat <<\EOF >install_kubernetes_kube_proxy.sh
#!/bin/bash
set -e
set -o pipefail
set -u
cd "$(dirname "$0")"
date
mkdir -p /opt/kubernetes/kube-proxy
cp -a kubernetes/kube-proxy/kube-proxy            /opt/kubernetes/kube-proxy/kube-proxy
cp -a kubernetes/kube-proxy/kube-proxy-config.yml /opt/kubernetes/kube-proxy/kube-proxy-config.yml
cp -a kubernetes/kube-proxy/kube-proxy.conf       /opt/kubernetes/kube-proxy/kube-proxy.conf
cp -a kubernetes/kube-proxy/kube-proxy.kubeconfig /opt/kubernetes/kube-proxy/kube-proxy.kubeconfig
cp -a kubernetes/kube-proxy/kube-proxy.service    /usr/lib/systemd/system/kube-proxy.service
systemctl daemon-reload
systemctl enable kube-proxy
systemctl restart kube-proxy
date
EOF
date
`

func X_stage_kubernetes_kube_proxy(args ...string) {
	log.Println(util.GetCurrentFuncName())

	log.Printf("args: %+v\n", args)

	var script string
	script = strings.TrimSpace(V_stage_kubernetes_kube_proxy)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
