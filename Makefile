WORKDIR=/go/src/github.com/naoina/vocanew
IMAGE=vocanew:latest
DOCKER=docker
DOCKER_BUILD=$(DOCKER) build -t $(IMAGE)
DOCKER_RUN=$(DOCKER) run -v $(PWD):$(WORKDIR) -w $(WORKDIR) $(IMAGE)

.PHONY: all build

all: build

docker:
	$(DOCKER_BUILD) .

build:
	$(DOCKER_RUN) /bin/bash -c 'dep ensure && kocha build'

run:
	$(DOCKER_RUN) /bin/bash -c 'dep ensure && kocha run'
