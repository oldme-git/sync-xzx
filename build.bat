go env -w GOARCH=386
go build -o sync-xzx.exe -tags desktop,production -ldflags "-w -s -H windowsgui"
go env -w GOARCH=amd64
