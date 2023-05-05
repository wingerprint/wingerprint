build:
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w"

buildStealth:
	GOOS=windows GOARCH=amd64 garble -literals -seed random -tiny build

buildDocker:
	docker run --rm -v "$$PWD":/usr/src/wingerprint -w /usr/src/wingerprint golang sh -c "go install mvdan.cc/garble@latest && GOOS=windows GOARCH=amd64 garble -literals -seed random -tiny build"
