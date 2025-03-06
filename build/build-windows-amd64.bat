$env:GOOS="windows"; $env:GOARCH="amd64"; $env:CGO_ENABLED="0"; go build -o out/netchecker-amd64.exe
chmod +x out/netchecker-amd64.exe
