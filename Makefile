export TARGET_USERRPC=userrpc
export TARGET_USER=user
export DOCKER_TARGET_USERRPC=hub.zpf.com/web_kit_scaffold/$(TARGET_USERRPC)
export DOCKER_TARGET_USER=hub.zpf.com/web_kit_scaffold/$(TARGET_USER)
export DOCKER_BUILDER_TARGET_USERRPC=$(DOCKER_TARGET_USERRPC).builder
export DOCKER_BUILDER_TARGET_USER=$(DOCKER_TARGET_USER).builder

.PHONY: build
build:
	go build -o ./$(TARGET_USERRPC) ./cmd/$(TARGET_USERRPC)/
	go build -o ./$(TARGET_USER) ./cmd/$(TARGET_USER)/
.PHONY: docker-build
docker-build:
	DOCKER_BUILDKIT=0 docker build --build-arg TARGET_USERRPC --build-arg GOPRYXY	--target builder -t $(DOCKER_BUILDER_TARGET_USERRPC) -f deploy/Dockerfile-userrpc .
	DOCKER_BUILDKIT=0 docker build --build-arg TARGET_USER --build-arg GOPRYXY	--target builder -t $(DOCKER_BUILDER_TARGET_USER) -f deploy/Dockerfile-user .
.PHONY: docker
docker:
	DOCKER_BUILDKIT=0 docker build --build-arg TARGET_USERRPC -t $(DOCKER_TARGET_USERRPC):dev -f deploy/Dockerfile-userrpc .
	DOCKER_BUILDKIT=0 docker build --build-arg TARGET_USER -t $(DOCKER_TARGET_USER):dev -f deploy/Dockerfile-user .
docker-release:
	docker push $(DOCKER_BUILDER_TARGET_USER)
	docker push $(DOCKER_BUILDER_TARGET_USERRPC)
	docker push $(DOCKER_TARGET_USER):release
	docker push $(DOCKER_TARGET_USERRPC):release
docker-push:
	docker push $(DOCKER_TARGET_USER):dev
	docker push $(DOCKER_TARGET_USERRPC):dev
