#!/bin/bash
bazel build //socket-test:server && ./bazel-bin/socket-test/server
