import socket

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect(("199.232.64.223", 80))
print(s)
s.close()