package config

type Grpc struct {
	Server GrpcServer `mapstructure:"server" json:"server" yaml:"server"`
	Client GrpcClient `mapstructure:"client" json:"client" yaml:"client"`
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
