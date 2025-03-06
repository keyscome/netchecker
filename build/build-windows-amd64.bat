$env:GOOS="windows"; $env:GOARCH="amd64"; $env:CGO_ENABLED="0"; go build -o netchecker.exe
chmod +x netchecker.exe
