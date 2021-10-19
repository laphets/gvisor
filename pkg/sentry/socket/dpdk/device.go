package dpdk

import "gvisor.dev/gvisor/pkg/sentry/device"

var socketDevice = device.NewAnonDevice()
