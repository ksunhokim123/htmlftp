# Introduction
🐁 hosting site that is used in HTML class

## Features
 - Account registration using authorization keys
 - Human-readable + memorable keys
 - The administration page
 - Basic web app(html, js, css)
 - Virtual FTP account

## Installation

## initial

```
go get github.com/sunho/mouse-ftp/server
```

## mouseftp

Go to mouse-ftp and follow the guide to install it.

## api server

Go to [GOPATH](https://github.com/golang/go/wiki/GOPATH)/github.com/sunho/shower-server/server and execute the following command.

```
go build
```

Then you should see "server" binary file. You can configure the server by tweaking "config.yaml" The meanings of entries are explained in the table below.

| field | description | type |
| --- | --- | --- |
| userfile | the path of a file in which the users' data are stored. | path |
| keyfile | the path of a file in which the keys' data are stored | path |
| username | admin username | string |
| password | admin password | string |
| address | address for mousehosting api server | ip / port |
| ftp address |  |   |
| api | address for mouseftp api server this should not be exposed to the public | ip / port |
| ftp | address for mouseftp ftp server | ip / port |

If you omit the ip, it will bind for every ip available.

After you finished the configuration, you can open the api server via.

```
./server
```

### client

Go to [GOPATH](https://github.com/golang/go/wiki/GOPATH)/github.com/sunho/shower-server. And execute these commands.

```
npm install
npx webpack
```

After the process, you should see "dist" folder. You can provide the whole directory through web server such as apache or nginx. Or you can just execute "dist/index.html"
