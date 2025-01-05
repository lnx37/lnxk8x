package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"lnxk8x/lnxk8s"
	"lnxk8x/task"
	"lnxk8x/util"
)

func GenerateEnv() {
	var err error

	var file string
	file = "env.sh"

	var content string
	content = util.ENV_SH
	content = strings.TrimSpace(content)
	content = fmt.Sprintf("%s\n", content)

	var content2 []byte
	content2 = []byte(content)

	err = ioutil.WriteFile(file, content2, 0644)
	if err != nil {
		panic(err)
	}
}

func CopyEnv() {
	var err error

	var content []byte
	content, err = ioutil.ReadFile("env.sh")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("/tmp/env.sh", content, 0644)
	if err != nil {
		panic(err)
	}
}

func ParseEnv() ([]string, []string, []string) {
	var err error

	var etcd_ip_list []string
	var master_ip_list []string
	var worker_ip_list []string

	var GetIpList func(command string) ([]string, error)
	GetIpList = func(command string) ([]string, error) {
		var err error

		var ip_list []string

		var cmd *exec.Cmd
		cmd = exec.Command("sh", "-c", command)

		var combined_output []byte
		combined_output, err = cmd.CombinedOutput()

		var output string
		output = string(combined_output)

		var fields []string
		fields = strings.Fields(output)

		var ip string
		var field string
		for _, field = range fields {
			ip = field
			ip_list = append(ip_list, ip)
		}

		return ip_list, err
	}

	var cmd string

	cmd = `source ./env.sh && echo "${ETCD_IP_LIST[@]}"`
	etcd_ip_list, err = GetIpList(cmd)
	if err != nil {
		panic(err)
	}

	cmd = `source ./env.sh && echo "${MASTER_IP_LIST[@]}"`
	master_ip_list, err = GetIpList(cmd)
	if err != nil {
		panic(err)
	}

	cmd = `source ./env.sh && echo "${WORKER_IP_LIST[@]}"`
	worker_ip_list, err = GetIpList(cmd)
	if err != nil {
		panic(err)
	}

	return etcd_ip_list, master_ip_list, worker_ip_list
}

func Bootstrap() {
	var etcd_ip_list []string
	var master_ip_list []string
	var worker_ip_list []string
	etcd_ip_list, master_ip_list, worker_ip_list = ParseEnv()
	log.Printf("etcd_ip_list: %+v\n", etcd_ip_list)
	log.Printf("master_ip_list: %+v\n", master_ip_list)
	log.Printf("worker_ip_list: %+v\n", worker_ip_list)

	// bootstrap.sh
	// |-- init.sh
	// |-- download.sh
	// |-- unpkg.sh
	// |-- make.sh
	// |-- stage.sh
	// |-- install.sh
	// `-- setup.sh

	// init.sh
	lnxk8s.X_init_ssh()
	lnxk8s.X_init_check()

	// download.sh
	{
		lnxk8s.X_download_cfssl()
		lnxk8s.X_download_etcd()
		lnxk8s.X_download_cni_plugins()
		lnxk8s.X_download_containerd()
		lnxk8s.X_download_runc()
		lnxk8s.X_download_crictl()
		lnxk8s.X_download_kubernetes()
	}

	// unpkg.sh
	{
		lnxk8s.X_unpkg_cfssl()
		lnxk8s.X_unpkg_etcd()
		lnxk8s.X_unpkg_cni_plugins()
		lnxk8s.X_unpkg_containerd()
		lnxk8s.X_unpkg_runc()
		lnxk8s.X_unpkg_crictl()
		lnxk8s.X_unpkg_kubernetes()
	}

	// make.sh
	{
		var etcd_ip string
		for _, etcd_ip = range etcd_ip_list {
			lnxk8s.X_make_etcd(etcd_ip)
		}

		lnxk8s.X_make_containerd()
		lnxk8s.X_make_crictl()
		lnxk8s.X_make_kubernetes_common()
		lnxk8s.X_make_kubernetes_kubectl()
		lnxk8s.X_make_kubernetes_kube_apiserver()
		lnxk8s.X_make_kubernetes_kube_controller_manager()
		lnxk8s.X_make_kubernetes_kube_scheduler()
		lnxk8s.X_make_kubernetes_kubelet()
		lnxk8s.X_make_kubernetes_kube_proxy()
	}

	// stage.sh
	{
		lnxk8s.X_stage_kubectl()

		var etcd_ip string
		for _, etcd_ip = range etcd_ip_list {
			lnxk8s.X_stage_etcd(etcd_ip)
		}

		// stage_kubernetes_master.sh
		lnxk8s.X_stage_kubernetes_common()
		lnxk8s.X_stage_kubernetes_kubectl()
		lnxk8s.X_stage_kubernetes_kube_apiserver()
		lnxk8s.X_stage_kubernetes_kube_controller_manager()
		lnxk8s.X_stage_kubernetes_kube_scheduler()

		// stage_kubernetes_worker.sh
		lnxk8s.X_stage_cni_plugins()
		lnxk8s.X_stage_containerd()
		lnxk8s.X_stage_runc()
		lnxk8s.X_stage_crictl()
		lnxk8s.X_stage_kubernetes_common()
		lnxk8s.X_stage_kubernetes_kubelet()
		lnxk8s.X_stage_kubernetes_kube_proxy()
	}

	lnxk8s.X_distribute()

	// install.sh
	{
		lnxk8s.X_install_kubectl()

		var etcd_ip string
		for _, etcd_ip = range etcd_ip_list {
			lnxk8s.X_install_etcd(etcd_ip)
		}

		// install_kubernetes_master.sh
		var master_ip string
		for _, master_ip = range master_ip_list {
			lnxk8s.X_install_kubernetes_common(master_ip)
			lnxk8s.X_install_kubernetes_kubectl(master_ip)
			lnxk8s.X_install_kubernetes_kube_apiserver(master_ip)
			lnxk8s.X_install_kubernetes_kube_controller_manager(master_ip)
			lnxk8s.X_install_kubernetes_kube_scheduler(master_ip)
		}

		// install_kubernetes_worker.sh
		var worker_ip string
		for _, worker_ip = range worker_ip_list {
			lnxk8s.X_install_cni_plugins(worker_ip)
			lnxk8s.X_install_containerd(worker_ip)
			lnxk8s.X_install_runc(worker_ip)
			lnxk8s.X_install_crictl(worker_ip)
			lnxk8s.X_install_kubernetes_common(worker_ip)
			lnxk8s.X_install_kubernetes_kubelet(worker_ip)
			lnxk8s.X_install_kubernetes_kube_proxy(worker_ip)
		}

		// task.X_install_calico()
		task.X_install_flannel()
		task.X_install_coredns()
		task.X_install_metrics_server()
	}

	// setup.sh
	{
		task.X_setup()
	}
}

// install_kubernetes_master.sh
func AddMaster(master_ip string) {
	lnxk8s.X_distribute(master_ip)

	lnxk8s.X_install_kubernetes_common(master_ip)
	lnxk8s.X_install_kubernetes_kubectl(master_ip)
	lnxk8s.X_install_kubernetes_kube_apiserver(master_ip)
	lnxk8s.X_install_kubernetes_kube_controller_manager(master_ip)
	lnxk8s.X_install_kubernetes_kube_scheduler(master_ip)
}

// install_kubernetes_worker.sh
func AddWorker(worker_ip string) {
	lnxk8s.X_distribute(worker_ip)

	lnxk8s.X_install_cni_plugins(worker_ip)
	lnxk8s.X_install_containerd(worker_ip)
	lnxk8s.X_install_runc(worker_ip)
	lnxk8s.X_install_crictl(worker_ip)
	lnxk8s.X_install_kubernetes_common(worker_ip)
	lnxk8s.X_install_kubernetes_kubelet(worker_ip)
	lnxk8s.X_install_kubernetes_kube_proxy(worker_ip)
}

func main() {
	defer util.TimeTaken(time.Now())

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var err error

	var debug bool
	var add_worker bool
	var worker_ip string

	flag.BoolVar(&debug, "debug", true, "Debug")
	flag.BoolVar(&add_worker, "add-worker", false, "Add Worker")
	flag.StringVar(&worker_ip, "worker-ip", "", "Worker IP")

	flag.Parse()

	log.Println("debug:", debug)
	util.SETTINGS.DEBUG = debug
	log.Printf("SETTINGS: %+v\n", util.SETTINGS)

	log.Println("add_worker:", add_worker)
	log.Println("worker_ip:", worker_ip)

	if add_worker {
		if worker_ip == "" {
			log.Println("invalid argument, need an ip")
			os.Exit(0)
		}
	}

	if !add_worker {
		_, err = os.Stat("env.sh")
		if err != nil {
			log.Println("env.sh doesn't exist, generated one, edit it, then retry")
			GenerateEnv()
			os.Exit(0)
		}

		// CopyEnv()

		Bootstrap()
	}

	if add_worker {
		AddWorker(worker_ip)
	}
}
