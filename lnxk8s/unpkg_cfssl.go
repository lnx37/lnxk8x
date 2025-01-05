package lnxk8s

import (
	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

var V_unpkg_cfssl string = `
#!/bin/bash
set -e
set -o pipefail
set -u
cd "$(dirname "$0")"
date
mkdir -p pkg
mkdir -p artifact/cfssl
cd pkg
chown root:root cfssl-certinfo_1.6.4_linux_amd64
chown root:root cfssl_1.6.4_linux_amd64
chown root:root cfssljson_1.6.4_linux_amd64
chmod +x cfssl-certinfo_1.6.4_linux_amd64
chmod +x cfssl_1.6.4_linux_amd64
chmod +x cfssljson_1.6.4_linux_amd64
cp -a cfssl-certinfo_1.6.4_linux_amd64 ../artifact/cfssl/cfssl-certinfo
cp -a cfssl_1.6.4_linux_amd64          ../artifact/cfssl/cfssl
cp -a cfssljson_1.6.4_linux_amd64      ../artifact/cfssl/cfssljson
cp -a cfssl-certinfo_1.6.4_linux_amd64 /usr/local/bin/cfssl-certinfo
cp -a cfssl_1.6.4_linux_amd64          /usr/local/bin/cfssl
cp -a cfssljson_1.6.4_linux_amd64      /usr/local/bin/cfssljson
date
`

func X_unpkg_cfssl(args ...string) {
	log.Println(util.GetCurrentFuncName())

	log.Printf("args: %+v\n", args)

	var script string
	script = strings.TrimSpace(V_unpkg_cfssl)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
