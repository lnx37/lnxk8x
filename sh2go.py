import os
import shutil
from jinja2 import Environment
from jinja2 import FileSystemLoader

lnxk8s = '../lnxk8s'

if os.path.exists('lnxk8s'):
    shutil.rmtree('lnxk8s')
os.mkdir('lnxk8s')

files = os.listdir(lnxk8s)
files = sorted(files)

targets = [
    'init_ssh.sh',
    'init_check.sh',

    'download_cfssl.sh',
    'download_etcd.sh',
    'download_cni_plugins.sh',
    'download_containerd.sh',
    'download_runc.sh',
    'download_crictl.sh',
    # 'download_crio.sh',
    # 'download_docker.sh',
    # 'download_cri_dockerd.sh',
    'download_kubernetes.sh',

    'unpkg_cfssl.sh',
    'unpkg_etcd.sh',
    'unpkg_cni_plugins.sh',
    'unpkg_containerd.sh',
    'unpkg_runc.sh',
    'unpkg_crictl.sh',
    # 'unpkg_crio.sh',
    # 'unpkg_docker.sh',
    # 'unpkg_cri_dockerd.sh',
    'unpkg_kubernetes.sh',

    'make_etcd.sh',
    'make_containerd.sh',
    'make_crictl.sh',
    # 'make_crio.sh',
    # 'make_docker.sh',
    # 'make_cri_dockerd.sh',
    'make_kubernetes_common.sh',
    'make_kubernetes_kubectl.sh',
    'make_kubernetes_kube_apiserver.sh',
    'make_kubernetes_kube_controller_manager.sh',
    'make_kubernetes_kube_scheduler.sh',
    'make_kubernetes_kubelet.sh',
    'make_kubernetes_kube_proxy.sh',

    'stage_kubectl.sh',
    'stage_etcd.sh',

    # stage_kubernetes_master.sh
    'stage_kubernetes_common.sh',
    'stage_kubernetes_kubectl.sh',
    'stage_kubernetes_kube_apiserver.sh',
    'stage_kubernetes_kube_controller_manager.sh',
    'stage_kubernetes_kube_scheduler.sh',

    # stage_kubernetes_worker.sh
    'stage_cni_plugins.sh',
    'stage_containerd.sh',
    'stage_runc.sh',
    'stage_crictl.sh',
    # 'stage_crio.sh',
    # 'stage_docker.sh',
    # 'stage_cri_dockerd.sh',
    'stage_kubernetes_common.sh',
    'stage_kubernetes_kubelet.sh',
    'stage_kubernetes_kube_proxy.sh',

    'distribute.sh',

    'install_kubectl.sh',
    'install_etcd.sh',

    # install_kubernetes_master.sh
    'install_kubernetes_common.sh',
    'install_kubernetes_kubectl.sh',
    'install_kubernetes_kube_apiserver.sh',
    'install_kubernetes_kube_controller_manager.sh',
    'install_kubernetes_kube_scheduler.sh',

    # install_kubernetes_worker.sh
    'install_common.sh',
    'install_cni_plugins.sh',
    'install_containerd.sh',
    'install_runc.sh',
    'install_crictl.sh',
    # 'install_crio.sh',
    # 'install_docker.sh',
    # 'install_cri_dockerd.sh',
    'install_kubernetes_common.sh',
    'install_kubernetes_kubelet.sh',
    'install_kubernetes_kube_proxy.sh',
]

for file in files:
    if file in targets:
        # print(file)

        shfile = os.path.join(lnxk8s, file)
        # print(shfile)

        gofile = os.path.join('lnxk8s', file)
        gofile = gofile.replace('.sh', '.go')
        print('%s -> %s' % (shfile, gofile))

        with open(shfile, 'r') as f, open(gofile, 'w') as f2:
            lines2 = []
            lines = f.readlines()
            for line in lines:
                line = line.rstrip()
                if line.startswith('# '):
                    continue
                if line == '#':
                    continue
                if line == '':
                    continue
                # if line == 'cd "$(dirname "$0")"':
                #     continue
                if line == 'set -x':
                    continue
                lines2.append(line)
            content = os.linesep.join(lines2)

            func = 'X_%s' % file.replace('.sh', '')
            var = 'V_%s' % file.replace('.sh', '')

            template = 'template.go'
            destination = gofile
            data = {
                'func': func,
                'var': var,
                'content': content,
            }

            env = Environment(loader=FileSystemLoader('template'), trim_blocks=True)
            template = env.get_template(template)
            output = template.render(data=data)
            f2.write(output)
            f2.write(os.linesep)
