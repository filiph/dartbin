// Copyright (c) 2017, filiph. All rights reserved. Use of this source code
// is governed by a BSD-style license that can be found in the LICENSE file.

import 'dart:io';

import 'package:args/args.dart';
import 'package:path/path.dart' as path;

main(List<String> arguments) {
  var argParser = new ArgParser(allowTrailingOptions: true)
    ..addOption("dart",
        defaultsTo: Platform.resolvedExecutable,
        help: "Full path to the dart executable.");
  var args = argParser.parse(arguments);

  if (args.rest.length != 1) {
    print("Program will bundle Dart snapshot into a standalone executable.\n");
    print("usage: dartbin [options] path/to/file.snapshot\n");
    print(argParser.usage);
    exitCode = 2;
    return;
  }

  var dir = new Directory("go_src");

  var dartExePath = args['dart'];
  var dartVmBytesFilename = path.join(dir.path, "dartvmbytes.go");

  try {
    buildGoFile('dartvmbytes', dartExePath, dartVmBytesFilename);
  } on FileSystemException catch (e) {
    print("Couldn't build $dartVmBytesFilename: $e");
    exitCode = 1;
    return;
  }

  // TODO: add to README: dart --snapshot=test.snapshot bin/main.dart
  var snapshotPath = args.rest.single;
  var snapshotBytesFilename = path.join(dir.path, "snapshotbytes.go");

  try {
    buildGoFile('snapshotbytes', snapshotPath, snapshotBytesFilename);
  } on FileSystemException catch (e) {
    print("Couldn't build $snapshotBytesFilename: $e");
    exitCode = 1;
    return;
  }

  // env GOPATH=/Users/filiph/dev/dartbin/go_src GOOS=windows GOARCH=386 go build -v
  // TODO: env GOOS=windows GOARCH=386 go build -v
}

void buildGoFile(String variableName, String binaryPath, String goSourcePath) {
  var file = new File(binaryPath);
  var outFile = new File(goSourcePath);

  var bytes = file.readAsBytesSync();
  print("Reading of $file done: ${bytes.length} bytes");

  var out = new StringBuffer();
  out.writeln('package main\n');
  out.write("var $variableName = [...]byte{");
  for (int i = 0; i < bytes.length - 1; i++) {
    out.write(bytes[i].toString());
    out.write(', ');
    if ((i + 1) % 10 == 0) {
      out.write('\n');
    }
  }
  // Last byte (so that we don't end with a comma)
  out.write(bytes.last.toString());
  out.writeln('}');

  outFile.writeAsStringSync(out.toString());
}
