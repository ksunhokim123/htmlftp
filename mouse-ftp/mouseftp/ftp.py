import os
import shutil

from hashlib import md5
from pyftpdlib.authorizers import DummyAuthorizer
from pyftpdlib.handlers import FTPHandler
from pyftpdlib.servers import FTPServer

class DummyMD5Authorizer(DummyAuthorizer):
    def validate_authentication(self, username, password, handler):
        password = password.encode('latin1')
        hash = md5(password).hexdigest()
        try:
            if self.user_table[username]['pwd'] != hash:
                raise KeyError
        except KeyError:
            raise AuthenticationFailed

class ftp_server:
   def __init__(self, address, home):
       self.authorizer = DummyMD5Authorizer()
       self.address = address
       self.home = home

   def run(self):
       self.handler = FTPHandler
       self.handler.authorizer = self.authorizer
       self.handler.banner = 'Welcome to Sunho Mouse FTP server. %s'
       self.server = FTPServer(self.address, self.handler)
       self.server.serve_forever()

   def add_user(self,user,password):
       password = password.encode('latin1')
       hash = md5(password).hexdigest()
       home_path = str(self.home + '/' + user)
       if not os.path.exists(home_path):
           os.mkdir(home_path)
       self.authorizer.add_user(str(user), hash, home_path, perm='elradfmwMT')

   def remove_user(self,user):
       home_path = str(self.home + '/' + user)
       self.authorizer.remove_user(str(user))
       shutil.rmtree(home_path)

   def get_users(self):
       keys = []
       for key in self.authorizer.user_table.keys():
           keys.append(key)
       return keys
