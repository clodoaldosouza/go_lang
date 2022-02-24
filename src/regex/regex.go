package main

import (
	"./models"
	"fmt"
	"regexp"
)

var expression = "(?P<HEADER>#)(?P<ACCOUNT>\\d{4})(?P<PASS>\\w{8})(?P<PARAM>(?:A|B|C|D).(\\w{2}))(?P<CRC>.{2})(?P<DATE>.{10})(?P<TAIL>#)"
var msg = "#50003E0B28F2C3DB4A1644034395#"

func main() {
	pattern := regexp.MustCompile(expression)
	matches := pattern.FindStringSubmatch(msg)

	received1 := received.Received{}
	received1.Account = matches[pattern.SubexpIndex("ACCOUNT")]
	received1.Password = matches[pattern.SubexpIndex("PASS")]
	received1.Crc = matches[pattern.SubexpIndex("CRC")]
	received1.Date = matches[pattern.SubexpIndex("DATE")]

	received2 := received.Received{
		matches[pattern.SubexpIndex("ACCOUNT")],
		matches[pattern.SubexpIndex("PASS")],
		matches[pattern.SubexpIndex("CRC")],
		matches[pattern.SubexpIndex("DATE")]}

	fmt.Println(received1)
	fmt.Println(received2)
}
