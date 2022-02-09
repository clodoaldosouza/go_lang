package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {
	readOnly()
	println("")
	OpenWrite()
	println("")
	anotherWayOpenFile()
}

func OpenWrite() {
	pwd, _ := os.Getwd()
	file, err := os.OpenFile(pwd+"/src/open_file/log.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		println("Error:", err)
		file.Close()
		return
	}
	file.WriteString(time.Now().Format("02/01/2006 - 15:04:05 = ") + "Append File..." + "\n")
	file.Close()
}

func anotherWayOpenFile() {
	pwd, _ := os.Getwd()
	file, err := ioutil.ReadFile(pwd + "/src/open_file/log.log") // its self close the file
	if err != nil {
		println("Error:", err)
		return
	}
	println(string(file)) // convert bytes to string

}

func readOnly() {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/src/open_file/myFile.txt")
	if err != nil {
		println("Error:", err)
		file.Close()
		return
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		fmt.Println(strings.TrimSpace(line))
		if err == io.EOF {
			file.Close()
			break
		}
	}
	file.Close()
}
