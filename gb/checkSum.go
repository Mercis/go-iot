package gb

import (
	"fmt"
	"strconv"
	"strings"
)

//校验和计算
func CheckSum(data string) string {
	var parse = 0
	for i := 0; i < len(data); i++ {
		if i%2 == 0 {
			d := data[i : i+2]
			num, err := strconv.ParseInt(d, 16, 32)
			if err != nil {
				continue
			}
			//int64 转 int
			strInt64 := strconv.FormatInt(num, 10)
			id16, _ := strconv.Atoi(strInt64)
			//舍去8位(1字节)以上的进位位后
			parse += id16
			parse = parse & 255
		}
	}

	checkSum := toHex(parse)
	if len(checkSum) == 1 {
		checkSum = "0" + checkSum
	}
	return checkSum
}

//10进制转16进制
func toHex(ten int) string {
	m := 0
	hex := make([]int, 0)
	for {
		m = ten % 16
		ten = ten / 16
		if ten == 0 {
			hex = append(hex, m)
			break
		}
		hex = append(hex, m)
	}
	hexStr := []string{}
	for i := len(hex) - 1; i >= 0; i-- {
		if hex[i] >= 10 {
			hexStr = append(hexStr, fmt.Sprintf("%c", 'A'+hex[i]-10))
		} else {
			hexStr = append(hexStr, fmt.Sprintf("%d", hex[i]))
		}
	}
	return strings.Join(hexStr, "")
}
