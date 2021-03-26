package main

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
)

func main() {
	// fmt.Printf("Version: %s\n", runtime.Version())
	// fmt.Printf("Kernel: %s\n", runtime.GOOS)
	// fmt.Printf("GOARCH: %s\n", runtime.GOARCH)

	// Hostname, err := os.Hostname()
	// handError(err)
	// fmt.Printf("Hostname: %s\n", Hostname)
	// u, er := user.Current()
	// handError(er)
	// fmt.Printf("User: %s\n", u.Username)
	// fmt.Print(host.Info())
	fmt.Print(cpu.Info())
}

func handError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
