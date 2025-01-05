package lnxk8s

import (
	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

var V_make_kubernetes_kubectl string = `
#!/bin/bash
set -e
set -o pipefail
set -u
cd "$(dirname "$0")"
date
source ./env.sh
echo "${KUBE_APISERVER_IP}"
mkdir -p artifact/kubernetes/kubectl
cd artifact/kubernetes/kubectl
cat <<EOF >kubectl-csr.json
{
  "CN": "kubernetes-admin",
  "key": {
    "algo": "rsa",
    "size": 2048
  },
  "names": [
    {
      "C": "CHANGEME",
      "ST": "CHANGEME",
      "L": "CHANGEME",
      "O": "system:masters",
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
  kubectl-csr.json |cfssljson -bare kubectl
KUBECONFIG="./kubectl.kubeconfig"
KUBE_APISERVER="https://${KUBE_APISERVER_IP}:6443"
[ -f "$KUBECONFIG" ] && rm -f "$KUBECONFIG"
kubectl config set-cluster kubernetes \
  --certificate-authority=../ca.pem \
  --embed-certs=true \
  --kubeconfig="${KUBECONFIG}" \
  --server="${KUBE_APISERVER}"
kubectl config set-credentials kubernetes-admin \
  --client-certificate=./kubectl.pem \
  --client-key=./kubectl-key.pem \
  --embed-certs=true \
  --kubeconfig="${KUBECONFIG}"
kubectl config set-context kubernetes-admin@kubernetes \
  --cluster=kubernetes \
  --kubeconfig="${KUBECONFIG}" \
  --user=kubernetes-admin
kubectl config use-context kubernetes-admin@kubernetes --kubeconfig="${KUBECONFIG}"
date
`

func X_make_kubernetes_kubectl(args ...string) {
	log.Println(util.GetCurrentFuncName())

	log.Printf("args: %+v\n", args)

	var script string
	script = strings.TrimSpace(V_make_kubernetes_kubectl)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
