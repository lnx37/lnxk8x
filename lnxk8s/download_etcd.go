package lnxk8s

import (
	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

var V_download_etcd string = `
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
wget -c "https://github.com/etcd-io/etcd/releases/download/v3.5.12/etcd-v3.5.12-linux-amd64.tar.gz" -O etcd-v3.5.12-linux-amd64.tar.gz
}
[ "${NETWORK_REGION}" = "cn" ] && {
wget -c "http://199.115.230.237:12345/kubernetes/pkg/etcd-v3.5.12-linux-amd64.tar.gz" -O etcd-v3.5.12-linux-amd64.tar.gz
}
date
`

func X_download_etcd(args ...string) {
	log.Println(util.GetCurrentFuncName())

	log.Printf("args: %+v\n", args)

	var script string
	script = strings.TrimSpace(V_download_etcd)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
