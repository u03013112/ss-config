package config

import (
	"log"
	"net"
	"time"

	"github.com/u03013112/ss-config/ping"
	"golang.org/x/net/icmp"
)

// ScheduleInit : 计时器初始化
func ScheduleInit() {
	go func() {
		for {
			checkConfig()
			time.Sleep(time.Second * 30) //每3秒检查一次
		}
	}()
}

func checkConfig() {
	configList := getConfigListAll()
	if len(configList) > 0 {
		ipList := make([]string, 0) //所有需要ping的IP
		for _, c := range configList {
			if addrs, err := net.LookupHost(c.Domain); err == nil {
				if len(addrs) > 0 {
					c.IP = addrs[0]
					ipList = append(ipList, c.IP)
				} else {
					c.Status = -2 //没有地址
					updateConfig(c)
				}
			} else {
				c.Status = -1 //dns失败
				updateConfig(c)
			}
		}

		// 批量ping，这个是抄袭cad的探测代码
		bp, err := ping.NewBatchPinger(ipList, 1, time.Second*1, time.Second*3, true)
		if err != nil {
			log.Print(err)
			return
		}
		successIPList := make([]string, 0)
		bp.OnRecv = func(pkt *icmp.Echo, srcAddr string) {
			successIPList = append(successIPList, srcAddr)
			// 成功的IP直接改数据库
			for _, c := range configList {
				if c.IP == srcAddr {
					c.Status = 1
					updateConfig(c)
				}
			}
		}
		bp.OnFinish = func(stMap map[string]*ping.Statistics) {
			failedIPList := make([]string, 0)
			for _, v1 := range ipList {
				isExist := false
				for _, v2 := range successIPList {
					if v1 == v2 {
						isExist = true
					}
				}
				if !isExist {
					failedIPList = append(failedIPList, v1)
				}
			}
			for _, v := range failedIPList {
				for _, c := range configList {
					if c.IP == v {
						c.Status = 0
						updateConfig(c)
					}
				}
			}
		}
		bp.Run()
	}
}
