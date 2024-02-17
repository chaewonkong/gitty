GREEN=\n\033[1;32;40m
NC=\033[0m # No Color

# 라이브러리 설치
# tidy: 미사용 라이브러리 제거
# vendor: vendor에 라이브러리 설치
ref:
	@go mod tidy -v
	@go mod vendor -v
.PHONY: ref

build:
	@GOOS=linux GOARCH=amd64 go build -o build/gitty-intel
	@GOARCH=arm64 go build -o build/gitty-arm
	@GOOS=windows GOARCH=amd64 go build -o build/gitty-win.exe
.PHONY: build

comp:
	@tar -czvf bin/gitty-macos-intel.tar.gz build/gitty-intel
	@tar -czvf bin/gitty-macos-arm.tar.gz build/gitty-arm
	@tar -czvf bin/gitty-windows.tar.gz build/gitty-win.exe