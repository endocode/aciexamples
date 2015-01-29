package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func readFile(file string) {
	fmt.Println(file + ":")
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Cannot read file "+file+": ", err)
		return
	}
	fmt.Print(string(dat))
}

func main() {
	fmt.Println("Hello! Who am I? Where am I?")
	fmt.Println("Args: " + strings.Join(os.Args, " "))

	readFile("/proc/mounts")
}
