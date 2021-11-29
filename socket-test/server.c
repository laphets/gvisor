#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <errno.h>
#include <string.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <netdb.h>
#include <arpa/inet.h>
#include <sys/wait.h>
#include <signal.h>

#define MYPORT "8080"  // the port users will be connecting to
#define BACKLOG 6     // how many pending connections queue will hold

// get sockaddr, IPv4 or IPv6:
void *get_in_addr(struct sockaddr *sa)
{
    if (sa->sa_family == AF_INET)
    {
        return &(((struct sockaddr_in *)sa)->sin_addr);
    }

    return &(((struct sockaddr_in6 *)sa)->sin6_addr);
}

int main(void)
{
    struct sockaddr_storage their_addr;
    socklen_t addr_size;
    struct addrinfo hints, *res;
    int sockfd, new_fd;
    char s[INET6_ADDRSTRLEN];

    // !! don't forget your error checking for these calls !!
    // first, load up address structs with getaddrinfo():

    memset(&hints, 0, sizeof hints);
    hints.ai_family = AF_INET;  // use IPv4 or IPv6, whichever
    hints.ai_socktype = SOCK_STREAM;
    hints.ai_flags = AI_PASSIVE;     // fill in my IP for me

    getaddrinfo(NULL, MYPORT, &hints, &res);

    // make a socket, bind it, and listen on it:

    sockfd = socket(res->ai_family, res->ai_socktype, res->ai_protocol);
    if (bind(sockfd, res->ai_addr, res->ai_addrlen) == -1)
    {
        close(sockfd);
        perror("server: bind");
    }

    if (listen(sockfd, BACKLOG) == -1)
    {
        perror("listen");
        exit(1);
    }

    // now accept an incoming connection:
    char recvbuf[1024];

    // accept() recv() at the same process
    addr_size = sizeof their_addr;
    new_fd = accept(sockfd, (struct sockaddr *)&their_addr, &addr_size);
    if (new_fd == -1)
    {
        printf("socket accept return new_fd = -1\n");
        perror("accept");
        exit(1);
    }

    printf("sockfd=%d, new_fd after accept = %d\n", sockfd, new_fd);
    // printf("Accept peer_addr: family=%u port=%u,\n", their_addr.sin_family, their_addr.sin_port);
    printf("peeraddrlen %d\n", addr_size);

    // inet_ntop(their_addr.ss_family,
    //           get_in_addr((struct sockaddr *)&their_addr),
    //           s, sizeof s);
    printf("server: got connection from \n");

    while(1){
        memset(recvbuf, 0, 1024);
        int n = recv(new_fd, recvbuf, 1024, 0);
        if (n == -1){
            perror("recv");
            exit(1);
        }
        // printf("recv_from_client %s\n", recvbuf);

        if (send(new_fd, recvbuf, n, 0) == -1)
            perror("send");
    }
    printf("send done, shut down server\n");
    close(new_fd);
    exit(0);

    
    close(new_fd); // parent doesn't need this
    return 0;


}
