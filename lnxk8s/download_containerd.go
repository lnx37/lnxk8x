package lnxk8s

import (
	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

var V_download_containerd string = `
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
wget -c "https://github.com/containerd/containerd/releases/download/v1.6.28/containerd-1.6.28-linux-amd64.tar.gz" -O containerd-1.6.28-linux-amd64.tar.gz
}
[ "${NETWORK_REGION}" = "cn" ] && {
wget -c "http://199.115.230.237:12345/kubernetes/pkg/containerd-1.6.28-linux-amd64.tar.gz" -O containerd-1.6.28-linux-amd64.tar.gz
}
date
`

func X_download_containerd(args ...string) {
	log.Println(util.GetCurrentFuncName())

	log.Printf("args: %+v\n", args)

	var script string
	script = strings.TrimSpace(V_download_containerd)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
