package gb

import (
	"fmt"
)

func Init(message string) string {
	lenth := len(message)
	if lenth < 1 {
		//fmt.Println("未接收到数据")
		return "40408c0d013011380f1c0914010000000000ffff00000000080003132323"
	}
	var promoter = message[:4]
	var businessSerialNumber = message[4:8]
	var agreementVersionNumber = message[8:12]
	var timeTag = message[12:24]
	var sourceAddress = message[24:36]
	var destinationAddress = message[36:48]
	var applicationDataUnitLength = message[48:52]
	var command = message[52:54]
	var applicationDataUnit = message[54 : lenth-6]
	var checkSum = message[lenth-6 : lenth-4]
	var terminator = message[lenth-4 : lenth]
	var checkString = businessSerialNumber + agreementVersionNumber + timeTag + sourceAddress + destinationAddress + applicationDataUnitLength
	var check = CheckSum(checkString + command + applicationDataUnit)
	//fmt.Println("check is ", check)
	if check != checkSum {
		fmt.Println("校验和不正确")
	}
	//fmt.Println("checkString is ", checkString+"03"+applicationDataUnit)
	var checkNum = CheckSum(checkString + "03" + applicationDataUnit)
	//fmt.Println("checknum is ", checkNum)
	var sendData = promoter + checkString + "03" + checkNum + terminator
	return sendData
}
