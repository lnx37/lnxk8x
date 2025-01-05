package lnxk8s

import (
	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

var V_make_kubernetes_common string = `
#!/bin/bash
set -e
set -o pipefail
set -u
cd "$(dirname "$0")"
date
mkdir -p artifact/kubernetes
cd artifact/kubernetes
cat <<EOF >ca-config.json
{
  "signing": {
    "default": {
      "expiry": "876000h"
    },
    "profiles": {
      "kubernetes": {
        "expiry": "876000h",
        "usages": [
          "signing",
          "key encipherment",
          "server auth",
          "client auth"
        ]
      }
    }
  }
}
EOF
cat <<EOF >ca-csr.json
{
  "CN": "kubernetes-ca",
  "key": {
    "algo": "rsa",
    "size": 2048
  },
  "names": [
    {
      "C": "CHANGEME",
      "ST": "CHANGEME",
      "L": "CHANGEME",
      "O": "kubernetes",
      "OU": "CHANGEME"
    }
  ]
}
EOF
cfssl gencert -initca ca-csr.json |cfssljson -bare ca
cat <<EOF >front-proxy-ca-config.json
{
  "signing": {
    "default": {
      "expiry": "876000h"
    },
    "profiles": {
      "kubernetes": {
        "expiry": "876000h",
        "usages": [
          "signing",
          "key encipherment",
          "server auth",
          "client auth"
        ]
      }
    }
  }
}
EOF
cat <<EOF >front-proxy-ca-csr.json
{
  "CN": "front-proxy-ca",
  "key": {
     "algo": "rsa",
     "size": 2048
  }
}
EOF
cfssl gencert -initca front-proxy-ca-csr.json |cfssljson -bare front-proxy-ca
cat <<EOF >front-proxy-client-csr.json
{
  "CN": "front-proxy-client",
  "key": {
     "algo": "rsa",
     "size": 2048
  }
}
EOF
cfssl gencert \
  -ca=front-proxy-ca.pem \
  -ca-key=front-proxy-ca-key.pem \
  -config=front-proxy-ca-config.json \
  -profile=kubernetes \
  front-proxy-client-csr.json |cfssljson -bare front-proxy-client
openssl genrsa -out sa.key 2048
openssl rsa -in sa.key -pubout -out sa.pub
date
`

func X_make_kubernetes_common(args ...string) {
	log.Println(util.GetCurrentFuncName())

	log.Printf("args: %+v\n", args)

	var script string
	script = strings.TrimSpace(V_make_kubernetes_common)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
