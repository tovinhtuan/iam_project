package server

import (
	"context"
	"errors"

	iamgrpc "github.com/tovinhtuan/my_proto/gen_go/iam_service"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ( 
	Server struct {
		iamgrpc.IamServiceServer
	}
)

func (serv *Server) RollbackUserCredentials(ctx context.Context, req *iamgrpc.RollbackUserCredentialsRequest) (*emptypb.Empty, error) {
	return nil, errors.New("Testing")
}

func (serv *Server) mustEmbedUnimplementedIamServiceServer() {}