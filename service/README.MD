## Hi

I'm simple gRPC service on Golang (MathService).
I'm representing a mathematical function:
`MathFunc(MathObject) returns (MathResponse)`

`MathObject {
    int a;
    int b;
    int c;
}`

`MathResponse {
    int result;
}`

### 1) You need to start with two proto-generated files (XXX_grpc.pb.go and XXX.pb.go)

See or edit .proto-file (api/proto/mathfunc.proto)
then:
```bash
protoc --go-grpc_out=. mathfunc.proto
protoc --go_out=. mathfunc.proto          
```

Then, prehaps, need set relation with go import:
```bash
go mod tidy          
```

### 2) You need implement interface
```golang
type MathServiceServer interface {
	MathFunc(context.Context, *MathObject) (*MathResponse, error)
	mustEmbedUnimplementedMathServiceServer()
}
``` 
(from file mathfunc_grpc.pb.go)

Create new file (**math_func.go**)
Create import to our proto 

`pb "github.com/alex-mj/simple-grpc-srvc/service/api/proto/math"`

We may create very simple struct:
```golang
type GRPCServer struct {}
``` 
But we need implement also `mustEmbedUnimplementedMathServiceServer()` to have forward compatible.

So embed `UnimplementedMathServiceServer` from mathfunc_grpc.pb.go to our struct:
```golang
type GRPCServer struct {
	pb.UnimplementedMathServiceServer
} 
```

And implement math function (it's our busineslogical), in this example-case it's only simple adding.
```golang
// MathFunc implements the example of func our service.
func (GRPCServer) MathFunc(ctx context.Context, mObj *pb.MathObject) (*pb.MathResponse, error) {
	return &pb.MathResponse{Result: mObj.GetA() + mObj.GetB() + mObj.GetC()}, nil
}
```
### 3) You need run grpc service

In cmd/main.go release run-code:

create grpc server
```golang
    s := grpc.NewServer()
```
link our handler to the grpc server
```golang
	pb.RegisterMathServiceServer(s, &api.MathServer{})
```
If all OK (when interface implemented) go next. We need in listener for our server.
```golang
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
```
link it with grpc server
```golang
	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
```
try it launch
```bash 
go run cmd/main.go
```
if OK? go to client -=>
