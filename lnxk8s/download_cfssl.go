package lnxk8s

import (
	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

var V_download_cfssl string = `
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
wget -c "https://github.com/cloudflare/cfssl/releases/download/v1.6.4/cfssl-certinfo_1.6.4_linux_amd64" -O cfssl-certinfo_1.6.4_linux_amd64
wget -c "https://github.com/cloudflare/cfssl/releases/download/v1.6.4/cfssl_1.6.4_linux_amd64"          -O cfssl_1.6.4_linux_amd64
wget -c "https://github.com/cloudflare/cfssl/releases/download/v1.6.4/cfssljson_1.6.4_linux_amd64"      -O cfssljson_1.6.4_linux_amd64
}
[ "${NETWORK_REGION}" = "cn" ] && {
wget -c "http://199.115.230.237:12345/kubernetes/pkg/cfssl-certinfo_1.6.4_linux_amd64" -O cfssl-certinfo_1.6.4_linux_amd64
wget -c "http://199.115.230.237:12345/kubernetes/pkg/cfssl_1.6.4_linux_amd64"          -O cfssl_1.6.4_linux_amd64
wget -c "http://199.115.230.237:12345/kubernetes/pkg/cfssljson_1.6.4_linux_amd64"      -O cfssljson_1.6.4_linux_amd64
}
date
`

func X_download_cfssl(args ...string) {
	log.Println(util.GetCurrentFuncName())

	log.Printf("args: %+v\n", args)

	var script string
	script = strings.TrimSpace(V_download_cfssl)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
