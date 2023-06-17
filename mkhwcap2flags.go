//go:build ignore

// Generate hwcap2flags_*.go files using last copy of elf.h_include
// copied from https://git.musl-libc.org/cgit/musl/tree/include/elf.h
package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func addFile(goarch string, structlines []string) {
	w := new(bytes.Buffer)
	_, _ = fmt.Fprintf(w, "//go:build linux && %s\n", goarch)
	_, _ = fmt.Fprintln(w, "// Code generated using 'go generate'; DO NOT EDIT.")
	_, _ = fmt.Fprintln(w, "// hwcap2flags files contain the HWCAP flags for each architecture")
	_, _ = fmt.Fprintln(w, "//go:generate go run mkhwcap2flags.go")
	_, _ = fmt.Fprintln(w, "package processorfeatures")
	_, _ = fmt.Fprintln(w, "var MachineFeatures = []ProcessorFeature{")
	_, _ = fmt.Fprintln(w, strings.Join(structlines, "\n"))
	_, _ = fmt.Fprintln(w, "}")

	// gofmt result
	b := w.Bytes()
	b, err := format.Source(b)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(fmt.Sprintf("hwcap2flags_%s.go", goarch), b, 0666); err != nil {
		log.Fatalf("can't write output: %v\n", err)
	}

}

func main() {
	filename := "elf.h_include"
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	buf, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	if len(buf) == 0 {
		log.Printf("%s is empty", filename)
		return
	}
	records := bytes.Split(buf, []byte("\n"))
	// Build list of elf architectures
	var arches [][]byte
	regex := regexp.MustCompile("#define R_.*_NONE")
	archset := make(map[string]int)
	none := []byte("_NONE")
	for i, r := range records {
		if !regex.Match(r) {
			continue
		}
		before, _, b1 := bytes.Cut(r, none)
		if b1 {
			a := bytes.Replace(before, []byte("#define R_"), []byte{}, 1)
			arches = append(arches, a)
			archset[string(a)] = i
		}
	}
	if len(arches) == 0 {
		log.Println("no architecture found")
		return
	}
	log.Printf("%s", arches)
	// Lines of each architecture
	selected := make(map[string][]string, len(arches))
	for _, prefix := range arches {
		elfarch := string(prefix)
		for _, r := range records[archset[elfarch]:] {
			if len(r) == 0 {
				continue
			}
			if bytes.Contains(r, prefix) {
				selected[elfarch] = append(selected[elfarch], string(r))
				continue
			}
			break // Values are contiguous except blank lines
		}
		arch := strings.ToLower(elfarch)
		// Map name of architecture from elf to Go.
		goarch := arch
		switch goarch {
		case "390":
			goarch = "s390x"
		case "aarch64":
			goarch = "arm64"
		case "x86_64":
			goarch = "amd64"
		case "386", "arm", "mips", "ppc", "ppc64", "riscv", "sparc":
			// Name is identical
		default:
			log.Printf("%s is not supported in Go. Skipping %d definitions.\n", elfarch, len(selected[elfarch]))
			continue
		}
		var structlines []string
		pattern := fmt.Sprintf("#define R_%s_", elfarch)
		replacer := strings.NewReplacer("\t", " ", "\n", "", pattern, "")
		for i, s := range selected[elfarch] {
			// TODO Cope with complex lines
			if !strings.Contains(s, pattern) {
				continue
			}
			s = replacer.Replace(s)
			s = strings.TrimSpace(s)
			if s == "" {
				continue
			}
			r := s[:strings.Index(s, " ")]
			v := s[strings.LastIndex(s, " ")+1:]
			n, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				// Although it is not generated, IA64 has hex values: 0x...
				n, err = strconv.ParseInt(v[2:], 16, 64)
				if err != nil {
					for _, a := range arches {
						if strings.Contains(v, string(a)) {
							// Reference is currently the same index of record
							if strings.Contains(selected[string(a)][i], r) {
								v = selected[string(a)][i]
								v = v[1:] // Remove {
								v, _, b := strings.Cut(v, ",")
								if b {
									n, err = strconv.ParseInt(v, 10, 64)
								}
								break
							}
						}
					}
				}
				if err != nil {
					log.Printf("%s (%s): %v", v, s, err)
					continue
				}
			}
			structlines = append(structlines, fmt.Sprintf(`{%v, "", "%s", "%s"},`, n, strings.ToLower(r), s))
		}
		selected[elfarch] = structlines
		log.Printf("%s (%s): %d records parsed to %d", goarch, elfarch, len(selected[elfarch]), len(structlines))
		addFile(goarch, structlines)
		if goarch == "ppc64" {
			addFile("ppc64le", structlines)
		}
	}
}
