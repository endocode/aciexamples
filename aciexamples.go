package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func readDir(dir string) {
	fmt.Println("Dir: " + dir + ":")
	d, err := os.Open(dir)
	if d == nil {
		fmt.Println("cannot read directory "+dir+": ", err)
		return
	}
	files, err := d.Readdirnames(-1)
	if files == nil {
		fmt.Println("cannot get files in directory "+dir+": ", err)
		return
	}
	for j := 0; j < len(files); j++ {
		fmt.Println(files[j])
	}
	d.Close()

}

func readFile(file string) {
	fmt.Println("File: " + file + ":")
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

	wd, err := os.Getwd()
	if err == nil {
		fmt.Println("Working directory: " + wd)
	} else {
		fmt.Printf("Cannot get working directory: %v", err)
	}

	fmt.Println("Environment:\n" + strings.Join(os.Environ(), "\n") + "\n")

	readDir("/mnt/vol1")
	readDir("/mnt/vol2")
	readDir("/mnt/vol3")
	readDir("/mnt/vol4")
	readDir("/mnt/vol5")
}
