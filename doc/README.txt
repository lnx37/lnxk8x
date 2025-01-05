-- intro
a wrapper of https://github.com/lnx37/lnxk8s



-- requirement
2c2g



-- usage
chmod +x lnxk8x
./lnxk8x

./lnxk8x --help
./lnxk8x --add-worker --worker-ip="172.22.25.100"



-- upstream
cfssl          v1.6.4   (20230411) (latest:v1.6.5 :(20240306)) https://github.com/cloudflare/cfssl/releases
cni-plugins    v1.2.0   (20230117) (latest:v1.5.1 :(20240617)) https://github.com/containernetworking/plugins/releases
containerd     v1.6.28  (20240201) (latest:v1.7.20:(20240718)) https://github.com/containerd/containerd/releases
coredns        v1.10.1  (20230207) (latest:v1.11.1:(20231019)) https://github.com/kubernetes/kubernetes/tree/v1.28.10/cluster/addons/dns/coredns
crictl         v1.28.0  (20230814) (latest:v1.30.1:(20240711)) https://github.com/kubernetes-sigs/cri-tools/releases
etcd           v3.5.12  (20240131) (latest:v3.5.15:(20240720)) https://github.com/etcd-io/etcd/releases
flannel        v0.24.2  (20240119) (latest:v0.25.5:(20240717)) https://github.com/flannel-io/flannel/releases
kubernetes     v1.28.10 (20240515) (latest:v1.30.3:(20240717)) https://github.com/kubernetes/kubernetes/releases
metrics-server v0.7.0   (20240123) (latest:v0.7.1 :(20240327)) https://github.com/kubernetes-sigs/metrics-server/releases
runc           v1.1.11  (20240102) (latest:v1.1.13:(20240614)) https://github.com/opencontainers/runc/releases
