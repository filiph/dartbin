# dartbin

## Prerequisites

* Go SDK installed

## Manual steps

1. Generate snapshot of your app (`dart --snapshot=file.snapshot file.dart`)
2. Run the source code generator (`dart bin/main.dart file.snapshot`)
3. Change to the `go_src` directory (`cd go_src`)
4. Compile the Go package 
   (`env GOPATH=/full/path/to/go_src go build -v`)
5. Grab the `go_src/go_src` file – that's your executable – and rename it.

For other architectures (like Windows when you're running this on Mac, or 
vice versa), you'll need to provide a matching executable 
(`dart bin/main.dart --dart /full/path/to/dart.exe file.snapshot`) 
and you'll need to run the Go compilation with the correct GOOS and GOARCH 
(`env GOPATH=/full/path/to/go_src GOOS=windows GOARCH=386 go build -v`)

