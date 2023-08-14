#  Go语言的安装与开发环境

下载：

国内：[http://studygolang.com/dl](https://gitee.com/link?target=http%3A%2F%2Fstudygolang.com%2Fdl)

[https://golang.org/dl/](https://gitee.com/link?target=https%3A%2F%2Fgolang.org%2Fdl%2F)

```
# 设置国内镜像
go env -w GOPROXY=https://goproxy.cn,direct
# 开启 Go Module
go env -w GO111MODULE=on
# goimports
go get -v golang.org/x/tools/cmd/goimports
```

开发环境：vi, emacs, idea, eclipse, vs, sublime … + go 插件

IDE：Goland, liteIDE