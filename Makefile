APP ?= cb
PROJECT_PATH ?= ./src
ENV ?=
DOCKERFILEBUILD ?= Dockerfile.build
REGISTRY ?= clvx 
IMAGE ?= cb
TAG ?= $(shell git rev-parse --short HEAD)
PORT ?= 3000
URL ?= http://localhost:3000

build:
	cd ${PROJECT_PATH} && go build -o bin/${APP} && cd ../

build-docker:
	docker build -f ${DOCKERFILEBUILD} --output src/bin/ .

push:
	docker build -t ${REGISTRY}/${IMAGE}:${TAG} .
	docker push ${REGISTRY}/${IMAGE}:${TAG}

run:
	./src/bin/${APP} run -u ${URL} 

serve:
	./src/bin/${APP} serve -p ${PORT}

clean:
	rm -rf ./src/bin 2> /dev/null
