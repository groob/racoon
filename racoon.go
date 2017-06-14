package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	var (
		flManifest = flag.String("manifest", "dep_bootstrap", "manifest to loop through")
		flInterval = flag.Duration("interval", 2*time.Second, "interval between checks")
	)
	flag.Parse()

	ticker := time.NewTicker(*flInterval).C
	done := make(chan bool)
	for {
		if err := runMunkiID(*flManifest, done); err != nil {
			log.Println(err)
			os.Exit(1)
		}
		<-done
		<-ticker
	}
}

func runMunkiID(id string, done chan<- bool) error {
	idArg := fmt.Sprintf("--id=%s", id)
	checkArgs := []string{"--munkipkgsonly", "--checkonly", idArg}
	installArgs := []string{"--munkipkgsonly", "--installonly", idArg}
	cmdName := "/usr/local/munki/managedsoftwareupdate"

	runWithArgs := func(args ...string) error {
		cmd := exec.Command(cmdName, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}

	if err := runWithArgs(checkArgs...); err != nil {
		return err
	}

	if err := runWithArgs(installArgs...); err != nil {
		return err
	}

	go func() { done <- true }()
	return nil
}
