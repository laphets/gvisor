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

void dump_socket_req(socket_req_t* req) {
    char ch[200];
    memset(ch, 0, sizeof(ch));
    memcpy(ch, req->data, req->size);
    printf("[socket_req] sizeof: %d, fd: %d, size: %d, op: %d, data: %s\n", sizeof(socket_req_t), req->fd, req->size, req->op, ch);
}

void dump_socket_rsp(socket_rsp_t *rsp)
{
    char ch[200];
    memset(ch, 0, sizeof(ch));
    memcpy(ch, rsp->data, rsp->size);
    printf("[socket_rsp] sizeof: %d, fd: %d, size: %d, op: %d, result: %d, data: %s\n", sizeof(socket_rsp_t), rsp->fd, rsp->size, rsp->op, rsp->result, ch);
}

#define SOCK_ADDR "/tmp/gvisor.sock"

int16_t socket_fd_map[65536];

int socket_handler(socket_req_t *req, socket_rsp_t* rsp) {
    int sockfd, ret;
    struct sockaddr_in serv_addr;
    rsp->fd = req->fd;
    rsp->op = req->op;
    switch (req->op)
    {
    case OpCreate:
        sockfd = socket(AF_INET, SOCK_STREAM, 0);
        socket_fd_map[req->fd] = sockfd;
        rsp->result = 0;
        break;

    case OpConnect:
        printf("connect\n");
        ret = connect(socket_fd_map[req->fd], (struct sockaddr *)req->data, req->size);
        printf("connect done %d\n", ret);
        rsp->result = ret;
        break;

    case OpSend:
        break;

    case OpRelease:
        rsp->result = close(socket_fd_map[req->fd]);
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