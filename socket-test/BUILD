config_setting(
    name = "custom",
    values = {
        "define": "fstack=true"
    }
)

cc_binary(
    name = "server",
    srcs = ["socket-server.c", "socket.h"],
    linkopts = ["-lpthread"],
    defines = select({
        ":custom": ["FSTACK"],
        "//conditions:default": [],
    })
)