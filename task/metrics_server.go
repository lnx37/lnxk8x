package task

import (
	_ "embed"

	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

//go:embed yml/components.yaml
var V_metrics_server string

var V_install_metrics_server string = `
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

func X_install_metrics_server(args ...string) {
	log.Println(util.GetCurrentFuncName())

	V_install_metrics_server = fmt.Sprintf(V_install_metrics_server, V_metrics_server)

	var script string
	script = strings.TrimSpace(V_install_metrics_server)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
