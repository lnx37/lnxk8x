package lnxk8s

import (
	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

var V_stage_crictl string = `
#!/bin/bash
set -e
set -o pipefail
set -u
cd "$(dirname "$0")"
date
mkdir -p artifact
cd artifact
cat <<\EOF >install_crictl.sh
#!/bin/bash
set -e
set -o pipefail
set -u
cd "$(dirname "$0")"
date
cp -a crictl/crictl.yaml /etc/crictl.yaml
cp -a crictl/crictl      /usr/local/bin/crictl
date
EOF
date
`

func X_stage_crictl(args ...string) {
	log.Println(util.GetCurrentFuncName())

	log.Printf("args: %+v\n", args)

	var script string
	script = strings.TrimSpace(V_stage_crictl)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
