package lnxk8s

import (
	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

var V_unpkg_etcd string = `
#!/bin/bash
set -e
set -o pipefail
set -u
cd "$(dirname "$0")"
date
mkdir -p pkg
mkdir -p artifact/etcd
cd pkg
[ -d etcd-v3.5.12-linux-amd64 ] && rm -rf etcd-v3.5.12-linux-amd64
tar xzf etcd-v3.5.12-linux-amd64.tar.gz
chown -R root:root etcd-v3.5.12-linux-amd64
chmod +x etcd-v3.5.12-linux-amd64/etcd
chmod +x etcd-v3.5.12-linux-amd64/etcdctl
cp -a etcd-v3.5.12-linux-amd64/etcd    ../artifact/etcd/etcd
cp -a etcd-v3.5.12-linux-amd64/etcdctl ../artifact/etcd/etcdctl
date
`

func X_unpkg_etcd(args ...string) {
	log.Println(util.GetCurrentFuncName())

	log.Printf("args: %+v\n", args)

	var script string
	script = strings.TrimSpace(V_unpkg_etcd)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
