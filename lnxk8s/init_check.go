package lnxk8s

import (
	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

var V_init_check string = `
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
  if (ssh root@"$IP" "which getenforce >/dev/null 2>&1"); then
    result="$(ssh root@"$IP" "getenforce" 2>/dev/null)"
    [ "$result" != "Disabled" ] && echo "it seems selinux enabled"
  fi
done
for IP in "${UNIQUE_IP_LIST[@]}"; do
  if (root@"$IP" "systemctl status firewalld >/dev/null 2>&1"); then
    echo "it seems firewalld started"
  fi
done
for IP in "${UNIQUE_IP_LIST[@]}"; do
  result="$(ssh root@"$IP" "cat /proc/swaps |grep -v '^Filename' |wc -l" 2>/dev/null)"
  [ "$result" -ne 0 ] && echo "it seems swap is on"
  result="$(ssh root@"$IP" "cat /etc/fstab |grep 'swap' |grep -v '^#' |wc -l" 2>/dev/null)"
  [ "$result" -ne 0 ] && echo "it seems swap is on"
done
date
`

func X_init_check(args ...string) {
	log.Println(util.GetCurrentFuncName())

	log.Printf("args: %+v\n", args)

	var script string
	script = strings.TrimSpace(V_init_check)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
