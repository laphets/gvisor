#include <stdio.h>
#include <stdlib.h>
#include <errno.h>
#include <string.h>
#include <sys/types.h>
#include <sys/epoll.h>
#include <fcntl.h>
#include <sys/socket.h>
#include <sys/un.h>
#include <errno.h>
#include <pthread.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <linux/vm_sockets.h>
#include <sys/time.h>
#include <stdint.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <fcntl.h>  // for open
#include <unistd.h> // for close

#include "socket.h"

#ifdef FSTACK
#include "ff_socket.h"
#include "ff_config.h"
#include "ff_api.h"
#include "ff_epoll.h"

#define ss_fcntl ff_fcntl
#define ss_socket ff_socket
#define ss_bind ff_bind
#define ss_listen ff_listen
#define ss_connect ff_connect
#define ss_accept ff_accept
#define ss_close ff_close
#define ss_shutdown ff_shutdown
#define ss_getpeername ff_getpeername
#define ss_getsockname ff_getsockname
#define ss_getsockopt ff_getsockopt
#define ss_setsockopt ff_setsockopt
#define ss_write ff_write
#define ss_read ff_read
#define ss_send ff_send
#define ss_recv ff_recv
#define ss_epoll_ctl ff_epoll_ctl
#define ss_epoll_wait ff_epoll_wait
#define ss_epoll_create ff_epoll_create

#else
#define ss_fcntl fcntl
#define ss_socket socket
#define ss_bind bind
#define ss_listen listen
#define ss_connect connect
#define ss_accept accept
#define ss_close close
#define ss_shutdown shutdown
#define ss_getpeername getpeername
#define ss_getsockname getsockname
#define ss_getsockopt getsockopt
#define ss_setsockopt setsockopt
#define ss_write write
#define ss_read read
#define ss_send send
#define ss_recv recv
#define ss_epoll_ctl epoll_ctl
#define ss_epoll_wait epoll_wait
#define ss_epoll_create epoll_create
#endif

void dump_socket_req(socket_req_t *req)
{
    char ch[200];
    if (req->op != OpRecv)
    {
        memset(ch, 0, sizeof(ch));
        memcpy(ch, req->data, req->size);
    }
    printf("[socket_req] sizeof: %d, fd: %d, size: %d, op: %d\n", sizeof(socket_req_t), req->fd, req->size, req->op);
}

void dump_socket_rsp(socket_rsp_t *rsp)
{
    char ch[200];
    printf("[socket_rsp] sizeof: %d, fd: %d, size: %d, op: %d, result: %d, data: ", sizeof(socket_rsp_t), rsp->fd, rsp->size, rsp->op, rsp->result);
    memset(ch, 0, sizeof(ch));
    memcpy(ch, rsp->data, rsp->size);
    for (int i = 0; i < rsp->size; i++)
    {
        printf("%d ", rsp->data[i]);
    }
    printf("\n");
}

#define SOCK_ADDR "/tmp/gvisor.sock"

int16_t socket_fd_map[65536];

int socket_handler(socket_req_t *req, socket_rsp_t* rsp) {
    int sockfd, ret;

    struct sockaddr_in serv_addr;
    struct sockaddr_storage peer_addr;
    struct sockaddr_in sa;

    socklen_t peer_addr_size;
    int sa_len;
    int level, name, optlen, backlog;
    void *opt;

    rsp->fd = req->fd;
    rsp->op = req->op;
    switch (req->op)
    {
    case OpCreate:
        sockfd = ss_socket(AF_INET, SOCK_STREAM, 0);
        socket_fd_map[req->fd] = sockfd;
        rsp->result = 0;
        break;

    case OpConnect:
        printf("connect\n");
        ret = ss_connect(socket_fd_map[req->fd], (struct sockaddr *)req->data, req->size);
        printf("connect done %d\n", ret);
        rsp->result = ret;
        break;

    case OpAccept:
        printf("accept\n");
        peer_addr_size = sizeof peer_addr;
        sockfd = ss_accept(socket_fd_map[req->fd], &peer_addr, &peer_addr_size);
        int reserved_fd = *((int *)req->data);
        socket_fd_map[reserved_fd] = sockfd;
        printf("map serverfd %d to gvisor %d\n", sockfd, reserved_fd);
        memcpy(rsp->data, &peer_addr, peer_addr_size);
        rsp->size = peer_addr_size;
        printf("peeraddrlen %d\n", rsp->size);
        printf("\n");
        break;

    case OpBind:
        printf("bind\n");
        // sa = *(struct sockaddr_in *)req->data;
        // printf("Bind sockaddr: family=%u addr=%u, port=%d\n", sa.sin_family, sa.sin_addr, sa.sin_port);
        ret = ss_bind(socket_fd_map[req->fd], (struct sockaddr *)req->data, req->size);
        printf("bind done %d\n", ret);
        rsp->result = ret;
        break;
        
    case OpListen:
        printf("listen\n");
        backlog = *((int32_t *)req->data); // read int backlog from req->data
        printf("listen backlog=%d\n",backlog);
        ret = ss_listen(socket_fd_map[req->fd], backlog);
        printf("listen done %d\n", ret);
        rsp->result = ret;
        break;

    case OpGetSockName:
        rsp->result = ss_getsockname(socket_fd_map[req->fd], &sa, &sa_len);
        memcpy(rsp->data, &sa, sa_len);
        rsp->size = sa_len;
        break;

    case OpGetPeerName:
        rsp->result = ss_getpeername(socket_fd_map[req->fd], &sa, &sa_len);
        memcpy(rsp->data, &sa, sa_len);
        rsp->size = sa_len;
        break;

    case OpSend:
        rsp->result = ss_send(socket_fd_map[req->fd], req->data, req->size, 0);
        break;

    case OpRecv:
        rsp->result = ss_recv(socket_fd_map[req->fd], rsp->data, req->size, 0);
        rsp->size = rsp->result;
        break;

    case OpRelease:
        rsp->result = ss_close(socket_fd_map[req->fd]);
        break;

    case OpSetSockopt:
        level = *((int32_t *)req->data);
        name = *((int32_t *) (req->data + 4));
        optlen = *((int32_t *) (req->data + 8));
        opt = req->data + 12;
        printf("setsockopt level=%d, name=%d, optlen=%d\n", level, name, optlen);
        rsp->result = ss_setsockopt(socket_fd_map[req->fd], level, name, opt, optlen);
        printf("setsockopt done %d\n", rsp->result);
        break;
    }
    return 0;
}

int main()
{
    int server_fd, client_fd, rc;
    socklen_t len;
    struct sockaddr_un server_sockaddr;
    struct sockaddr_un client_sockaddr;
    char buf[256];
    int backlog = 10;
    memset(&server_sockaddr, 0, sizeof(struct sockaddr_un));
    memset(&client_sockaddr, 0, sizeof(struct sockaddr_un));
    memset(buf, 0, 256);

    server_fd = socket(AF_UNIX, SOCK_STREAM, 0);
    if (server_fd == -1)
    {
        printf("SOCKET ERROR: %d\n", errno);
        exit(1);
    }

    server_sockaddr.sun_family = AF_UNIX;
    strcpy(server_sockaddr.sun_path, SOCK_ADDR);
    len = sizeof(server_sockaddr);

    unlink(SOCK_ADDR);
    rc = bind(server_fd, (struct sockaddr *)&server_sockaddr, len);
    if (rc == -1)
    {
        printf("BIND ERROR: %d\n", errno);
        close(server_fd);
        exit(1);
    }

    rc = listen(server_fd, backlog);
    if (rc == -1)
    {
        printf("LISTEN ERROR: %d\n", errno);
        close(server_fd);
        exit(1);
    }
    printf("Listen for incoming socket requests...\n");


    client_fd = accept(server_fd, (struct sockaddr *)&client_sockaddr, &len);
    if (client_fd == -1)
    {
        printf("ACCEPT ERROR: %d\n", errno);
        close(server_fd);
        close(client_fd);
        exit(1);
    }

    len = sizeof(client_sockaddr);
    rc = getpeername(client_fd, (struct sockaddr *)&client_sockaddr, &len);
    if (rc == -1)
    {
        printf("GETPEERNAME ERROR: %d\n", errno);
        close(server_fd);
        close(client_fd);
        exit(1);
    }
    else
    {
        printf("Client socket filepath: %s\n", client_sockaddr.sun_path);
    }


    socket_req_t req;
    socket_rsp_t rsp;
    int ret;
    for (;;)
    {
        if ((ret = recv(client_fd, &req, sizeof(socket_req_t), 0)) != sizeof(socket_req_t))
        {
            printf("error reading uds request: %d\n", ret);
            break;
        }
        dump_socket_req(&req);
        memset(&rsp, 0, sizeof(rsp));
        socket_handler(&req, &rsp);
        if ((ret = send(client_fd, &rsp, sizeof(socket_rsp_t), 0)) != sizeof(socket_rsp_t))
        {
            printf("error writing uds request: %d\n", ret);
            break;
        }
        dump_socket_rsp(&rsp);
    }

    close(server_fd);
    close(client_fd);

    return 0;
}