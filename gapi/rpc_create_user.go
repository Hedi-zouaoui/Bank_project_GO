package gapi

import (
	"context"

	db "github.com/Hedi-zouaoui/go_project/db/sqlc"
	"github.com/Hedi-zouaoui/go_project/pb"
	util "github.com/Hedi-zouaoui/go_project/utils"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	
	

	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil , status.Errorf(codes.Internal , "faile dto hash password  %s " , err )
	}

	arg := db.CreateUserParams{
		Username:       req.GetUsername(),
		HashedPassword: hashedPassword,
		FullName:       req.GetFullName(),
		Email:          req.GetEmail(),
	}

	user , err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				
				return nil , status.Errorf(codes.AlreadyExists , "username already exist  %s " , err )
			}
		}
	return nil , status.Errorf(codes.Internal , "faile to create user   %s " , err )
	}
	rsp := &pb.CreateUserResponse{
		User: convertUser(user) , 

	}
	return rsp , nil 
} 
	
	