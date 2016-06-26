package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	usr, err := user.Current()
	check(err)

	log.Printf("Platform: %v", runtime.GOOS)
	log.Printf("Home dir variable: %v", os.Getenv("HOME"))

	log.Printf("Current home dir: %v", usr.HomeDir)
	configf := filepath.Join(usr.HomeDir, ".ssh", "config")
	log.Printf("ssh config is at: %v", configf)

	f, err := ioutil.ReadFile(configf)
	check(err)
	log.Printf("File: %v", string(f))

	data, err := os.Open(configf)
	check(err)
	defer data.Close()
	log.Printf("Data: %v", data)

}
