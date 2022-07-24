package api

import (
	"context"

	pb "github.com/alex-mj/simple-grpc-srvc/service/api/proto/math"
)

// GRPCServer is a server that exposes the MathFunc service.
type GRPCServer struct {
	pb.UnimplementedMathServiceServer
}

// MathFunc implements the example of func our service.
func (GRPCServer) MathFunc(ctx context.Context, mObj *pb.MathObject) (*pb.MathResponse, error) {
	return &pb.MathResponse{Result: mObj.GetA() + mObj.GetB() + mObj.GetC()}, nil
}
