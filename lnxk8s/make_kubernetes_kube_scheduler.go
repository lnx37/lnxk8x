package lnxk8s

import (
	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

var V_make_kubernetes_kube_scheduler string = `
#!/bin/bash
set -e
set -o pipefail
set -u
cd "$(dirname "$0")"
date
source ./env.sh
echo "${KUBE_APISERVER_IP}"
mkdir -p artifact/kubernetes/kube-scheduler
cd artifact/kubernetes/kube-scheduler
cat <<EOF >kube-scheduler-csr.json
{
  "CN": "system:kube-scheduler",
  "key": {
    "algo": "rsa",
    "size": 2048
  },
  "names": [
    {
      "C": "CHANGEME",
      "ST": "CHANGEME",
      "L": "CHANGEME",
      "O": "system:kube-scheduler",
      "OU": "CHANGEME"
    }
  ]
}
EOF
cfssl gencert \
  -ca=../ca.pem \
  -ca-key=../ca-key.pem \
  -config=../ca-config.json \
  -profile=kubernetes \
  kube-scheduler-csr.json |cfssljson -bare kube-scheduler
KUBECONFIG="./kube-scheduler.kubeconfig"
KUBE_APISERVER="https://${KUBE_APISERVER_IP}:6443"
[ -f "$KUBECONFIG" ] && rm -f "$KUBECONFIG"
kubectl config set-cluster kubernetes \
  --certificate-authority=../ca.pem \
  --embed-certs=true \
  --kubeconfig="$KUBECONFIG" \
  --server="${KUBE_APISERVER}"
kubectl config set-credentials system:kube-scheduler \
  --client-certificate=./kube-scheduler.pem \
  --client-key=./kube-scheduler-key.pem \
  --embed-certs=true \
  --kubeconfig="$KUBECONFIG"
kubectl config set-context system:kube-scheduler@kubernetes \
  --cluster=kubernetes \
  --kubeconfig="$KUBECONFIG" \
  --user=system:kube-scheduler
kubectl config use-context system:kube-scheduler@kubernetes --kubeconfig="$KUBECONFIG"
cat <<EOF >kube-scheduler.conf
KUBE_SCHEDULER_OPTS=" \\
  --authentication-kubeconfig=/opt/kubernetes/kube-scheduler/kube-scheduler.kubeconfig \\
  --authorization-kubeconfig=/opt/kubernetes/kube-scheduler/kube-scheduler.kubeconfig \\
  --bind-address=127.0.0.1 \\
  --kubeconfig=/opt/kubernetes/kube-scheduler/kube-scheduler.kubeconfig \\
  --leader-elect=true \\
"
EOF
cat <<\EOF >kube-scheduler.service
[Unit]
Description=Kubernetes kube-scheduler
Documentation=https://kubernetes.io/docs/
Wants=network-online.target
After=network-online.target
[Service]
EnvironmentFile=/opt/kubernetes/kube-scheduler/kube-scheduler.conf
ExecStart=/opt/kubernetes/kube-scheduler/kube-scheduler $KUBE_SCHEDULER_OPTS
Restart=always
[Install]
WantedBy=multi-user.target
EOF
date
`

func X_make_kubernetes_kube_scheduler(args ...string) {
	log.Println(util.GetCurrentFuncName())

	log.Printf("args: %+v\n", args)

	var script string
	script = strings.TrimSpace(V_make_kubernetes_kube_scheduler)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
