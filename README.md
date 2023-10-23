# golearn
Local repo to learn about golang , libs and concepts.
As initla project for cf metrics development.


# dev steps
- Enable proto in your project
  ```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  ```

- Generate Go Code from Protobuf
Navigate to your project directory in the terminal and run the following command:
```
protoc --go_out=. --go_opt=paths=source_relative -I. log/event/clientanalytics.proto
```
-  Install go protobuf compiler library
```
sudo apt install protobuf-compiler
sudo apt install golang-goprotobuf-dev
```

- Add proto file to your project
- Generate Go Code from Protobuf
```
protoc --go_out=. --go_opt=paths=source_relative internal_user_log.proto
protoc --go_out=. --go_opt=paths=source_relative clientanalytics.proto
```

- run command 
```
go run cmd/main.go
```