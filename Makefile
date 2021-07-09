.PHONE: build-docker
build-docker:
	docker build -t builder -f Dockerfile .


.PHONE: pbuf-gen

BUFDIR = "/tmp/buf-generate"
pbuf: build-docker
	 docker run --rm --user "$(shell id -u):$(shell id -g)" -v $(PWD):/app builder /scripts/build.sh
