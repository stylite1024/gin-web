# gin-web

[English](./README.md)

## 简介

gin-web 是一个基于 go gin 和 vue3 框架的 web 模板。它采用的是前后端不分离模式，可用于快速开发 web 应用。

## 目录结构

详看：[project-layout](https://github.com/golang-standards/project-layout)

```
.
├── air.toml
├── cmd
│   ├── root.go
│   └── version.go
├── config
│   └── info.go
├── docker-compose.yml
├── Dockerfile
├── Dockerfile.alpine
├── docs
├── go.mod
├── go.sum
├── internal
│   ├── handler
│   │   └── handler.go
│   ├── logger
│   ├── logic
│   └── router
│       └── router.go
├── LICENSE
├── main.go
├── Makefile
├── pkg
│   └── util
│       ├── ip.go
│       └── response.go
├── README.md
├── README-zh.md
├── scripts
│   ├── docker-entrypoint.sh
│   └── nginx.conf
├── test
│   └── gin_test.go
└── web
    ├── static
    │   ├── api
    │   │   └── test.js
    │   ├── assets
    │   │   ├── img
    │   │   │   └── favicon.ico
    │   │   └── js
    │   │       └── request.js
    │   ├── index2.html
    │   ├── index.html
    │   └── pages
    │       └── demo
    │           └── demo.html
    └── static.go
```

## 快速开始

### 要求

golang 版本 1.22.0 及以上

### 开发

```sh
git clone https://github.com/stylite1024/gin-web.git
cd gin-web
go mod tidy
go run main.go
```

浏览器访问：http://0.0.0.0:8080

### 部署

```sh
# build binary file
make build

# build docker image
make build-image
```

## 致谢

-   [go](https://github.com/golang/go)
-   [gin](https://github.com/gin-gonic/gin)
-   [gin-contrib](https://github.com/gin-contrib)
-   [vuejs](https://github.com/vuejs/vue)
-   [axios](https://github.com/axios/axios)
-   [element-plus-ui](https://github.com/element-plus/element-plus)
-   [cobra](https://github.com/spf13/cobra)
-   [cobra-cli](https://github.com/spf13/cobra-cli)
-   [viper](https://github.com/spf13/viper)

## 提问

Found an error? Is there something meaningless? Initiate an [issue](https://github.com/stylite1024/gin-web/issues) to me, thank you!

## 开源协议

[MIT](https://github.com/stylite1024/gin-web/blob/main/LICENSE)
