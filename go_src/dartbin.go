package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

func run() int {
	dir, temperr := ioutil.TempDir("", "dartbin")
	if temperr != nil {
		log.Fatal(temperr)
		return 1
	}

	defer os.RemoveAll(dir) // clean up

	// TODO: keep track of the temp dir and next time, don't copy

	// dartvmbytes := [...]byte{1, 3, 3, 4, 5, 5, 6, 4, 4}

	dartfn := filepath.Join(dir, dartexename)
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
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	runerr := cmd.Run()

	if runerr == nil {
		// The returned error is nil if the command runs,
		// has no problems copying stdin, stdout, and stderr,
		// and exits with a zero exit status.
		return 0
	}

	if exitError, ok := runerr.(*exec.ExitError); ok {
		// Use cmd.ProcessState.Sys() to get the actual exit code.
		ws := exitError.Sys().(syscall.WaitStatus)
		return ws.ExitStatus()
	} else {
		log.Fatal(runerr)
		return 1
	}
}

func main() {
	os.Exit(run())
}
