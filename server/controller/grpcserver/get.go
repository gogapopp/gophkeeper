package grpcserver

import (
	"context"

	pb "github.com/gogapopp/gophkeeper/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (gs *grpcServer) SyncData(ctx context.Context, in *pb.SyncRequest) (*pb.SyncResponse, error) {
	if in == nil {
		return nil, nil
	}
	uninqueKeys := in
	response, err := gs.get.GetDatas(ctx, uninqueKeys)
	if err != nil {
		gs.log.Error(err)
		return nil, status.Error(codes.InvalidArgument, "bad request")
	}
	return response, nil
}
