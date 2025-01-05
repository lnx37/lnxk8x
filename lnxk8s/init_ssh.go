package lnxk8s

import (
	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

var V_init_ssh string = `
#!/bin/bash
set -e
set -o pipefail
set -u
cd "$(dirname "$0")"
date
source ./env.sh
echo "${ETCD_IP_LIST[@]}"
echo "${MASTER_IP_LIST[@]}"
echo "${WORKER_IP_LIST[@]}"
UNIQUE_IP_LIST=()
IP_LIST=("${ETCD_IP_LIST[@]}" "${MASTER_IP_LIST[@]}" "${WORKER_IP_LIST[@]}")
UNIQUE_IP_STR="$(for IP in "${IP_LIST[@]}"; do echo "$IP"; done |sort -u)"
while read -r line; do UNIQUE_IP_LIST+=("$line"); done < <(echo "${UNIQUE_IP_STR}")
echo "${IP_LIST[@]}"
echo "${UNIQUE_IP_STR}"
echo "${UNIQUE_IP_LIST[@]}"
for IP in "${UNIQUE_IP_LIST[@]}"; do
  grep "^$IP " ~/.ssh/known_hosts >/dev/null 2>&1 || ssh-keyscan "$IP" >>~/.ssh/known_hosts
  timeout 5 ssh root@"$IP" "hostname"
done
date
`

func X_init_ssh(args ...string) {
	log.Println(util.GetCurrentFuncName())

	log.Printf("args: %+v\n", args)

	var script string
	script = strings.TrimSpace(V_init_ssh)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
