package = github.com/danmademe/geoip-api

.PHONY: release

release:
	mkdir -p release
	GOOS=linux GOARCH=amd64 go build -o release/geoip-api-linux $(package)
	GOOS=darwin GOARCH=amd64 go build -o release/geoip-api-macos $(package)
	GOOS=windows GOARCH=amd64 go build -o release/geoip-api-win64.exe $(package)
