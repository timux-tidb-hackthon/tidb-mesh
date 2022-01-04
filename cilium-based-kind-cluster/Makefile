init:
	cd cilium && make kind
	cd cilium && make kind-image
	cilium install --agent-image localhost:5000/cilium/cilium-dev:local