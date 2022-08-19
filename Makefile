APP ?= cb
PROJECT_PATH ?= ./src
ENV ?=
DOCKERFILEBUILD ?= Dockerfile.build
REGISTRY ?= clvx
IMAGE ?= cb
#TAG ?= $(shell git rev-parse --short HEAD)
TAG ?= latest
PORT ?= 3000
URL ?= http://localhost:3000

build:
	cd ${PROJECT_PATH} && go build -o bin/${APP} && cd ../

build-docker:
	docker build -f ${DOCKERFILEBUILD} --output src/bin/ .

package:
	docker build -t ${REGISTRY}/${IMAGE}:${TAG} .

package-oci:
	flux push artifact oci://docker.io/clvx/cb-manifests:$(shell git rev-parse --short HEAD) \
	--path="./manifests" \
	--source="$(shell git config --get remote.origin.url)" \
	--revision="$(shell git branch --show-current)/$(shell git rev-parse HEAD)"
	flux tag artifact oci://docker.io/clvx/cb-manifests:$(shell git rev-parse --short HEAD) --tag latest

push:
	docker push ${REGISTRY}/${IMAGE}:${TAG}

run:
	./src/bin/${APP} run -u ${URL} 

serve:
	./src/bin/${APP} serve -p ${PORT}

clean:
	rm -rf ./src/bin 2> /dev/null


