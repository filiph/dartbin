# dartbin

**DEPRECATION:** With the arrival of `dart2native`, the `dartbin` project is 
no longer needed. The [`dart2native` tool](https://dart.dev/tools/dart2native)
lets you easily create standalone executables for every supported platform.
Those executables start and run faster, and are smaller, than anything
that the `dartbin` project could have hoped for. I'm going to archive
this project, for posterity. -- Filip

---

A tool for packaging Dart programs into standalone executables.

## Prerequisites

* Go SDK installed

## Manual steps

1. Generate snapshot of your app (`dart --snapshot=file.snapshot file.dart`)
2. Run the source code generator (`dart bin/main.dart file.snapshot`)
3. Change to the `go_src` directory (`cd go_src`)
4. Compile the Go package 
   (`env GOPATH=/full/path/to/go_src go build -v`)
5. Grab the `go_src/go_src` file – that's your executable – and rename it 
   to your liking.

For other architectures (like Windows when you're running this on Mac, or 
vice versa), you'll need to:

* provide a matching executable 
  (`dart bin/main.dart --dart /full/path/to/dart.exe file.snapshot`) 
* run the Go compilation with the correct GOOS and GOARCH variables 
  (`env GOPATH=/full/path/to/go_src GOOS=windows GOARCH=386 go build -v`)

