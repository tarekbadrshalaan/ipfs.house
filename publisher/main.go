package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/jessevdk/go-flags"
)

// Opts ...
type Opts struct {
	Repo string `short:"r" long:"repo" description:"repository git url" required:"true"`
	Dir  string `short:"d" long:"directory" description:"directory output name" required:"true"`
}

// go build . && ./publisher -r https://github.com/tarekbadrshalaan/gopkg.git -d myrepo
func main() {
	// parse flags
	opt := &Opts{}
	_, err := flags.ParseArgs(opt, os.Args)

	// git clone
	// cmd := exec.Command("git", "clone", "--bare", "https://github.com/tarekbadrshalaan/gopkg.git", "myrepo")
	cmd := exec.Command("git", "clone", "--bare", opt.Repo, opt.Dir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	// git update-server-info
	cmd2 := exec.Command("git", "update-server-info")
	cmd2.Dir = opt.Dir
	cmd2.Stdout = os.Stdout
	cmd2.Stderr = os.Stderr
	err = cmd2.Run()
	if err != nil {
		log.Fatalf("cmd2.Run() failed with %s\n", err)
	}
	// add to ipfs
	// -r Add directory paths recursively.
	// -w to Wrap files with a directory object.
	// -H to Include files that are hidden.
	// -Q Write only final hash.
	cmd3 := exec.Command("ipfs", "add", "-Q", "-H", "-w", "-r", ".")
	cmd3.Dir = opt.Dir
	hashout, err := cmd3.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd3.Run() failed with %s\n", err)
	}
	hashoutstr := strings.TrimSpace(string(hashout))
	fmt.Printf("\nipfs get %s/%v\ngit clone http://localhost:8080/ipfs/%s/%v", hashoutstr, opt.Dir, hashoutstr, opt.Dir)

	// to get the directory use
	// ipfs get <<hash>>/<<opt.Dir>>

	// to execute git clone
	// git clone http://localhost:8080/ipfs/<<hash>>/<<opt.Dir>>
}
