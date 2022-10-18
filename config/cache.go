package config

type Cache struct {
	Username   string   `mapstructure:"username" json:"username" yaml:"username"`          // 用户名
	Password   string   `mapstructure:"password" json:"password" yaml:"password"`          // 密码
	Addr       string   `mapstructure:"addr" json:"addr" yaml:"addr"`                      // 单机地址
	Addrs      []string `mapstructure:"addrs" json:"addrs" yaml:"addrs"`                   // 集群地址
	MasterName string   `mapstructure:"master-name" json:"master-name" yaml:"master-name"` // 哨兵MasterName
	Options    string   `mapstructure:"options" json:"options" yaml:"options"`             // 0(单机)1(集群)
}
