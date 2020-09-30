package ipv6

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/garyburd/redigo/redis"
	"myGo/ping"
	"net"
	"strconv"
	"testing"
)

func Test_ipv6(t *testing.T) {

	//ip := "1119"
	//ip2 := "111A"
	//
	//"80:12:12:09:12:12:12:12"
	//type Case struct {
	//	Range string
	//	First string
	//	Last  string
	//}
	//
	//cases := []Case{
	//	Case{
	//		Range: "192.168.0.0/16",
	//		First: "192.168.0.0",
	//		Last:  "192.168.255.255",
	//	},
	//	Case{
	//		Range: "192.168.0.0/17",
	//		First: "192.168.0.0",
	//		Last:  "192.168.127.255",
	//	},
	//	Case{
	//		Range: "fe80::/64",
	//		First: "fe80::",
	//		Last:  "fe80::ffff:ffff:ffff:ffff",
	//	},
	//}

	var address string = "240E:005A:6C01:0001:FFFF:FFFF:FFFF:FFFF/111"

	_, network, _ := net.ParseCIDR(address)
	firstIP, lastIP := cidr.AddressRange(network)
	gotFirstIP := firstIP.String()
	gotLastIP := lastIP.String()
	fmt.Println(gotFirstIP)
	fmt.Println(gotLastIP)

}

//ip到数字
func ip2Long(ip string) uint32 {
	var long uint32
	binary.Read(bytes.NewBuffer(net.ParseIP(ip).To4()), binary.BigEndian, &long)
	return long
}

//数字到IP
func backtoIP4(ipInt int64) string {
	// need to do two bit shifting and “0xff” masking
	b0 := strconv.FormatInt((ipInt>>24)&0xff, 10)
	b1 := strconv.FormatInt((ipInt>>16)&0xff, 10)
	b2 := strconv.FormatInt((ipInt>>8)&0xff, 10)
	b3 := strconv.FormatInt(ipInt&0xff, 10)
	return b0 + "." + b1 + "." + b2 + "." + b3
}

func Test_ipRange(t *testing.T) {
	//c, err := redis.Dial("tcp", "192.168.137.129:6379")
	//if err != nil {
	//	fmt.Println("Connect to redis error", err)
	//	return
	//}
	//defer c.Close()
	ipSlice := make([]string, 0)
	ip1 := ip2Long("61.187.43.1")
	ip2 := ip2Long("61.187.43.255")
	x := ip2 - ip1
	fmt.Println(ip1, ip2, x)
	if x > 10000 {
		ip2 = ip1 + 10000
	}
	for i := ip1; i <= ip2; i++ {
		i := int64(i)
		ipSlice = append(ipSlice, backtoIP4(i))
		//_, err = c.Do("SET", backtoIP4(i), 1)
		//if err != nil {
		//	fmt.Println("redis set failed:", err)
		//}
	}
	//fmt.Println(ipSlice)
	ServerPing(ipSlice)

}

func ServerPing(ipSlice []string) {
	c, err := redis.Dial("tcp", "192.168.137.129:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()
	c.Do("AUTH", "123456")
	bp, err := ping.NewBatchPinger(ipSlice, true)

	if err != nil {
		fmt.Println(err)
	}

	bp.OnFinish = func(stMap map[string]*ping.Statistics) {
		for _, st := range stMap {
			if st.PacketsRecv == 1 {
				count, err := redis.Int(c.Do("HGET", "axe.ops_resource.ping", st.Addr))
				if count > 0 {
					count++
					if count > 10 {
						count = 10
					}
					c.Do("HSET", "axe.ops_resource.ping", st.Addr, count)
					if count >= 10 {
						fmt.Printf("IP:%s,可用\n", st.Addr)
					}
				} else {
					c.Do("HSET", "axe.ops_resource.ping", st.Addr, 1)
				}
				if err != nil {
					fmt.Println("redis set failed:", err)
				}
			} else {
				count, err := redis.Int(c.Do("HGET", "axe.ops_resource.ping", st.Addr))
				if count < 0 {
					count--
					if count < -10 {
						count = -10
					}
					c.Do("HSET", "axe.ops_resource.ping", st.Addr, count)
					if count <= -10 {
						fmt.Printf("IP:%s,不可用\n", st.Addr)
					}
				} else {
					_, err = c.Do("HSET", "axe.ops_resource.ping", st.Addr, -1)
				}
				if err != nil {
					fmt.Println("redis set failed:", err)
				}
			}
		}
	}
	err = bp.Run()
	if err != nil {
		fmt.Printf("run err %v \n", err)
	}
	bp.OnFinish(bp.Statistics())

}
