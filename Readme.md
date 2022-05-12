1.初始化go包
go mod init atem-common

go mod edit -replace github.com/coreos/bbolt@v1.3.4=go.etcd.io/bbolt@v1.3.4
 
go mod edit -replace google.golang.org/grpc=google.golang.org/grpc@v1.26.0
 
go mod tidy