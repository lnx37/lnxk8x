package lnxk8s

import (
	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

var V_download_runc string = `
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
mkdir -p runc_v1.1.11
wget -c "https://github.com/opencontainers/runc/releases/download/v1.1.11/runc.amd64" -O runc_v1.1.11/runc.amd64
}
[ "${NETWORK_REGION}" = "cn" ] && {
mkdir -p runc_v1.1.11
wget -c "http://199.115.230.237:12345/kubernetes/pkg/runc_v1.1.11/runc.amd64" -O runc_v1.1.11/runc.amd64
}
date
`

func X_download_runc(args ...string) {
	log.Println(util.GetCurrentFuncName())

	log.Printf("args: %+v\n", args)

	var script string
	script = strings.TrimSpace(V_download_runc)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
