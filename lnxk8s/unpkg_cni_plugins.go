package lnxk8s

import (
	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

var V_unpkg_cni_plugins string = `
#!/bin/bash
set -e
set -o pipefail
set -u
cd "$(dirname "$0")"
date
mkdir -p pkg
mkdir -p artifact/cni-plugins
cd pkg
[ -d cni-plugins-linux-amd64-v1.2.0 ] && rm -rf cni-plugins-linux-amd64-v1.2.0
mkdir -p cni-plugins-linux-amd64-v1.2.0
tar xzf cni-plugins-linux-amd64-v1.2.0.tgz -C cni-plugins-linux-amd64-v1.2.0
chown -R root:root cni-plugins-linux-amd64-v1.2.0
chmod +x cni-plugins-linux-amd64-v1.2.0/bandwidth
chmod +x cni-plugins-linux-amd64-v1.2.0/bridge
chmod +x cni-plugins-linux-amd64-v1.2.0/dhcp
chmod +x cni-plugins-linux-amd64-v1.2.0/dummy
chmod +x cni-plugins-linux-amd64-v1.2.0/firewall
chmod +x cni-plugins-linux-amd64-v1.2.0/host-device
chmod +x cni-plugins-linux-amd64-v1.2.0/host-local
chmod +x cni-plugins-linux-amd64-v1.2.0/ipvlan
chmod +x cni-plugins-linux-amd64-v1.2.0/loopback
chmod +x cni-plugins-linux-amd64-v1.2.0/macvlan
chmod +x cni-plugins-linux-amd64-v1.2.0/portmap
chmod +x cni-plugins-linux-amd64-v1.2.0/ptp
chmod +x cni-plugins-linux-amd64-v1.2.0/sbr
chmod +x cni-plugins-linux-amd64-v1.2.0/static
chmod +x cni-plugins-linux-amd64-v1.2.0/tuning
chmod +x cni-plugins-linux-amd64-v1.2.0/vlan
chmod +x cni-plugins-linux-amd64-v1.2.0/vrf
cp -a cni-plugins-linux-amd64-v1.2.0/bandwidth   ../artifact/cni-plugins/bandwidth
cp -a cni-plugins-linux-amd64-v1.2.0/bridge      ../artifact/cni-plugins/bridge
cp -a cni-plugins-linux-amd64-v1.2.0/dhcp        ../artifact/cni-plugins/dhcp
cp -a cni-plugins-linux-amd64-v1.2.0/dummy       ../artifact/cni-plugins/dummy
cp -a cni-plugins-linux-amd64-v1.2.0/firewall    ../artifact/cni-plugins/firewall
cp -a cni-plugins-linux-amd64-v1.2.0/host-device ../artifact/cni-plugins/host-device
cp -a cni-plugins-linux-amd64-v1.2.0/host-local  ../artifact/cni-plugins/host-local
cp -a cni-plugins-linux-amd64-v1.2.0/ipvlan      ../artifact/cni-plugins/ipvlan
cp -a cni-plugins-linux-amd64-v1.2.0/loopback    ../artifact/cni-plugins/loopback
cp -a cni-plugins-linux-amd64-v1.2.0/macvlan     ../artifact/cni-plugins/macvlan
cp -a cni-plugins-linux-amd64-v1.2.0/portmap     ../artifact/cni-plugins/portmap
cp -a cni-plugins-linux-amd64-v1.2.0/ptp         ../artifact/cni-plugins/ptp
cp -a cni-plugins-linux-amd64-v1.2.0/sbr         ../artifact/cni-plugins/sbr
cp -a cni-plugins-linux-amd64-v1.2.0/static      ../artifact/cni-plugins/static
cp -a cni-plugins-linux-amd64-v1.2.0/tuning      ../artifact/cni-plugins/tuning
cp -a cni-plugins-linux-amd64-v1.2.0/vlan        ../artifact/cni-plugins/vlan
cp -a cni-plugins-linux-amd64-v1.2.0/vrf         ../artifact/cni-plugins/vrf
date
`

func X_unpkg_cni_plugins(args ...string) {
	log.Println(util.GetCurrentFuncName())

	log.Printf("args: %+v\n", args)

	var script string
	script = strings.TrimSpace(V_unpkg_cni_plugins)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
