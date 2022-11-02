package config

type Grpc struct {
	Server GrpcServer `mapstructure:"j-server" json:"j-server" yaml:"j-server"`
	Client GrpcClient `mapstructure:"j-client" json:"j-client" yaml:"j-client"`
}

type GrpcServer struct {
	Network string `mapstructure:"network" json:"network" yaml:"network"` // 网络, tcp、udp等
	Host    string `mapstructure:"host" json:"host" yaml:"host"`          // 服务器IP
	Port    int    `mapstructure:"port" json:"port" yaml:"port"`          // 服务器port
}

type GrpcClient struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"` // 服务器IP
	Port int    `mapstructure:"port" json:"port" yaml:"port"` // 服务器port
}
