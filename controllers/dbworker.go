package controllers

import (
	"../db"
	"../models"
	pb "../protobuf"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DbWorkerServer struct {}

func CreateResponseCode(code int32, msg string) *pb.ResponseCode {
	return &pb.ResponseCode{
		Code: code,
		Msg: msg,
	}
}

func ConvertUserToProto(user *models.User) *pb.User {
	return &pb.User{
		Id: int32(user.ID),
		Username: user.Username,
		Email: user.Email,
		FullName: user.FullName,
		IsSuperUser: user.IsSuperUser,
	}
}

func (dbWorker *DbWorkerServer) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	gormDb := db.GetDB()
	var user []models.User

	if err := gormDb.Find(&user).Error; err != nil {
		return nil, status.Errorf(codes.Unknown, err.Error())
	} else {
		var userList []*pb.User
		for _, v := range user {
			userList = append(userList, ConvertUserToProto(&v))
		}
		return &pb.GetUsersResponse{
			ResCode: CreateResponseCode(0, "success"),
			Users: userList,
		}, nil
	}
}
func (dbWorker *DbWorkerServer) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	var user = models.User{
		Username: req.Username,
		Email: req.Email,
		Password: req.Password,
		FullName: req.FullName,
	}

	gormDb := db.GetDB()
	gormDb.Create(&user)

	gormDb.Last(&user)

	return &pb.AddUserResponse{
		ResCode: CreateResponseCode(0, "success"),
		User: ConvertUserToProto(&user),
	}, nil
}

func (dbWorker *DbWorkerServer) GetDiscussions(ctx context.Context, req *pb.GetDiscussionsRequest) (*pb.GetDiscussionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDiscussions not implemented")
}
func (dbWorker *DbWorkerServer) AddDiscussion(ctx context.Context, req *pb.AddDiscussionRequest) (*pb.AddDiscussionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDiscussion not implemented")
}
