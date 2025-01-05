package task

import (
	_ "embed"

	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

//go:embed yml/calico.yaml
var V_calico string

var V_install_calico string = `
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

func X_install_calico(args ...string) {
	log.Println(util.GetCurrentFuncName())

	V_install_calico = fmt.Sprintf(V_install_calico, V_calico)

	var script string
	script = strings.TrimSpace(V_install_calico)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
