init:
	mkdir -p ~/.minikube/files/etc/ssl/certs
	cp setup/audit-policy.yml ~/.minikube/files/etc/ssl/certs/audit-policy.yml
	minikube start --kubernetes-version=1.15.4 --extra-config=apiserver.audit-policy-file=/etc/ssl/certs/audit-policy.yml --extra-config=apiserver.audit-log-path=- --extra-config=apiserver.authorization-mode=RBAC
	helm init
	kubectl rollout status deployment/tiller-deploy -n kube-system
	helm install --name falco stable/falco
	kubectl rollout status ds/falco
	helm repo add elastic https://helm.elastic.co
	helm install --name elasticsearch elastic/elasticsearch
	kubectl rollout status ds/elasticsearch

slide:
	reveal-md slides.md --theme night -w

test-sox:
	cd fitnessfn && go test . -run SOX -v

clean:
	minikube delete

pipeline:
	fly -t concourse set-pipeline --pipeline hello --config pipeline.yml
	fly -t concourse unpause-pipeline -p hello