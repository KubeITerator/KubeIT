package grpc

import (
	"context"
	"fmt"
	"google.golang.org/genproto/googleapis/spanner/admin/database/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func GetUserPermissions(ctx context.Context, db *database.Database, funcName string) (bool, error) {

	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return false, status.Errorf(codes.DataLoss, "UnaryEcho: failed to get metadata")
	}

	if t, found := md["authorization"]; found {
		fmt.Println(t)
		return true, nil

	} else {
		return false, status.Errorf(codes.DataLoss, "No token found: failed to get authorization token")
	}

}
