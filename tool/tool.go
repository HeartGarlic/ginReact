package tool

import (
	"github.com/thinkeridea/go-extend/exnet"
	"net"
	"reflect"
)

type Tool struct {
}

//IpToInt 将ip地址转换为整数
func (t *Tool) IpToInt(ip string) int64 {
	n, _ := exnet.IPString2Long(ip)
	return int64(n)
}

// IntToIp 将整数转换为ip地址
func (t *Tool) IntToIp(ip int) string {
	s, _ := exnet.Long2IPString(uint(ip))
	Ip1 := net.ParseIP(s) // 会得到一个16字节的byte，主要为了兼容ipv6
	n, _ := exnet.IP2Long(Ip1)
	Ip2, _ := exnet.Long2IP(n)
	reflect.DeepEqual(Ip1[12:], Ip2)
	return Ip2.String()
}
