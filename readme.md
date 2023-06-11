## 1st install dependencies

- go get -u google.golang.org/grpc
- go get google.golang.org/protobuf/cmd/protoc-gen-go
- go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
- go get github.com/jackc/pgx/v5
- go get github.com/jackc/pgx/v5/pgxpool

## 2nd generate proto files

> command for generating proto files
>- ```protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/blog.proto```
>- 
>- ```protoc proto/blog.proto --go_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:.```


## 3rd generate db (glog as in docker-compose)