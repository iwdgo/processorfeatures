// package setprocessor returns Go variables to set using bash export
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
			log.Fatalf("%v", err)
		}
		fmt.Printf("GOAMD64=%s", v)
		return
	case "arm", "arm64":
		v, err := processorfeatures.SetGOARMv()
		if err != nil {
			log.Fatalf("%v", err)
		}
		if v == "arm64" {
			fmt.Println("GOARCH=arm64 GOARM=")
			return
		}
		fmt.Printf("GOARCH=arm GOARM=%s", v)
		return
	default:
		fmt.Println("not implemented")
	}
}
