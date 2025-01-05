package lnxk8s

import (
	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

var V_make_kubernetes_kube_proxy string = `
#!/bin/bash
set -e
set -o pipefail
set -u
cd "$(dirname "$0")"
date
source ./env.sh
echo "${KUBE_APISERVER_IP}"
echo "${KUBE_PROXY_MODE}"
mkdir -p artifact/kubernetes/kube-proxy
cd artifact/kubernetes/kube-proxy
cat <<EOF >kube-proxy-csr.json
{
  "CN": "system:kube-proxy",
  "key": {
    "algo": "rsa",
    "size": 2048
  },
  "names": [
    {
      "C": "CHANGEME",
      "ST": "CHANGEME",
      "L": "CHANGEME",
      "O": "system:node-proxier",
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
  kube-proxy-csr.json |cfssljson -bare kube-proxy
KUBECONFIG="./kube-proxy.kubeconfig"
KUBE_APISERVER="https://${KUBE_APISERVER_IP}:6443"
[ -f "$KUBECONFIG" ] && rm -f "$KUBECONFIG"
kubectl config set-cluster default \
  --certificate-authority=../ca.pem \
  --embed-certs=true \
  --kubeconfig="$KUBECONFIG" \
  --server="${KUBE_APISERVER}"
kubectl config set-credentials default \
  --client-certificate=./kube-proxy.pem \
  --client-key=./kube-proxy-key.pem \
  --embed-certs=true \
  --kubeconfig="$KUBECONFIG"
kubectl config set-context default \
  --cluster=default \
  --kubeconfig="$KUBECONFIG" \
  --user=default
kubectl config use-context default --kubeconfig="$KUBECONFIG"
cat <<EOF >kube-proxy-config.yml
apiVersion: kubeproxy.config.k8s.io/v1alpha1
bindAddress: 0.0.0.0
bindAddressHardFail: false
clientConnection:
  acceptContentTypes: ""
  burst: 0
  contentType: ""
  # kubeconfig: /var/lib/kube-proxy/kubeconfig.conf
  kubeconfig: /opt/kubernetes/kube-proxy/kube-proxy.kubeconfig
  qps: 0
clusterCIDR: 10.244.0.0/16
configSyncPeriod: 0s
conntrack:
  maxPerCore: null
  min: null
  tcpCloseWaitTimeout: null
  tcpEstablishedTimeout: null
detectLocal:
  bridgeInterface: ""
  interfaceNamePrefix: ""
detectLocalMode: ""
enableProfiling: false
healthzBindAddress: ""
hostnameOverride: ""
iptables:
  localhostNodePorts: null
  masqueradeAll: false
  masqueradeBit: null
  minSyncPeriod: 0s
  syncPeriod: 0s
ipvs:
  excludeCIDRs: null
  minSyncPeriod: 0s
  scheduler: ""
  strictARP: false
  syncPeriod: 0s
  tcpFinTimeout: 0s
  tcpTimeout: 0s
  udpTimeout: 0s
kind: KubeProxyConfiguration
logging:
  flushFrequency: 0
  options:
    json:
      infoBufferSize: "0"
  verbosity: 0
metricsBindAddress: ""
mode: "${KUBE_PROXY_MODE}"
nodePortAddresses: null
oomScoreAdj: null
portRange: ""
showHiddenMetricsForVersion: ""
winkernel:
  enableDSR: false
  forwardHealthCheckVip: false
  networkName: ""
  rootHnsEndpointName: ""
  sourceVip: ""
EOF
cat <<EOF >kube-proxy.conf
KUBE_PROXY_OPTS=" \\
  --config=/opt/kubernetes/kube-proxy/kube-proxy-config.yml \\
"
EOF
cat <<\EOF >kube-proxy.service
[Unit]
Description=Kubernetes kube-proxy
Documentation=https://kubernetes.io/docs/
Wants=network-online.target
After=network-online.target containerd.service cri-docker.service docker.service
[Service]
EnvironmentFile=/opt/kubernetes/kube-proxy/kube-proxy.conf
ExecStart=/opt/kubernetes/kube-proxy/kube-proxy $KUBE_PROXY_OPTS
Restart=always
[Install]
WantedBy=multi-user.target
EOF
date
`

func X_make_kubernetes_kube_proxy(args ...string) {
	log.Println(util.GetCurrentFuncName())

	log.Printf("args: %+v\n", args)

	var script string
	script = strings.TrimSpace(V_make_kubernetes_kube_proxy)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
