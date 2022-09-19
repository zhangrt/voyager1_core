package config

type Zinx struct {
	Name           string `mapstructure:"name" json:"name" yaml:"name"`                                     // 服务器应用名称
	Host           string `mapstructure:"host" json:"host" yaml:"host"`                                     // 服务器IP
	TcpPort        int    `mapstructure:"tcp-port" json:"tcp-port" yaml:"tcp-port"`                         // 服务器监听端口
	MaxConn        int    `mapstructure:"max-conn" json:"max-conn" yaml:"max-conn"`                         // 允许的客户端链接最大数量
	WorkerPoolSize int    `mapstructure:"worker-pool-size" json:"worker-pool-size" yaml:"worker-pool-size"` // 工作任务池最大工作Goroutine数量
	LogDir         string `mapstructure:"log-dir" json:"log-dir" yaml:"log-dir"`                            // 日志文件夹
	LogFile        string `mapstructure:"log-file" json:"log-file" yaml:"log-file"`                         // 日志文件名称(如果不提供，则日志信息打印到Stderr)
}
