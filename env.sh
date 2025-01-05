#!/bin/bash

# shellcheck disable=SC2034

# MASTER_IP_LIST=("172.22.23.70")
# MASTER_IP_LIST=("172.22.23.70" "172.22.23.71" "172.22.23.72")
MASTER_IP_LIST=("172.22.23.70")

# WORKER_IP_LIST=("172.22.23.70")
# WORKER_IP_LIST=("172.22.23.70" "172.22.23.71" "172.22.23.72")
# WORKER_IP_LIST=("${MASTER_IP_LIST[0]}")
WORKER_IP_LIST=("${MASTER_IP_LIST[0]}")

# ETCD_IP_LIST=("172.22.23.70")
# ETCD_IP_LIST=("172.22.23.70" "172.22.23.71" "172.22.23.72")
# ETCD_IP_LIST=("${MASTER_IP_LIST[0]}")
ETCD_IP_LIST=("${MASTER_IP_LIST[0]}")

# KUBE_APISERVER_IP="172.22.23.70"
# KUBE_APISERVER_IP="${MASTER_IP_LIST[0]}"
KUBE_APISERVER_IP="${MASTER_IP_LIST[0]}"

# KUBERNETES_VERSION="v1.18.20"
# KUBERNETES_VERSION="v1.19.16"
# KUBERNETES_VERSION="v1.20.15"
# KUBERNETES_VERSION="v1.21.14"
# KUBERNETES_VERSION="v1.22.17"
# KUBERNETES_VERSION="v1.23.17"
# KUBERNETES_VERSION="v1.24.17"
# KUBERNETES_VERSION="v1.25.14"
# KUBERNETES_VERSION="v1.26.9"
# KUBERNETES_VERSION="v1.27.6"
# KUBERNETES_VERSION="v1.28.10"
# KUBERNETES_VERSION="v1.29.4"
# KUBERNETES_VERSION="v1.30.1"
KUBERNETES_VERSION="v1.28.10"

# CONTAINER_RUNTIME="containerd"
# CONTAINER_RUNTIME="crio"
# CONTAINER_RUNTIME="docker"
CONTAINER_RUNTIME="containerd"

# NETWORK_PLUGIN="calico"
# NETWORK_PLUGIN="cilium"
# NETWORK_PLUGIN="flannel"
NETWORK_PLUGIN="flannel"

# KUBE_PROXY_MODE="iptables"
# KUBE_PROXY_MODE="ipvs"
KUBE_PROXY_MODE="ipvs"

# NETWORK_REGION="us"
# NETWORK_REGION="cn"
NETWORK_REGION="cn"

# OS_ARCH="amd64"
# OS_ARCH="arm64"
OS_ARCH="amd64"

# DRY_RUN="no"
# DRY_RUN="yes"
DRY_RUN="no"

# WITH_DOWNLOAD="no"
# WITH_DOWNLOAD="yes"
WITH_DOWNLOAD="yes"
