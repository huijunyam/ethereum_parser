package helper

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvertIntToHexString(num int64) string {
	return fmt.Sprintf("0x%x", num)
}

func ConvertHexStringToInt(hex string) (int64, error) {
	result := isDecimal(hex)
	decimalVal, err := strconv.ParseUint(result, 16, 64)
	if err != nil {
		return int64(decimalVal), err
	}
	return int64(decimalVal), nil
}

func isDecimal(hexString string) string {
	// replace 0x or 0X with empty String
	number := strings.Replace(hexString, "0x", "", -1)
	number = strings.Replace(number, "0X", "", -1)

	return number
}
