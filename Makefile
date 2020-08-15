
provision:
	@echo "Provisioning GSet Cluster"	
	bash scripts/provision.sh

gset-build:
	@echo "Building GSet Docker Image"	
	docker build -t gset -f Dockerfile .

gset-run:
	@echo "Running Single GSet Docker Container"
	docker run -p 8080:8080 -d gset

info:
	echo "GSet Cluster Nodes"
	docker ps | grep 'gset'
	docker network ls | grep gset_network

clean:
	@echo "Cleaning GSet Cluster"
	docker ps -a | awk '$$2 ~ /gset/ {print $$1}' | xargs -I {} docker rm -f {}
	docker network rm gset_network

build:
	@echo "Building GSet Server"	
	go build -o bin/gset main.go

build:
	@echo "go fmt GSet Server"	
	go fmt ./...

test:
	@echo "Testing GSet"	
	go test -v --cover ./...