//go:build ignore

// Generate hwcapflags_*.go files using last copy of elf.h_include
package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// Reads https://git.musl-libc.org/cgit/musl/tree/include/elf.h
func featuresToFile(f *bytes.Buffer, prefix string, rawFeatures string) *bytes.Buffer {
	replacer := strings.NewReplacer("\t", " ", "\n", "")
	out := replacer.Replace(rawFeatures)
	features := strings.Split(out, "#define ")
	for i, s := range features {
		if s == "" {
			continue
		}
		index := strings.Index(s, " ")
		if index == -1 {
			log.Printf("%s (%d) has no blank space", s, i)
			continue
		}
		if index <= len(prefix) {
			log.Printf("blank space at %d is before prefix %s (%s)", index, prefix, s)
			continue
		}
		r := s[len(prefix):index]
		v := s[strings.LastIndex(s, " ")+1:]
		n, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			// Although it is not generated, IA64 has hex values: 0x...
			n, err = strconv.ParseInt(v[2:], 16, 64)
			if err != nil {
				// TODO PPC64 re-uses PPC values
				log.Printf("%s (%s): %v", v, s, err)
				continue
			}
		}
		_, err = fmt.Fprintln(f, fmt.Sprintf(`{%v, "", "%s", "%s"},`, n, strings.ToLower(r), s))
		if err != nil {
			log.Fatal(err)
		}
	}
	return f
}

func main() {
	f, err := os.Open("elf.h_include")
	if err != nil {
		log.Fatal(err)
	}
	buf, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	records := bytes.Split(buf, []byte("\n"))
	// Build list of prefix
	var arches [][]byte
	archset := make(map[string]int) // R_*_NONE index
	none := []byte("_NONE")
	for i, r := range records {
		before, _, b1 := bytes.Cut(r, none)
		if b1 {
			_, after, b2 := bytes.Cut(before, []byte(" R_"))
			if len(bytes.TrimSpace(after)) == 0 {
				continue
			}
			if b2 {
				arches = append(arches, after)
				archset[string(after)] = i
			}
		}
	}
	log.Printf("%s", arches)
	for _, prefix := range arches {
		selected := make([]string, 512)
		for _, r := range records[archset[string(prefix)]:] {
			if bytes.Contains(r, prefix) {
				selected = append(selected, string(r))
				continue
			}
			break // Values are contiguous
		}
		arch := strings.ToLower(string(prefix))
		w := new(bytes.Buffer)
		// Some elf arches are names differently in Go
		goarch := arch
		switch goarch {
		case "390":
			goarch = "s390x"
		case "aarch64":
			goarch = "arm64"
		case "x86_64":
			goarch = "amd64"
		case "386", "arm", "ppc", "ppc64", "riscv", "sparc":
			// Known architectures
		default:
			fmt.Printf("Architecture %s is not supported in Go. Skipping %d definitions.\n", goarch, len(selected))
			continue
		}
		log.Printf("parsing %d records for prefix %s (%s)", len(selected), string(prefix), goarch)
		_, _ = fmt.Fprintf(w, "//go:build %s\n", goarch)
		_, _ = fmt.Fprintf(w, "// Code generated using 'go generate'; DO NOT EDIT.\n")
		_, _ = fmt.Fprintln(w, "// package buildref contains the the HWCAP flags for each architecture")
		_, _ = fmt.Fprintln(w, "//go:generate go run mkhwcapflags.go")
		_, _ = fmt.Fprintln(w, "package hwcapflags")
		_, _ = fmt.Fprintln(w, `import "github.com/iwdgo/processorfeatures"`)
		_, _ = fmt.Fprintf(w, "var ProcessorFeatures = []processorfeatures.ProcessorFeature{\n")
		w = featuresToFile(w, string(prefix), strings.Join(selected, "\n"))
		_, _ = fmt.Fprintf(w, "}")

		// gofmt result
		b := w.Bytes()
		b, err = format.Source(b)
		if err != nil {
			panic(err)
		}

		if err := os.WriteFile(fmt.Sprintf("hwcapflags_%s.go", goarch), b, 0666); err != nil {
			log.Fatalf("can't write output: %v\n", err)
		}
	}
}
