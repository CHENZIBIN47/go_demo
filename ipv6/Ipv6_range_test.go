package ipv6

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func Test_ipv6_range(t *testing.T) {
	str := "sssdaaa;"

	str = strings.TrimRight(str, ";")

	fmt.Println(str)
}

func convert2CompleteIpV6(ip string) string {
	var ipV6 string = ip
	index := strings.Index(ip, "::")
	if index > 0 {
		size := 8 - len(strings.Split(ip, ":")) - 1
		var tmp string = ""
		for i := 0; i < size; i++ {
			tmp += ":0000"
		}
		tmp += ":"
		//ipV6 = ip.replace("::",tmp);
		ipV6 = strings.Replace(ipV6, "::", tmp, -1)
	} else if index == 0 {
		if ip == "::" {
			ipV6 = "0000:0000:0000:0000:0000:0000:0000:0000"
		} else {
			ipV6 = strings.Replace(ipV6, "::", "0000:0000:0000:0000:0000:0000:0000:", -1)
		}
	}
	return ipV6
}

func bin_to_hex(mask string) []string {
	//var hex_array = [{key:0,val:"0000"},{key:1,val:"0001"},{key:2,val:"0010"},{key:3,val:"0011"},{key:4,val:"0100"},{key:5,val:"0101"},{key:6,val:"0110"},{key:7,val:"0111"},
	//{key:8,val:"1000"},{key:9,val:"1001"},{key:'A',val:"1010"},{key:'B',val:"1011"},{key:'C',val:"1100"},{key:'D',val:"1101"},{key:'E',val:"1110"},{key:'F',val:"1111"}]
	mapData := make(map[string]string)
	mapData["0"] = "0000"
	mapData["1"] = "0001"
	mapData["2"] = "0010"
	mapData["3"] = "0011"
	mapData["4"] = "0100"
	mapData["5"] = "0101"
	mapData["6"] = "0110"
	mapData["7"] = "0111"
	mapData["8"] = "1000"
	mapData["9"] = "1001"
	mapData["A"] = "1010"
	mapData["B"] = "1011"
	mapData["C"] = "1100"
	mapData["D"] = "1101"
	mapData["E"] = "1110"
	mapData["F"] = "1111"

	var value string = ""

	list := make([]string, 0)

	var str string = ""
	parseInt, err := strconv.Atoi(mask)
	//parseInt, err := strconv.ParseInt(mask, 10, 64)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if parseInt != 0 {
		for i := 0; i < parseInt; i++ {
			str += "1"
		}

		for i := 0; i < 128-parseInt; i++ {
			str += "0"
		}
	}
	for {
		if len(str) > 4 {
			list = append(list, str[0:4])
			str = str[0:4]
		} else {
			break
		}
	}
	list = append(list, str)
	for i := 0; i < len(list); i++ {
		for key, array := range mapData {
			if list[i] == array {
				strings.Contains(value, key)
				break
			}
		}
	}

	valArr := make([]string, 0)
	for {
		if len(value) > 4 {
			valArr = append(valArr, str[0:4])
			value = value[0:4]
		} else {
			break
		}
	}
	valArr = append(valArr, value)
	return valArr

}

//func checkIpv6(ip,mask string,ranges[]){
//	//var flag bool = true
//	//masks := bin_to_hex(mask)
//	//ipv6 := strings.Split(convert2CompleteIpV6(ip),":")
//	//ranges = strings.Split(convert2CompleteIpV6(ranges),":")
//}
