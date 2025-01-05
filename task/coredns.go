package task

import (
	_ "embed"

	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

//go:embed yml/coredns.yaml.base
var V_coredns string

var V_install_coredns string = `
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

func X_install_coredns(args ...string) {
	log.Println(util.GetCurrentFuncName())

	V_install_coredns = fmt.Sprintf(V_install_coredns, V_coredns)

	var script string
	script = strings.TrimSpace(V_install_coredns)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
