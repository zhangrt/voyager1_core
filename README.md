# gallery

模块封装

## 使用组件时添加依赖

* go get github.com/zhangrt/voyager1_core

* Init初始化 import "github.com/zhangrt/voyager1_core"  -> NewInit()

## 代码结构

```shell
├── auth
│   ├── luna
│   └── star
├── cache
├── config
├── constant
├── global
│   ├── request
│   └── response
├── log
├── oss
├── util
└── zinx

```

| 文件夹       | 说明                    | 描述                        |
| ------------ | ----------------------- | --------------------------- |
| `auth`        | auth 组件                | auth 接口                 |
| `--luna`      | auth 平台组件            | auth  平台组件接口         |
| `--star`      | auth 业务组件            | auth 业务组件接口          |  
| `cache`       | cache组件                | cache接口                 |
| `config`      | 配置文件                 | 组件配置                   |
| `constant`    | constant常量             | constant常量              |
| `global`      | 全局对象                 | 全局对象                   |
| `--request`   | 入参结构体               | 接收前端发送到后端的数据。   |
| `--response`  | 出参结构体               | 返回给前端的数据结构体       |
| `log`         | 日志组件                 | 日志组件接口                |
| `oss`         | oss组件                  | oss组件接口                |
| `util`        | 工具包                   | 工具函数封装                |
| `zinx`        | zinx核心服务             | 基于zinx实现的高并发Server   |

## ProtoBuf

### 安装protoc

<https://github.com/protocolbuffers/protobuf/releases/>

* linux 安装  
测试protoc -h  
* windwos 安装  
测试 protoc --version

### 安装protobuf-go

<https://github.com/protocolbuffers/protobuf-go>

#### 获取 proto包(Go语言的proto API接口)

* go get  -v -u github.com/golang/protobuf/proto
* go get  -v -u github.com/golang/protobuf/protoc-gen-go

#### 编译protoc-gen-go

* cd $GOPATH/src/github.com/golang/protobuf/protoc-gen-go/
* go build

#### 将生成的 protoc-gen-go可执行文件，放在/bin目录下

* sudo cp protoc-gen-go /bin/

#### 或直接下载安装本地

<https://github.com/protocolbuffers/protobuf-go/releases>

### 创建.proto文件

#### 定义message格式

### 编译.proto文件

#### protoc 编译生成.go文件 编译后的.pb.go文件无法修改

* protoc --proto_path=IMPORT_PATH --go_out=DST_DIR path/to/file.proto

* protoc --go_out=. *.proto 将本地当前文件下所有.proto文件全部编译并将结果放在当前文件夹下

--proto_path，指定了 .proto 文件导包时的路径，可以有多个，如果忽略则默认当前目录。

--go_out， 指定了生成的go语言代码文件放入的文件夹

允许使用protoc --go_out=./ *.proto的方式一次性编译多个 .proto 文件

编译时，protobuf 编译器会把 .proto 文件编译成 .pd.go 文件

### 使用.proto开发

## 开发约束

### 文件格式

* 小写字母，当有多个单词时通过 _ 连接

### git commit message 约束

**_type: subject(scope)_**

* type：用于说明commit的类别，规定为如下几种
* feat：新增功能；
* fix：修复bug；
* docs：修改文档；
* refactor：代码重构，未新增任何功能和修复任何bug；
* build：改变构建流程，新增依赖库、工具等（例如webpack修改）；
* style：仅仅修改了空格、缩进等，不改变代码逻辑；
* perf：改善性能和体现的修改；
* chore：非src和test的修改；
* test：测试用例的修改；
* ci：自动化流程配置修改；
* revert：回滚到上一个版本；
* scope：【可选】用于说明commit的影响范围
* subject：commit的简要说明，尽量简短

## 代码质量检查

### pre-commit

pip install pre-commit，需要安装python环境

#### Doc

<https://pre-commit.com/>

### 使用git hook

* 1、执行 pre-commit install 将会在git hook目录安装pre-commit文件 无需修改

* 2、复制scripts目录下commit-msg 和 pre-push 两个文件到工程 ./.git/hook 目录下

* 3、建议使用sourcetree提交代码，可以看到hook的完整日志
