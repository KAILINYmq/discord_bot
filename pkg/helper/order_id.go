package helper

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"net"
	"strconv"
	"strings"
)

var (
	payNodeId    int64 = 1
	payOrderNode *snowflake.Node
)

func init() {
	// TODO 如果 pod 数量超过 255 怎么办？虽然目前并不会
	ip := GetLocalIP()
	ipLast := parseIPV4(ip)
	if ipLast != "" {
		id, err := strconv.ParseInt(ipLast, 10, 32)
		if err == nil {
			payNodeId = id
		}
	}
	fmt.Println("node", payNodeId)
	n, err := snowflake.NewNode(payNodeId)
	if err != nil {
		panic("snow error:" + err.Error())
	}

	payOrderNode = n
}

// TODO 多 pod 环境下怎么保证不会重复
func GenOrderId() string {
	return payOrderNode.Generate().String()
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func parseIPV4(ip string) string {
	if ip == "" {
		return ""
	}

	ads := strings.Split(ip, ".")
	if len(ads) != 4 {
		return ""
	}
	return ads[3]
}
