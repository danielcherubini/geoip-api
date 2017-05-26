package = github.com/danmademe/geoip-api

.PHONY: release

release:
	mkdir -p release
	GOOS=linux GOARCH=amd64 go build -o release/geoip-api-linux-amd64 $(package)
	GOOS=linux GOARCH=386 go build -o release/geoip-api-linux-386 $(package)
	GOOS=linux GOARCH=arm go build -o release/geoip-api-linux-arm $(package)
