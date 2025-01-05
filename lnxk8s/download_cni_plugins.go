package lnxk8s

import (
	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

var V_download_cni_plugins string = `
#!/bin/bash
set -e
set -o pipefail
set -u
cd "$(dirname "$0")"
date
source ./env.sh
echo "${NETWORK_REGION}"
mkdir -p pkg
cd pkg
[ "${NETWORK_REGION}" = "us" ] && {
wget -c "https://github.com/containernetworking/plugins/releases/download/v1.2.0/cni-plugins-linux-amd64-v1.2.0.tgz" -O cni-plugins-linux-amd64-v1.2.0.tgz
}
[ "${NETWORK_REGION}" = "cn" ] && {
wget -c "http://199.115.230.237:12345/kubernetes/pkg/cni-plugins-linux-amd64-v1.2.0.tgz" -O cni-plugins-linux-amd64-v1.2.0.tgz
}
date
`

func X_download_cni_plugins(args ...string) {
	log.Println(util.GetCurrentFuncName())

	log.Printf("args: %+v\n", args)

	var script string
	script = strings.TrimSpace(V_download_cni_plugins)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
