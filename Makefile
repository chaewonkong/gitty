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
	@GOOS=linux GOARCH=amd64 go build -o bin/intel/gitty
	@GOARCH=arm64 go build -o bin/arm/gitty
	@GOOS=windows GOARCH=amd64 go build -o bin/win/gitty.exe
.PHONY: build

comp:
	@tar -czvf dist/gitty-macos-intel.tar.gz -C bin/intel gitty
	@tar -czvf dist/gitty-macos-arm.tar.gz -C bin/arm gitty
	@tar -czvf dist/gitty-windows.tar.gz -C bin/win gitty.exe
.PHONY: comp