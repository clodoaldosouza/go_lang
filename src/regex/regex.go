package main

import (
	"fmt"
	"regexp"
)

var expression = "(?P<HEADER>#)(?P<ACCOUNT>\\d{4})(?P<PASS>\\w{8})(?P<PARAM>(?:A|B|C|D).(\\w{2}))(?P<CRC>.{2})(?P<DATE>.{10})(?P<TAIL>#)"
var msg = "#50003E0B28F2C3DB4A1644034395#"

type Received struct {
	account  string
	password string
	crc      string
	date     string
}

func main() {
	pattern := regexp.MustCompile(expression)
	matches := pattern.FindStringSubmatch(msg)

	received := Received{}
	received.account = matches[pattern.SubexpIndex("ACCOUNT")]
	received.password = matches[pattern.SubexpIndex("PASS")]
	received.crc = matches[pattern.SubexpIndex("CRC")]
	received.date = matches[pattern.SubexpIndex("DATE")]

	receivedMode2 := Received{
		matches[pattern.SubexpIndex("ACCOUNT")],
		matches[pattern.SubexpIndex("PASS")],
		matches[pattern.SubexpIndex("CRC")],
		matches[pattern.SubexpIndex("DATE")]}

	fmt.Println(received)
	fmt.Println(receivedMode2)
}
