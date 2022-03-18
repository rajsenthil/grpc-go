## golang with grpc

### Generating the proptobuf
Use `generate.sh` to geberate the protobuf code with greet.proto file.

#### Issues

if there is an issue `server does not implement greetpb.GreetServiceServer (missing mustEmbedUnimplementedGreetServiceServer method)`
then 
* Solution 1:

Change the code 
>> type server struct {}

to 

>> type server struct {
>>	greetpb.GreetServiceServer
>> }

* Solution 2:

>> protoc --go-grpc_out=require_unimplemented_servers=false[,other options...]:. can solve this problem.

```Note: 
With only **--go-grpc_out** the struct for message is not generated. 
To generate the struct body for message **--go_out** is needed as below
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative usermgmt/usermgmt.proto
```
