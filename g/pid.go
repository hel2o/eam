package g

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Init()  {
	version := flag.Bool("v", false, "show version")
	flag.Parse()

	if *version {
		fmt.Println(VERSION)
		os.Exit(0)
	}
	runPid()
}

func runPid() {
	var filename = "var/app.pid"
	var f *os.File
	var err1 error

	if checkFileIsExist(filename) {
		f, err1 = os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	} else {
		f, err1 = os.Create(filename)
	}
	check(err1)
	_, err1 = io.WriteString(f, strconv.Itoa(os.Getpid()))
	check(err1)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
