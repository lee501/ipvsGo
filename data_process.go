package ipvsGo

import (
	"net"
	"sort"
)

var Ranger IpRangers

func init() {
	var iprange = ReadDataFromJson("./result.json")
	IpStringToUint(iprange)
}

type IpRangers []*IpRanger
type IpRanger struct {
	Begin uint
	End	  uint
	Name  string
}
//转换切片
func IpStringToUint(ips []*IpRange){
	for _, ipv := range ips {
		r := new(IpRanger)
		r.Begin = IPString2Long(ipv.Start)
		r.End = IPString2Long(ipv.End)
		r.Name = ipv.Source
		Ranger = append(Ranger, r)
	}
	sort.Stable(Ranger)
}

func IPString2Long(ip string) uint {
	b := net.ParseIP(ip).To4()
	if b == nil {
		return 0
	}

	return uint(b[3]) | uint(b[2])<<8 | uint(b[1])<<16 | uint(b[0])<<24
}

func (ips IpRangers) Len() int {
	return len(ips)
}
func (ips IpRangers) Less(i, j int) bool {
	return ips[i].Begin < ips[j].Begin
}
func (ips IpRangers) Swap(i, j int) {
	ips[i], ips[j] = ips[j], ips[i]
}
