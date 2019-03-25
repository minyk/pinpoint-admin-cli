PWD := `pwd`

default: build

build: cleanall linux darwin windows

cleanall: clean-linux clean-darwin clean-windows
	rm -rf ./build

clean-linux:
	rm -f ./build/dcos-resources-linux

clean-darwin:
	rm -f ./build/dcos-resources-darwin

clean-windows:
	rm -f ./build/dcos-resources-windows.exe

linux: clean-linux
	docker run --rm -e "GOOS=linux" -e "GOARCH=amd64" -v $(PWD):/usr/src/github.com/minyk/pinpoint-cli -w /usr/src/github.com/minyk/pinpoint-cli golang:1.11 go build -ldflags="-s -w ${GO_LDFLAGS}" -v -o build/pinpoint-cli-linux-amd64

darwin: clean-darwin
	docker run --rm -e "GOOS=darwin" -e "GOARCH=amd64" -v $(PWD):/usr/src/github.com/minyk/pinpoint-cli -w /usr/src/github.com/minyk/pinpoint-cli golang:1.11 go build -ldflags="-s -w ${GO_LDFLAGS}" -v -o build/pinpoint-cli-darwin-amd64

windows: clean-windows
	docker run --rm -e "GOOS=windows" -e "GOARCH=amd64" -v $(PWD):/usr/src/github.com/minyk/pinpoint-cli -w /usr/src/github.com/minyk/pinpoint-cli golang:1.11 go build -ldflags="-s -w ${GO_LDFLAGS}" -v -o build/pinpoint-cli-windows-amd64.exe
