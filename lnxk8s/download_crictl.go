package lnxk8s

import (
	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

var V_download_crictl string = `
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
wget -c "https://github.com/kubernetes-sigs/cri-tools/releases/download/v1.28.0/crictl-v1.28.0-linux-amd64.tar.gz" -O crictl-v1.28.0-linux-amd64.tar.gz
}
[ "${NETWORK_REGION}" = "cn" ] && {
wget -c "http://199.115.230.237:12345/kubernetes/pkg/crictl-v1.28.0-linux-amd64.tar.gz" -O crictl-v1.28.0-linux-amd64.tar.gz
}
date
`

func X_download_crictl(args ...string) {
	log.Println(util.GetCurrentFuncName())

	log.Printf("args: %+v\n", args)

	var script string
	script = strings.TrimSpace(V_download_crictl)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
