package main

import (
	"regexp"
	"strconv"
)

func validIPAddress(IP string) string {
	if testIPv4(IP) {
		return "IPv4"
	}

	if testIPv6(IP) {
		return "IPv6"
	}

	return "Neither"
}

func testIPv4(IP string) bool {
	// regexp should be initialized only once when the app start
	// and keep it somewhere
	nums := regexp.MustCompile(`\.`).Split(IP, -1)

	if len(nums) != 4 {
		return false
	}

	for _, num := range nums {
		val, err := strconv.Atoi(num)
		if err != nil {
			return false
		}
		// no leading zeros
		if len(num) > 1 && num[0] == '0' {
			return false
		}
		if val < 0 || val > 255 {
			return false
		}
	}

	return true
}

func testIPv6(IP string) bool {
	// regexp should be initialized only once when the app start
	// and keep it somewhere
	nums := regexp.MustCompile(`:`).Split(IP, -1)

	if len(nums) > 8 {
		return false
	}

	if len(nums) < 3 { // least case "0::0"
		return false
	}

	omitFlag := false
	for i, num := range nums {
		if num == "" {
			if i != len(nums)-1 && i != 0 {
				omitFlag = true
				continue
			}
			return false // the first/last num should not be empty
		}
		val, err := strconv.ParseInt(num, 16, 32)
		if err != nil {
			return false
		}
		// no leading zeros
		if len(num) > 4 && num[0] == '0' {
			return false
		}
		if val < 0 || val > 0xFFFF {
			return false
		}
	}

	if !omitFlag && len(nums) != 8 {
		return false
	} else if omitFlag && len(nums) == 8 {
		return false
	}

	return true
}

func main() {
	println(validIPAddress("172.16.254.1")) // IPv4

	println(validIPAddress("2001:0db8:85a3:0:0:8A2E:0370:7334")) // IPv6

	println(validIPAddress("2001:0db8:85a3::8A2E:0370:7334")) // IPv6

	println(validIPAddress("2001:db8:85a3::8A2E:370:7334")) // IPv6

	println(validIPAddress("2001:0db8:85a3:0:0:8A2E:0370")) // Neither

	println(validIPAddress("2001:0db8:85a3:00000:0:8A2E:0370:7334")) // Neither

	println(validIPAddress("::1")) // Neither

	println(validIPAddress("::")) // Neither

	println(validIPAddress("1::1")) // IPv6

	println(validIPAddress("256.256.256.256")) // Neither
}
