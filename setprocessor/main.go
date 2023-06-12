// package setprocessor returns processor variables in the format of go env
package main

import (
	"fmt"
	"github.com/iwdgo/processorfeatures"
	"log"
	"runtime"
)

func main() {
	switch runtime.GOARCH {
	case "amd64":
		v, err := processorfeatures.SetGOAMD64v()
		if err != nil {
			log.Fatalf("Set GOAMD64 failed with %v", err)
		}
		fmt.Printf("GOAMD64=%s\n", v)
	case "arm", "arm64":
		v, err := processorfeatures.SetGOARMv()
		if err != nil {
			log.Fatalf("Set GOARM or GOARCH failed with %v", err)
		}
		if v == "arm64" {
			fmt.Println("GOARCH=arm64 GOARM=")
			return
		}
		fmt.Printf("GOARCH=arm GOARM=%s\n", v)
	case "386", "mips", "mips64", "mips64le", "ppc64", "ppc64le":
		fmt.Println("not implemented")
	default:
		fmt.Println("")
	}
}
