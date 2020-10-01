BINDIR         := $(CURDIR)/bin
PROTO_OUT      := genproto
MESH_VERSION   := 0.0.1
DOCKER_REPO    := docker.io/frredacs

BAR_OUT        := "bar"
BAR_PB_FILE    := $(CURDIR)/pb/bar.proto
BAR_PB_OUTPUT  := $(CURDIR)/pkg/
BAR_DOCKERFILE := Dockerfile-bar
BAR_TAG        := ${DOCKER_REPO}/bar:${MESH_VERSION}

FOO_OUT        := "foo"
FOO_PB_FILE    := $(CURDIR)/pb/foo.proto
FOO_PB_OUTPUT  := $(CURDIR)/pkg/
FOO_DOCKERFILE := Dockerfile-foo
FOO_TAG        := ${DOCKER_REPO}/foo:${MESH_VERSION}

.PHONY: all bar-api build-bar

all: build-bar build-foo

pkg/dacs/bar/bar.pb.go: pb/bar.proto
	@protoc -I $(CURDIR)/pb/ \
		-I${GOPATH}/src \
		-I${BAR_PB_OUTPUT} \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:${BAR_PB_OUTPUT} \
		${BAR_PB_FILE}

pkg/dacs/foo/foo.pb.go: pb/foo.proto
	@protoc -I $(CURDIR)/pb/ \
		-I${GOPATH}/src \
		-I${FOO_PB_OUTPUT} \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:${FOO_PB_OUTPUT} \
		${FOO_PB_FILE}

pkg/dacs/bar/bar.pb.gw.go: pb/bar.proto
	@protoc -I $(CURDIR)/pb/ \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:${BAR_PB_OUTPUT} \
		${BAR_PB_FILE}

pkg/dacs/foo/foo.pb.gw.go: pb/foo.proto
	@protoc -I $(CURDIR)/pb/ \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:${FOO_PB_OUTPUT} \
		${FOO_PB_FILE}

pkg/dacs/bar/bar.swagger.json: pb/bar.proto
	@protoc -I $(CURDIR)/pb/ \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:${BAR_PB_OUTPUT} \
		${BAR_PB_FILE}

pkg/dacs/foo/foo.swagger.json: pb/foo.proto
	@protoc -I $(CURDIR)/pb/ \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:${FOO_PB_OUTPUT} \
		${FOO_PB_FILE}

bar-api: pkg/dacs/bar/bar.pb.go pkg/dacs/bar/bar.pb.gw.go pkg/dacs/bar/bar.swagger.json ## Auto-generate grpc go sources

foo-api: pkg/dacs/foo/foo.pb.go pkg/dacs/foo/foo.pb.gw.go pkg/dacs/foo/foo.swagger.json ## Auto-generate grpc go sources

clean: ## Remove previous builds
	@rm -f pkg/dacs/bar/bar.pb.go pkg/dacs/bar/bar.pb.gw.go pkg/dacs/bar/bar.swagger.json '$(BINDIR)'/$(BAR_OUT)
	@rm -f pkg/dacs/foo/foo.pb.go pkg/dacs/foo/foo.pb.gw.go pkg/dacs/foo/foo.swagger.json '$(BINDIR)'/$(FOO_OUT)

build-bar: bar-api
	@go build -o '$(BINDIR)'/$(BAR_OUT) ./cmd/bar/main.go

build-foo: foo-api
	@go build -o '$(BINDIR)'/$(FOO_OUT) ./cmd/foo/main.go

build-bar-docker: bar-api
	@echo
	@echo "==> Creating Bar Docker Image <=="
	@docker build -t $(BAR_TAG) -f ${BAR_DOCKERFILE} .

build-foo-docker: foo-api
	@echo
	@echo "==> Creating Foo Docker Image <=="
	@docker build -t $(FOO_TAG) -f ${FOO_DOCKERFILE} .

build-docker: build-bar-docker build-foo-docker