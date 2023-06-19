package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/iwdgo/processorfeatures"
	"io"
	"log"
	"math/bits"
	"os"
	"runtime"
)

const (
	AT_NULL = 0

	AT_IGNORE = 1
	AT_EXECFD = 2
	AT_PHDR   = 3
	AT_PHENT  = 4
	AT_PHNUM  = 5
	AT_PAGESZ = 6
	AT_BASE   = 7
	AT_FLAGS  = 8
	AT_ENTRY  = 9
	AT_NOTELF = 10
	AT_UID    = 11
	AT_EUID   = 12
	AT_GID    = 13
	AT_EGID   = 14
	AT_CLKTCK = 17

	AT_PLATFORM = 15
	AT_HWCAP    = 16

	AT_FPUCW = 18

	AT_DCACHEBSIZE = 19
	AT_ICACHEBSIZE = 20
	AT_UCACHEBSIZE = 21

	AT_IGNOREPPC = 22

	AT_SECURE = 23

	AT_BASE_PLATFORM = 24

	AT_RANDOM = 25

	AT_HWCAP2 = 26

	AT_EXECFN = 31

	AT_SYSINFO      = 32
	AT_SYSINFO_EHDR = 33

	AT_L1I_CACHESHAPE = 34
	AT_L1D_CACHESHAPE = 35
	AT_L2_CACHESHAPE  = 36
	AT_L3_CACHESHAPE  = 37

	AT_L1I_CACHESIZE     = 40
	AT_L1I_CACHEGEOMETRY = 41
	AT_L1D_CACHESIZE     = 42
	AT_L1D_CACHEGEOMETRY = 43
	AT_L2_CACHESIZE      = 44
	AT_L2_CACHEGEOMETRY  = 45
	AT_L3_CACHESIZE      = 46
	AT_L3_CACHEGEOMETRY  = 47

	AT_MINSIGSTKSZ = 51
)

func main() {
	log.Println("reading /proc/self/auxv")
	filename := "/proc/self/auxv"
	// filename := filepath.Join("..", "cpufeatures", "auxvbin")
	buf, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("%s", hex.Dump(buf))
	/*
		f, err := os.Create("/tmp/auxvbin")
		defer f.Close()
		if err != nil {
			log.Printf("%v", err)
		}
		n, err := f.Write(buf)
		if err != nil {
			log.Printf("%v", err)
		}

	*/
	log.Printf("%s contains %v bytes (%d uint64)", filename, len(buf), len(buf)/8)
	if runtime.GOARCH == "arm" {
		log.Println("Skipping on 32 bits")
		return
	}
	if runtime.GOOS == "s390x" {
		log.Printf("Big endian. Skipping")
		return
	}
	var w, v uint64
	var hwcap2 uint32
	// var ptr *string
	reader := bytes.NewReader(buf)
	for err != io.EOF {
		err = binary.Read(reader, binary.LittleEndian, &w)
		if err != nil {
			fmt.Println(err)
			continue
		}
		switch w {
		case AT_NULL:
			err = binary.Read(reader, binary.LittleEndian, &v)
			// log.Println("last byte read")
			break
		case AT_PLATFORM:
			err = binary.Read(reader, binary.LittleEndian, &v)
			/*
				// TODO Invalid type **string
				err = binary.Read(reader, binary.LittleEndian, &ptr)
				if err != nil {
					log.Println(err)
					break
				}

			*/
			log.Printf("AT_PLATFORM (%d) [%x] ", w, w)
			// https://github.com/golang/go/issues/41205
			// go vet -unsafeptr=false
			// stringPtr := (*string)(unsafe.Pointer(uintptr(v)))
			// platform := *stringPtr
			// log.Println(*ptr)
		case AT_HWCAP:
			err = binary.Read(reader, binary.LittleEndian, &v)
			// First 32 bit word is significant
			log.Printf("AT_HWCAP (%d) [%x] = (%d) %064b", w, w, v, v)
		case AT_HWCAP2:
			err = binary.Read(reader, binary.LittleEndian, &hwcap2)
			// Keeping only the first 32 bits
			s := fmt.Sprintf("%032b", hwcap2)
			log.Printf("%s has %v", s, bits.LeadingZeros32(hwcap2))
			log.Printf("AT_HWCAP2 (%d) [%x]", hwcap2, hwcap2)
			var available []string
			var unavailable []string
			i := 0
			l := len(processorfeatures.MachineFeatures)
			for {
				b := s[i]
				if string(b) == "0" {
					unavailable = append(unavailable, processorfeatures.MachineFeatures[i].S)
				} else {
					available = append(available, processorfeatures.MachineFeatures[i].S)
				}
				i++
				if i >= l {
					sl := s[i:] // If sl contains ones, a feature is lost
					log.Printf("%d machinefeatures, ignored %s", l, sl)
					break
				}
				if i >= len(s) {
					log.Printf("%s machine features for 32 bits (ARM,...)", len(s))
					break
				}
			}
			log.Printf("available %v", available)
			log.Printf("unavailable %v", unavailable)
		default:
			err = binary.Read(reader, binary.LittleEndian, &v)
			// log.Printf("%d [%x] = %b", w, w, v)
		}
	}
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
}
