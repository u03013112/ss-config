package config

import (
	"context"
	"log"

	pb "github.com/u03013112/ss-pb/config"
)

// Srv ：服务
type Srv struct{}

// Test : for test only
func (s *Srv) Test(ctx context.Context, in *pb.TestRequest) (*pb.TestReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.TestReply{Message: "Hello " + in.GetName()}, nil
}

// GetSSConfig : 获取ss配置
func (s *Srv) GetSSConfig(ctx context.Context, in *pb.GetSSConfigRequest) (*pb.GetSSConfigReply, error) {
	ret := new(pb.GetSSConfigReply)
	ret.IP = "c9s2.jamjams.net"
	ret.Port = "58700"
	ret.Method = "aes-256-gcm"
	ret.Passwd = "xKpQV8wUVe"
	return ret, nil
}
