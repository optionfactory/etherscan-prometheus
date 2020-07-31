UID=$(shell id -u)
GID=$(shell id -g)

clean:
	rm -f bin/etherscan-exporter

build:
	@mkdir -p bin
	@docker run \
		--rm=true \
		-u $(UID):$(GID) \
		-v ${PWD}/main.go:/go/src/main.go \
		-v ${PWD}/collectors:/go/src/collectors \
		-v ${PWD}/go.mod:/tmp/go.mod \
		-v ${PWD}/bin:/go/bin \
		-v ${PWD}/build.sh:/go/build.sh \
		golang:1.14 /go/build.sh 

shell:
	@docker run \
		--rm=true \
		-it \
		-u $(UID):$(GID) \
		-v ${PWD}/main.go:/go/src/main.go \
		-v ${PWD}/collectors:/go/src/collectors \
		-v ${PWD}/go.mod:/tmp/go.mod \
		-v ${PWD}/bin:/go/bin \
		-v ${PWD}/build.sh:/go/build.sh \
		golang:1.14 /bin/bash