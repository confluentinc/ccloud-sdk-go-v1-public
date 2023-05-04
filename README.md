# ccloud-sdk-go-v1-public

> Public Golang SDK for Confluent Cloud.

## To generate / update `*pb.go` files

* Install `protoc-gen-go`, 

* run `go mod vendor` in `ccloud-sdk-go-v1-public`

* at `ccloud-sdk-go-v1-public` root dir, run the following command:
```
protoc --go_out=. --go_out=/ org/org.proto -I=./ -I=vendor/github.com/gogo/protobuf \
  -I=vendor/github.com/envoyproxy/protoc-gen-validate \
  -I=vendor 
```
