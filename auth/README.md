# 权限管理AUTH包

`众星 star 拱月 luna`

luna 作为基础平台将对所有注册star的客户端服务进行授权鉴权,通信方式有TCP、GRPC、Rest等多种方式可选

## 平台基础组件luna包

## 业务组件star包

## grpc

### proto

* 安装 protobuf编译器

* 安装go protobuf插件

* 编写proto文件 定义服务、rpc方法和message

<https://grpc.io/docs/languages/go/basics>

* 生成proto文件

`protoc --go_out=. --go_opt=paths=source_relative   --go-grpc_out=. --go-grpc_opt=paths=source_relative  auth.proto`
