package api

import (
	"context"

	pb "github.com/alex-mj/simple-grpc-srvc/service/api/proto/math"
)

// MathServer is a server that exposes the MathFunc service.
type MathServer struct {
	pb.UnimplementedMathServiceServer
}

// MathFunc implements the example of func our math service.
func (MathServer) MathFunc(ctx context.Context, mObj *pb.MathObject) (*pb.MathResponse, error) {
	return &pb.MathResponse{Result: mObj.GetA() + mObj.GetB() + mObj.GetC()}, nil
}
