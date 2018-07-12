# Introduction
🐁 ftp server which provides a user api via http

## Brief explanantion
The server provides a ftp server which you can expose to the public. It also provides a http user api which you **SHOULD NOT** expose to the public as it is a simple REST api that can even delete users and their files easily. 
## Installation
```
git clone https://github.com/sunho/mouse-ftp
cd mouse-ftp
./setup.py install
```
## Excecution
```
mouseftp
```
For more detailed instructions
```
mouseftp -h
```
## API
| endpoint | method | description |
|--|--|--|
| /users | post | add user by form(username, password) |
| /users | get | get users (json) |
| /users/(username) | delete | delete user |

Every response has error entry which contains useful informations. If the error is 'success', it means the task was successfully done. Error codes are not supported yet. It will be added soon. 

