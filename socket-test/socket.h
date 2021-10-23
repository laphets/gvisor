#include <stdio.h>
#include <stdlib.h>
#include <errno.h>
#include <string.h>
#include <sys/types.h>

typedef enum
{
    OpCreate,
    OpBind,
    OpListen,
    OpAccept,
    OpConnect,
    OpShutdown,
    OpSetSockopt,
    OpGetSockopt,
    OpRelease,
    OpSend,
    OpRecv,
    OpGetSockName,
    OpGetPeerName,
} socket_op_t;

typedef struct __attribute__((__packed__)) socket_req
{
    int16_t fd : 16;
    int16_t size : 16;
    socket_op_t op : 8;
    char data[123];
} socket_req_t;

typedef struct __attribute__((__packed__)) socket_rsp
{
    int16_t fd : 16;
    int16_t size : 16;
    socket_op_t op : 8;
    int16_t result : 16;
    char data[121];
} socket_rsp_t;