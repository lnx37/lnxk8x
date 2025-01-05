package lnxk8s

import (
	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

var V_install_etcd string = `
#!/bin/bash
set -e
set -o pipefail
set -u
cd "$(dirname "$0")"
date
[ "$#" -ne 1 ] && echo "invalid argument, need an ip" && exit 1
ETCD_IP="$1"
echo "${ETCD_IP}"
ETCD_NAME="$(echo "${ETCD_IP}" |sed "s/\./_/g")"
echo "${ETCD_NAME}"
ssh root@"${ETCD_IP}" "bash /opt/artifact/install_etcd_${ETCD_NAME}.sh"
date
`

func X_install_etcd(args ...string) {
	log.Println(util.GetCurrentFuncName())

	log.Printf("args: %+v\n", args)

	var script string
	script = strings.TrimSpace(V_install_etcd)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
