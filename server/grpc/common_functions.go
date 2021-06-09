package grpc

import (
	"context"
	"golang.org/x/oauth2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	db "kubeIT/database"
	"kubeIT/pkg/grpc/user"
	"kubeIT/server/helpers"
	"strings"
)

func GetUserPermissions(ctx context.Context, db *db.Database, authorizer *helpers.Authorizer) (*user.User, error) {

	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, status.Errorf(codes.DataLoss, "UnaryEcho: failed to get metadata")
	}

	if t, found := md["authorization"]; found {

		if len(t) != 1 {
			return nil, status.Errorf(codes.InvalidArgument, "Unknown authorization format")
		}

		splits := strings.Split(t[0], " ")
		if splits[0] != "Bearer" {
			return nil, status.Errorf(codes.InvalidArgument, "Token type must be Bearer")
		}

		verif, _, claims := authorizer.Verify(ctx, &oauth2.Token{AccessToken: splits[1]})

		if !verif {
			return nil, status.Errorf(codes.Unauthenticated, "Token expired")
		}

		us, err := db.GetUserBySub(claims.Sub)

		return us, err

	} else {
		return nil, status.Errorf(codes.DataLoss, "No token found: failed to get authorization token")
	}

}
