$env:GOOS="windows"; $env:GOARCH="amd64"; $env:CGO_ENABLED="0"; go build -o pulse.exe
chmod +x pulse.exe
