package config

import (
	"context"
	"log"
	"time"

	tester "github.com/u03013112/ss-pb/tester"
	user "github.com/u03013112/ss-pb/user"
	"google.golang.org/grpc"
)

const (
	userAddress  = "user:50000"
	testerAddrss = "tester:50004"
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
		// log.Fatalf("did not connect: %v", err)
		return nil, err
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

func grpcGetTestLineList() (*tester.GetSSLineListReply, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(testerAddrss, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	t := tester.NewSSTesterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := t.GetSSLineList(ctx, &tester.GetSSLineListRequest{})
	return r, err
}

func grpcGetTestLineConfig(lineID int64) (*tester.GetSSLineConfigReply, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(testerAddrss, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	t := tester.NewSSTesterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := t.GetSSLineConfig(ctx, &tester.GetSSLineConfigRequest{
		LineID: lineID,
	})
	return r, err
}
