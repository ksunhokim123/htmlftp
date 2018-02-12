from pyftpdlib.authorizers import DummyAuthorizer
from pyftpdlib.handlers import FTPHandler
from pyftpdlib.servers import FTPServer
authorizer = DummyAuthorizer()
authorizer.add_user("user", "12345", "/home/test", perm="elradfmwMT")
handler = FTPHandler
handler.authorizer = authorizer
server = FTPServer(("127.0.0.1", 5693), handler)
server.serve_forever()
