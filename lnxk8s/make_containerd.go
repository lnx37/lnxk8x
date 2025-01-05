package lnxk8s

import (
	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

var V_make_containerd string = `
#!/bin/bash
set -e
set -o pipefail
set -u
cd "$(dirname "$0")"
date
mkdir -p artifact/containerd
cd artifact/containerd
cat <<\EOF >containerd.service
[Unit]
Description=containerd container runtime
Documentation=https://containerd.io
After=network.target local-fs.target
[Service]
ExecStartPre=-/sbin/modprobe overlay
ExecStart=/usr/local/bin/containerd
Type=notify
Delegate=yes
KillMode=process
Restart=always
RestartSec=5
LimitNPROC=infinity
LimitCORE=infinity
LimitNOFILE=infinity
TasksMax=infinity
OOMScoreAdjust=-999
[Install]
WantedBy=multi-user.target
EOF
./containerd config default >config.toml
cp -a config.toml config.toml_bk_raw
sed -i "s|registry.k8s.io|registry.aliyuncs.com/google_containers|g" config.toml
sed -i "s|SystemdCgroup = false|SystemdCgroup = true|g" config.toml
sed -i '/\[plugins."io.containerd.grpc.v1.cri".registry.mirrors\]/a\
\        [plugins."io.containerd.grpc.v1.cri".registry.mirrors."docker.io"]\
\          endpoint=["https://docker.m.daocloud.io", "https://cagucih8.mirror.aliyuncs.com/"]' config.toml
diff -u config.toml_bk_raw config.toml || true
date
`

func X_make_containerd(args ...string) {
	log.Println(util.GetCurrentFuncName())

	log.Printf("args: %+v\n", args)

	var script string
	script = strings.TrimSpace(V_make_containerd)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
