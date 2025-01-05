package task

import (
	_ "embed"

	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

//go:embed yml/kube-flannel.yml
var V_flannel string

var V_install_flannel string = `
#!/bin/bash

set -e
set -o pipefail
set -u

date

cat <<EOF |kubectl apply -f -
%s
EOF

date
`

func X_install_flannel(args ...string) {
	log.Println(util.GetCurrentFuncName())

	V_install_flannel = fmt.Sprintf(V_install_flannel, V_flannel)

	var script string
	script = strings.TrimSpace(V_install_flannel)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
