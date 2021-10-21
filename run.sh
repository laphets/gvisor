#!/bin/bash
chmod 777 /tmp/gvisor.sock
CGO_ENABLED=0 GO111MODULE=on go build -o /tmp/runsc gvisor.dev/gvisor/runsc && docker run -it -v $(pwd):/gvsior -v /tmp:/tmp --rm --runtime=gvisor-dpdk laphets/gvisor-test /bin/bash