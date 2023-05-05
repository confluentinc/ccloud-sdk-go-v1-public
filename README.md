# ccloud-sdk-go-v1-public

> Public Golang SDK for Confluent Cloud.

## To generate / update `*pb.go` files

* Install `protoc-gen-go`:
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

* Run `go mod vendor`

* Run the following command, replacing `org/org.go` with the relevant file
```
protoc --go_out=. --go_out=/ org/org.proto -I=./ -I=vendor/github.com/gogo/protobuf \
  -I=vendor/github.com/envoyproxy/protoc-gen-validate \
  -I=vendor 
```

* Remove the `vendor/` directory and open a PR
