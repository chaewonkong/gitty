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
	@GOOS=linux GOARCH=amd64 go build -o dist/intel/gitty
	@GOARCH=arm64 go build -o dist/arm/gitty
	@GOOS=windows GOARCH=amd64 go build -o dist/win/gitty.exe
.PHONY: build

comp:
	@tar -czvf comp/gitty-macos-intel.tar.gz dist/intel/gitty
	@tar -czvf comp/gitty-macos-arm.tar.gz dist/arm/gitty
	@tar -czvf comp/gitty-windows.tar.gz dist/win/gitty.exe