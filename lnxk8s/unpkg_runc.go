package lnxk8s

import (
	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

var V_unpkg_runc string = `
#!/bin/bash
set -e
set -o pipefail
set -u
cd "$(dirname "$0")"
date
mkdir -p pkg
mkdir -p artifact/runc
cd pkg
chown -R root:root runc_v1.1.11
chmod +x runc_v1.1.11/runc.amd64
cp -a runc_v1.1.11/runc.amd64 ../artifact/runc/runc
date
`

func X_unpkg_runc(args ...string) {
	log.Println(util.GetCurrentFuncName())

	log.Printf("args: %+v\n", args)

	var script string
	script = strings.TrimSpace(V_unpkg_runc)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
