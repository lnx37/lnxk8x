package lnxk8s

import (
	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

var V_download_kubernetes string = `
#!/bin/bash
set -e
set -o pipefail
set -u
cd "$(dirname "$0")"
date
source ./env.sh
echo "${NETWORK_REGION}"
echo "${KUBERNETES_VERSION}"
mkdir -p pkg
cd pkg
[ "${NETWORK_REGION}" = "us" ] && {
mkdir -p "kubernetes_${KUBERNETES_VERSION}"
wget -c --no-check-certificate "https://dl.k8s.io/${KUBERNETES_VERSION}/kubernetes-server-linux-amd64.tar.gz" -O "kubernetes_${KUBERNETES_VERSION}/kubernetes-server-linux-amd64.tar.gz"
}
[ "${NETWORK_REGION}" = "cn" ] && {
mkdir -p "kubernetes_${KUBERNETES_VERSION}"
wget -c "http://199.115.230.237:12345/kubernetes/pkg/kubernetes_${KUBERNETES_VERSION}/kubernetes-server-linux-amd64.tar.gz" -O "kubernetes_${KUBERNETES_VERSION}/kubernetes-server-linux-amd64.tar.gz"
}
date
`

func X_download_kubernetes(args ...string) {
	log.Println(util.GetCurrentFuncName())

	log.Printf("args: %+v\n", args)

	var script string
	script = strings.TrimSpace(V_download_kubernetes)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
