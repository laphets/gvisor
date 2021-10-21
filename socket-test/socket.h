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
} socket_op_t;

typedef struct socket_req
{
    int16_t fd;
    int16_t size;
    socket_op_t op : 8;
    char data[123];
} socket_req_t;

typedef struct socket_rsp
{
    int16_t fd;
    int16_t size;
    socket_op_t op;
    int16_t result;
    char data[121];
} socket_rsp_t;