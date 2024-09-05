package user

import (
	"fmt"
	"strconv"
)
func parseParamIDtoInt(id string) int {
	parsedID, err := strconv.ParseInt(id, 10, 64) // 10 base, 64 bits

	if err != nil {
		fmt.Println(err)
		return 0
	}

	return int(parsedID)
}

func FormattedIPAddress(IpAddress string) string {
	if IpAddress == "::1" {
		return "127.0.0.1"
	}

	return IpAddress
}