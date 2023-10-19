GO_Linux_ENV := CGO_ENABLED=0 GOOS=linux
GO_Darwin_ENV := CGO_ENABLED=0 GOOS=darwin
GO_Windows_ENV := CGO_ENABLED=0 GOOS=windows
GO_ENV_AMD64 := GOARCH=amd64
GO_ENV_ARM64:= GOARCH=arm64
DIST_DIR_AMD64 := dist/x86_64
DIST_DIR_ARM64 := dist/aarch64

build: 
	$(GO_Linux_ENV) $(GO_ENV_AMD64) go build -o bin/parseAOF_linux_x86_64 ./src
	$(GO_Linux_ENV) $(GO_ENV_ARM64) go build -o bin/parseAOF_linux_aarch64 ./src
	$(GO_Darwin_ENV) $(GO_ENV_AMD64) go build -o bin/parseAOF_macos_x86_64 ./src
	$(GO_Darwin_ENV) $(GO_ENV_ARM64) go build -o bin/parseAOF_macos_arm64 ./src
	$(GO_Windows_ENV) $(GO_ENV_AMD64) go build -o bin/parseAOF_win_x86_64 ./src
	$(GO_Windows_ENV) $(GO_ENV_ARM64) go build -o bin/parseAOF_win_aarch64 ./src