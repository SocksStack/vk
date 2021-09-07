# vk
快速构建 Go 项目脚手架

# 生成项目结构

```
demo
 ├── bootstrap
 │   └── bootstrap.go
 ├── config
 │   ├── default.yml
 │   ├── development.yml
 │   ├── production.yml
 │   └── test.yml
 ├── handler
 │   └── index.go
 ├── middleware
 │   └── init.go
 ├── router
 │   └── router.go
 ├── main.go
 ├── go.mod
 ├── go.sum
 └── README.md
```

# Usage
```shell
# 创建项目
vk new demo

# 创建 service
vk make:service test
# 或者
vk ms test

# 创建 handler
vk make:handler test
# 或者
vk mh test

# 创建 middleware
vk make:middleware test
# 或者
vk mm test
```
