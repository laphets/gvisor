# gVisor

This branch is a synthetic branch, containing only Go sources, that is
compatible with standard Go tools. See the master branch for authoritative
sources and tests.

## Run
See `run.sh` for build and use `runsc` in docker.

The `runs.sh` will build `runsc` to `/tmp/runsc`, and use `laphets/gviosr-test` as the container image for testing.

Add the following configs to `/etc/docker/daemon.json` to register `runsc` runtime,
```json
"runtimes": {
    "gvisor-dpdk": {
        "path": "/tmp/runsc",
        "runtimeArgs": [
            "--network=dpdk",
            "--debug-log",
            "/tmp/gvisor/logs/runsc.log.%TEST%.%TIMESTAMP%.%COMMAND%",
            "--net-raw",
            "--debug",
            "--strace",
            "--log-packets"
        ]
    },
    "gvisor": {
        "path": "/tmp/runsc",
        "runtimeArgs": [
            "--debug-log",
            "/tmp/gvisor/logs/runsc.log.%TEST%.%TIMESTAMP%.%COMMAND%",
            "--net-raw",
            "--debug",
            "--strace",
            "--log-packets"
        ]
    },
}
```

Then restart docker daemon by `systemctl restart docker`.