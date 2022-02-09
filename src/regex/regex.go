package main

import (
	"fmt"
	"regexp"
	"souza.com/index/src/regex/models"
)

var expression = "(?P<HEADER>#)(?P<ACCOUNT>\\d{4})(?P<PASS>\\w{8})(?P<PARAM>(?:A|B|C|D).(\\w{2}))(?P<CRC>.{2})(?P<DATE>.{10})(?P<TAIL>#)"
var msg = "#50003E0B28F2C3DB4A1644034395#"

func main() {
	pattern := regexp.MustCompile(expression)
	matches := pattern.FindStringSubmatch(msg)

	received := models.Received{}
	received.Account = matches[pattern.SubexpIndex("ACCOUNT")]
	received.Password = matches[pattern.SubexpIndex("PASS")]
	received.Crc = matches[pattern.SubexpIndex("CRC")]
	received.Date = matches[pattern.SubexpIndex("DATE")]

	receivedMode2 := models.Received{
		matches[pattern.SubexpIndex("ACCOUNT")],
		matches[pattern.SubexpIndex("PASS")],
		matches[pattern.SubexpIndex("CRC")],
		matches[pattern.SubexpIndex("DATE")]}

	fmt.Println(received)
	fmt.Println(receivedMode2)
}
