NAME = sr-apps
CONTAINERS = srapps-nse
PODS = simple-client srapps-nse
CHECK = export ITERATIONS=10 && export BATCHES=1 && scripts/check_srapps.sh

#export VPP_AGENT=ligato/vpp-agent:v2.1.1

include $(TOP)/mk/targets.mk

# If GOPATH is made up of several paths, use the first one for our targets in this Makefile
GO_TOP := $(shell echo ${GOPATH} | cut -d ':' -f1)
export GO_TOP
export OUT_DIR=$(GO_TOP)/out
BUILDTYPE_DIR:=debug
export SRAPPS_OUT_LINUX:=$(OUT_DIR)/linux_amd64/$(BUILDTYPE_DIR)
# scratch dir for building isolated images
DOCKER_BUILD_TOP:=${SRAPPS_OUT_LINUX}/docker_build

.PHONY: build-srapps
build-srapps:
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -ldflags '-extldflags "-static"' -o ${GOPATH}/bin/linux_amd64/srapps_nse ./examples/sr-apps/sr-nse/cmd/...

.PHONY: docker-srapps
docker-srapps: build-srapps
	mkdir -p ${DOCKER_BUILD_TOP}/srapps/etc
	cp -p ./examples/sr-apps/sr-nse/Dockerfile.sr-nse ${DOCKER_BUILD_TOP}/srapps/
	cp -p ${GOPATH}/bin/linux_amd64/srapps_nse ${DOCKER_BUILD_TOP}/srapps/
	cp -pr ./examples/universal-cnf/vppagent/etc/* ${DOCKER_BUILD_TOP}/srapps/etc/
	cd ${DOCKER_BUILD_TOP}/srapps && docker build -t ${ORG}/sr-apps_ucnf-nse:${TAG} --build-arg VPP_AGENT=${VPP_AGENT} -f Dockerfile.sr-nse .
