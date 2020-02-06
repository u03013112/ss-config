package config

import (
	"context"
	"encoding/json"
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
	if in.Token != "u03013112" { //秘密入口，不校验
		role, err := grpcGetRole(in.Token)
		if err != nil {
			return nil, err
		}
		print(role) //现在role暂时没啥用
	}
	configList := getConfigList()
	return chooseSSConfig(configList, in.LineID), nil
}

// 选择算法，现在就是直接取第一个
func chooseSSConfig(configList []*Config, lineID int64) *pb.GetSSConfigReply {
	if len(configList) > 0 {
		if lineID < 0 || lineID >= int64(len(configList)) {
			lineID = 0
		}
		ret := new(pb.GetSSConfigReply)
		ret.IP = configList[lineID].IP
		ret.Port = configList[lineID].Port
		ret.Method = configList[lineID].Method
		ret.Passwd = configList[lineID].Passwd
		return ret
	}
	return nil
}

// GetSSLineList :
func (s *Srv) GetSSLineList(ctx context.Context, in *pb.GetSSLineListRequest) (*pb.GetSSLineListReply, error) {
	ret := &pb.GetSSLineListReply{
		Error: "",
	}
	uInfo, err := grpcGetUserInfo(in.Token)
	if err != nil {
		return nil, err
	}
	if uInfo.Type == "android" {
		if uInfo.Status != "normal" {
			ret.Error = uInfo.Status
			return ret, nil
		}
	}
	configList := getConfigList()
	j, _ := json.Marshal(configList)
	json.Unmarshal(j, &ret.List)
	return ret, nil
}

// GetSSLineConfig :
func (s *Srv) GetSSLineConfig(ctx context.Context, in *pb.GetSSLineConfigRequest) (*pb.GetSSLineConfigReply, error) {
	ret := &pb.GetSSLineConfigReply{
		Error: "",
	}
	uInfo, err := grpcGetUserInfo(in.Token)
	if err != nil {
		return nil, err
	}
	if uInfo.Type == "android" {
		if uInfo.Status != "normal" {
			ret.Error = uInfo.Status
			return ret, nil
		}
	}
	configList := getConfigList()
	index := in.LineID
	if index < 0 || int(index) > len(configList) {
		ret.Error = "线路不可用"
		return ret, nil
	}
	j, _ := json.Marshal(configList[index])
	json.Unmarshal(j, ret)
	return ret, nil
}

// SetPasswd :
func (s *Srv) SetPasswd(ctx context.Context, in *pb.SetPasswdRequest) (*pb.Void, error) {
	setPasswd(in.Passwd)
	return &pb.Void{}, nil
}
