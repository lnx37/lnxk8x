package lnxk8s

import (
	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

var V_make_kubernetes_kube_controller_manager string = `
#!/bin/bash
set -e
set -o pipefail
set -u
cd "$(dirname "$0")"
date
source ./env.sh
echo "${KUBE_APISERVER_IP}"
mkdir -p artifact/kubernetes/kube-controller-manager
cd artifact/kubernetes/kube-controller-manager
cat <<EOF >kube-controller-manager-csr.json
{
  "CN": "system:kube-controller-manager",
  "key": {
    "algo": "rsa",
    "size": 2048
  },
  "names": [
    {
      "C": "CHANGEME",
      "ST": "CHANGEME",
      "L": "CHANGEME",
      "O": "system:kube-controller-manager",
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
  kube-controller-manager-csr.json |cfssljson -bare kube-controller-manager
KUBECONFIG="./kube-controller-manager.kubeconfig"
KUBE_APISERVER="https://${KUBE_APISERVER_IP}:6443"
[ -f "$KUBECONFIG" ] && rm -f "$KUBECONFIG"
kubectl config set-cluster kubernetes \
  --certificate-authority=../ca.pem \
  --embed-certs=true \
  --kubeconfig="$KUBECONFIG" \
  --server="${KUBE_APISERVER}"
kubectl config set-credentials system:kube-controller-manager \
  --client-certificate=./kube-controller-manager.pem \
  --client-key=./kube-controller-manager-key.pem \
  --embed-certs=true \
  --kubeconfig="$KUBECONFIG"
kubectl config set-context system:kube-controller-manager@kubernetes \
  --cluster=kubernetes \
  --kubeconfig="$KUBECONFIG" \
  --user=system:kube-controller-manager
kubectl config use-context system:kube-controller-manager@kubernetes --kubeconfig="$KUBECONFIG"
cat <<EOF >kube-controller-manager.conf
KUBE_CONTROLLER_MANAGER_OPTS=" \\
  --allocate-node-cidrs=true \\
  --authentication-kubeconfig=/opt/kubernetes/kube-controller-manager/kube-controller-manager.kubeconfig \\
  --authorization-kubeconfig=/opt/kubernetes/kube-controller-manager/kube-controller-manager.kubeconfig \\
  --bind-address=127.0.0.1 \\
  --client-ca-file=/opt/kubernetes/ca.pem \\
  --cluster-cidr=10.244.0.0/16 \\
  --cluster-name=kubernetes \\
  --cluster-signing-cert-file=/opt/kubernetes/ca.pem \\
  --cluster-signing-key-file=/opt/kubernetes/ca-key.pem \\
  --controllers=*,bootstrapsigner,tokencleaner \\
  --kubeconfig=/opt/kubernetes/kube-controller-manager/kube-controller-manager.kubeconfig \\
  --leader-elect=true \\
  --requestheader-client-ca-file=/opt/kubernetes/front-proxy-ca.pem \\
  --root-ca-file=/opt/kubernetes/ca.pem \\
  --service-account-private-key-file=/opt/kubernetes/sa.key \\
  --service-cluster-ip-range=10.96.0.0/12 \\
  --use-service-account-credentials=true \\
"
EOF
cat <<\EOF >kube-controller-manager.service
[Unit]
Description=Kubernetes kube-controller-manager
Documentation=https://kubernetes.io/docs/
Wants=network-online.target
After=network-online.target
[Service]
EnvironmentFile=/opt/kubernetes/kube-controller-manager/kube-controller-manager.conf
ExecStart=/opt/kubernetes/kube-controller-manager/kube-controller-manager $KUBE_CONTROLLER_MANAGER_OPTS
Restart=always
[Install]
WantedBy=multi-user.target
EOF
date
`

func X_make_kubernetes_kube_controller_manager(args ...string) {
	log.Println(util.GetCurrentFuncName())

	log.Printf("args: %+v\n", args)

	var script string
	script = strings.TrimSpace(V_make_kubernetes_kube_controller_manager)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
