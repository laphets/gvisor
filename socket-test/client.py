import socket

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
print(s)
s.connect(("0.0.0.0", 8080))


print(s)
for i in range(10):
    data = b"test%d"%(i)
    s.send(data)
    print("recv:", s.recv(len(data)))
s.close()
