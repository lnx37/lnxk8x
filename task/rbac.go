package task

import (
	_ "embed"

	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

//go:embed yml/kubelet_tls_bootstrapping.yml
var V_kubelet_tls_bootstrapping string

//go:embed yml/kube_apiserver_to_kubelet.yml
var V_kube_apiserver_to_kubelet string

var V_setup string = `
#!/bin/bash

set -e
set -o pipefail
set -u

date

cat <<EOF |kubectl apply -f -
%s
EOF

cat <<EOF |kubectl apply -f -
%s
EOF

date
`

func X_setup(args ...string) {
	log.Println(util.GetCurrentFuncName())

	V_setup = fmt.Sprintf(V_setup, V_kubelet_tls_bootstrapping, V_kube_apiserver_to_kubelet)

	var script string
	script = strings.TrimSpace(V_setup)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
