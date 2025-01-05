package lnxk8s

import (
	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

var V_stage_kubernetes_common string = `
#!/bin/bash
set -e
set -o pipefail
set -u
cd "$(dirname "$0")"
date
mkdir -p artifact
cd artifact
cat <<\EOF >install_kubernetes_common.sh
#!/bin/bash
set -e
set -o pipefail
set -u
cd "$(dirname "$0")"
date
if (which yum >/dev/null 2>&1); then
  yum install conntrack -y -q
  yum install ipset -y -q
  yum install ipvsadm -y -q
fi
if (which apt >/dev/null 2>&1); then
  apt update >/dev/null 2>&1
  apt install conntrack -y >/dev/null 2>&1
  apt install ipset -y >/dev/null 2>&1
  apt install ipvsadm -y >/dev/null 2>&1
fi
cat <<EOT >/etc/modules-load.d/kubernetes.conf
br_netfilter
EOT
modprobe br_netfilter
cat <<EOT >/etc/sysctl.d/kubernetes.conf
net.bridge.bridge-nf-call-ip6tables=1
net.bridge.bridge-nf-call-iptables=1
net.ipv4.ip_forward=1
EOT
sysctl -p /etc/sysctl.d/kubernetes.conf
date
EOF
date
`

func X_stage_kubernetes_common(args ...string) {
	log.Println(util.GetCurrentFuncName())

	log.Printf("args: %+v\n", args)

	var script string
	script = strings.TrimSpace(V_stage_kubernetes_common)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
