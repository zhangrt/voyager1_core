# Luna-Server & star-Client

* 2022-09-26，发布第一版Luna Server

* 2022-09-27，目前暂停更新维护此模块,后续更新待定

## Luna-Server

`基于zinx-server搭建的tcp server`

## star-Client

`star`

## 测试类

`zinx_test.go`

## Proto 编译

* cd zinx/pb

* protoc --go_out=. zinx.proto

## 配置文件

`/conf/zinx.json 在外部引入组件时，优先会读取外部工程中/conf目录下zinx.json文件`
