# 权限管理AUTH包

`众星 star 拱月 luna 的设计思想`

luna 作为基础平台将对所有注册star的客户端服务进行授权鉴权,通信方式有GRPC、Zinx、Rest等多种方式可选

## 平台基础组件luna包

平台服务组件handler中有标准的授权鉴权拦截器

## 业务组件star包

## grpc（v0.1推荐）、Zinx

* grpc和zinx都可使用tcp或udp等不同的通信方式，目前建议使用由GRPC实现的稳定版本

* zinx版本采用长链接形式在star（client）端设计上需要重构

* zinx和Grpc对比在于Zinx更轻量、自定义性更强且实现好了会更高效，而Grpc封装的更彻底、开发和使用也更简单

* 使用GRPC在每次更改协议，只需要重新生成pb.go文件，业务代码的改动相比较会更少更简单，不必专门维护协议底层实现的代码

### proto 协议

* 安装 protobuf编译器

* 安装go protobuf插件

* 编写proto文件 定义服务、rpc方法和message

* <https://grpc.io/docs/languages/go/basics>

* 使用GRPC生成proto相关文件

  `cd auth/grpc/pb`

  `protoc --go_out=. --go_opt=paths=source_relative   --go-grpc_out=. --go-grpc_opt=paths=source_relative  auth.proto`

* 实现定义的GRPC接口 pb.UnimplementedAuthServiceServer
