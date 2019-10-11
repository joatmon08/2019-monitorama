init:
	mkdir -p ~/.minikube/files/etc/ssl/certs
	cp setup/audit-policy.yml ~/.minikube/files/etc/ssl/certs/audit-policy.yml
	minikube start --kubernetes-version=1.15.4 --extra-config=apiserver.audit-policy-file=/etc/ssl/certs/audit-policy.yml --extra-config=apiserver.audit-log-path=- --extra-config=apiserver.authorization-mode=RBAC
	cd setup && docker-compose up -d
	helm init
	kubectl rollout status deployment/tiller-deploy -n kube-system
	helm install --name falco stable/falco

slide:
	reveal-md slides.md --theme night -w

clean:
	cd setup && docker-compose down
	minikube delete