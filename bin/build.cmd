
go build -ldflags "-X main.build=`git rev-parse HEAD`" main/cmd.go