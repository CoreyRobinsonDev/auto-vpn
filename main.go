package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"
)

func main() {
	pathPtr := flag.String("p", "", "path to listen on")
	cleanPtr := flag.Bool("clean", false, "remove .ovpn files from path")
	versionPtr := flag.Bool("v", false, "version")
	flag.Parse()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	if *versionPtr {
		fmt.Println("auto-vpn v1.0.0")
		os.Exit(0)
	}

	if *cleanPtr {
		entries := Unwrap(os.ReadDir(*pathPtr))
		for _, entry := range entries {
			if strings.Contains(entry.Name(), ".ovpn") {
				Expect(os.Remove(*pathPtr + "/" + entry.Name()))
			}
		}
	}

	fmt.Printf("\r- Listening on path \x1b[32m%s\x1b[0m for .ovpn file",*pathPtr)
	for {
		select {
		case <- sigChan:
			fmt.Println()
			os.Exit(0)		
		default:
			entries := Unwrap(os.ReadDir(*pathPtr))
			for _, entry := range entries {
				if strings.Contains(entry.Name(), ".ovpn") {
					fmt.Println()
					sudoExe := Unwrap(exec.LookPath("sudo"))
					syscall.Exec(sudoExe, []string{"sudo", "openvpn", *pathPtr + "/" + entry.Name()}, os.Environ())
				}
			}
			time.Sleep(time.Millisecond*200)
			fmt.Printf("\r\\ Listening on path \x1b[32m%s\x1b[0m for .ovpn file",*pathPtr)
			time.Sleep(time.Millisecond*200)
			fmt.Printf("\r| Listening on path \x1b[32m%s\x1b[0m for .ovpn file",*pathPtr)
			time.Sleep(time.Millisecond*200)
			fmt.Printf("\r/ Listening on path \x1b[32m%s\x1b[0m for .ovpn file",*pathPtr)
			time.Sleep(time.Millisecond*200)
			fmt.Printf("\r- Listening on path \x1b[32m%s\x1b[0m for .ovpn file",*pathPtr)
		}
	}
}

func Expect(err error) {
	_, file, line, ok := runtime.Caller(1)
	progName := "auto-vpn"
	if !ok { 
		if err != nil {
			fmt.Fprintf(
				os.Stderr, 
				"\x1b[2m%s\x1b[0m %v\n",
				progName,
				err,
			)
			os.Exit(1)
		}
	} else {
		if err != nil {
			fmt.Fprintf(
				os.Stderr, 
				"\x1b[2m%s\x1b[0m %v\n",
				fmt.Sprintf("%s(%s:%d): ",progName, file,line),
				err,
			)
			os.Exit(1)
		}
	}
}

func Unwrap[T any](result T, err error) T {
	_, file, line, ok := runtime.Caller(1)
	progName := "auto-vpn"
	if !ok {
		if err != nil {
			fmt.Fprintf(
				os.Stderr, 
				"\x1b[2m%s\x1b[0m %v\n",
				progName,
				err,
			)
			os.Exit(1)
		}
	} else {
		if err != nil {
			fmt.Fprintf(
				os.Stderr, 
				"\x1b[2m%s\x1b[0m %v\n",
				fmt.Sprintf("%s(%s:%d): ",progName,file,line),
				err,
			)
			os.Exit(1)
		}
	}
	return result
}
