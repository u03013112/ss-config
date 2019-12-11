package config

import (
	"context"
	"log"
	"time"

	user "github.com/u03013112/ss-pb/user"
	"google.golang.org/grpc"
)

const (
	userAddress = "user:50000"
)

func grpcGetRole(token string) (string, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(userAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	u := user.NewSSUserClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := u.GetRoles(ctx, &user.GetRolesRequest{Token: token})
	if err != nil {
		log.Printf("could not GetRoles: %v", err)
		return "", err
	}
	return r.Role, nil
}

func grpcGetUserInfo(token string) (*user.GetUserInfoReply, error) {
	conn, err := grpc.Dial(userAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	u := user.NewSSUserClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := u.GetUserInfo(ctx, &user.GetUserInfoRequest{Token: token})
	if err != nil {
		log.Printf("could not GetRoles: %v", err)
		return nil, err
	}
	return r, nil
}
