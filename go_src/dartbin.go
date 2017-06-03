package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func run() int {
	dir, err := ioutil.TempDir("", "dartbin")
	if err != nil {
		log.Fatal(err)
		return 1
	}

	defer os.RemoveAll(dir) // clean up

	// TODO: keep track of the temp dir and next time, don't copy

	// dartvmbytes := [...]byte{1, 3, 3, 4, 5, 5, 6, 4, 4}

	dartfn := filepath.Join(dir, "dart")
	if err := ioutil.WriteFile(dartfn, dartvmbytes[:], 0751); err != nil {
		log.Fatal(err)
		return 1
	}

	exefn := filepath.Join(dir, "exe.snapshot")
	if err := ioutil.WriteFile(exefn, snapshotbytes[:], 0666); err != nil {
		log.Fatal(err)
		return 1
	}

	params := make([]string, 0)
	params = append(params, exefn)
	for i := 1; i < len(os.Args); i++ {
		params = append(params, os.Args[i])
	}
	cmd := exec.Command(dartfn, params...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	cmd.Run()

	var code int
	if cmd.ProcessState.Success() {
		code = 0
	} else {
		code = 1
	}
	return code
}

func main() {
	os.Exit(run())
}
