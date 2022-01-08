init:
	cd cilium && make kind
	cd cilium && make kind-image
	cilium install --agent-image localhost:5000/cilium/cilium-dev:local

dev:
	cd cilium && make kind-image
	kubectl delete pods -l k8s-app=cilium -n kube-system

cloud-init:
	cd cilium && make kind
	cilium install --agent-image registry-vpc.cn-heyuan.aliyuncs.com/handlerww/cilium-dev:latest

cloud-dev:
	kubectl delete pods -l k8s-app=cilium -n kube-system

deploy-tidb-operator:
	kubectl create -f https://raw.githubusercontent.com/pingcap/tidb-operator/master/manifests/crd.yaml
	helm repo add pingcap https://charts.pingcap.org/
	kubectl create namespace tidb-admin
	helm install --namespace tidb-admin tidb-operator pingcap/tidb-operator --version v1.2.5 \
		--set operatorImage=registry.cn-beijing.aliyuncs.com/tidb/tidb-operator:v1.2.5 \
		--set tidbBackupManagerImage=registry.cn-beijing.aliyuncs.com/tidb/tidb-backup-manager:v1.2.5 \
		--set scheduler.kubeSchedulerImageName=registry.cn-hangzhou.aliyuncs.com/google_containers/kube-scheduler

deploy-tidb:
	kubectl create ns pingcap
	kubectl ns pingcap
	kubectl apply -f example-tidb-clusters.yaml
	kubectl apply -f tenant-isolation-policy.yaml

enable-cilium-ui:
	cilium hubble enable --ui

download-cilium-cli-ui:
	curl -L --remote-name-all https://github.com/cilium/cilium-cli/releases/latest/download/cilium-linux-amd64.tar.gz && sudo tar xzvfC cilium-linux-amd64.tar.gz /usr/local/bin && rm cilium-linux-amd64.tar.gz
	export HUBBLE_VERSION=$(curl -s https://raw.githubusercontent.com/cilium/hubble/master/stable.txt)
	curl -L --remote-name-all https://github.com/cilium/hubble/releases/download/$HUBBLE_VERSION/hubble-linux-amd64.tar.gz && sudo tar xzvfC hubble-linux-amd64.tar.gz /usr/local/bin && rm hubble-linux-amd64.tar.gz

logs:
	kubectl get pod -l k8s-app=cilium -n kube-system |grep cilium | awk '{system ("kubectl logs -n kube-system " $1 " > " $1 ".log && code " $1 ".log")}'

test:
	go test -timeout 30s -run ^Test github.com/cilium/cilium/proxylib/tidb -v