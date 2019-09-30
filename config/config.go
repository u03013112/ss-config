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
	configList := getConfigList()
	return chooseSSConfig(configList), nil
}

// 选择算法，现在就是直接取第一个
func chooseSSConfig(configList []*Config) *pb.GetSSConfigReply {
	if len(configList) > 0 {
		ret := new(pb.GetSSConfigReply)
		ret.IP = configList[0].IP
		ret.Port = configList[0].Port
		ret.Method = configList[0].Method
		ret.Passwd = configList[0].Passwd
		return ret
	}
	return nil
}
