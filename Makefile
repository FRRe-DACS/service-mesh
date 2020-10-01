BINDIR         := $(CURDIR)/bin
BAR_OUT        := "bar"
BAR_PB_FILE    := $(CURDIR)/pb/bar.proto
BAR_PB_OUTPUT  := $(CURDIR)/pkg/
PROTO_OUT      := genproto

MESH_VERSION   := 0.0.1
DOCKER_REPO    := docker.io/frredacs

BAR_DOCKERFILE := Dockerfile-bar
BAR_TAG        := ${DOCKER_REPO}/bar:${MESH_VERSION}

FOO_DOCKERFILE := Dockerfile-foo
FOO_TAG        := ${DOCKER_REPO}/foo:${MESH_VERSION}

.PHONY: all api build

all: build

pkg/dacs/bar/bar.pb.go: pb/bar.proto
	@protoc -I $(CURDIR)/pb/ \
		-I${GOPATH}/src \
		-I${BAR_PB_OUTPUT} \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:${BAR_PB_OUTPUT} \
		${BAR_PB_FILE}

pkg/dacs/bar/bar.pb.gw.go: pb/bar.proto
	@protoc -I $(CURDIR)/pb/ \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:${BAR_PB_OUTPUT} \
		${BAR_PB_FILE}

pkg/dacs/bar/bar.swagger.json: pb/bar.proto
	@protoc -I $(CURDIR)/pb/ \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:${BAR_PB_OUTPUT} \
		${BAR_PB_FILE}

bar-api: pkg/dacs/bar/bar.pb.go pkg/dacs/bar/bar.pb.gw.go pkg/dacs/bar/bar.swagger.json ## Auto-generate grpc go sources

clean: ## Remove previous builds
	@rm -f pkg/dacs/bar/bar.pb.go pkg/dacs/bar/bar.pb.gw.go pkg/dacs/bar/bar.swagger.json $(BAR_OUT)

build-bar: bar-api
	@go build -o '$(BINDIR)'/$(BAR_OUT) ./cmd/bar/main.go

build-bar-docker: bar-api
	@docker build -t $(BAR_TAG) -f ${BAR_DOCKERFILE} .