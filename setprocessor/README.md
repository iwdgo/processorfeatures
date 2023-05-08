### How to

```bash
          go install github.com/iwdgo/processorfeatures/setprocessor@latest
          $(go env GOPATH)/bin/setprocessor # display value
          go env -w $(setprocessor)         # sets go environment variables
          export $(setprocessor)            # sets shell variables
```

Examples:
 
- amd64 : `GOAMD64=v3`
- arm64 : `GOARCH=arm64 GOARM=`
- arm   : `GOARCH=arm GOARM=7` or default `GOARCH=arm GOARM=6`
- or `not implemented`
