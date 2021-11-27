#!/bin/bash
bazel build //socket-test:server --define fstack=true && ./bazel-bin/socket-test/server