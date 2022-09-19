package config

type Server struct {
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	AUTHKey AUTHKey `mapstructure:"auth-key" json:"auth-key" yaml:"auth-key"`
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Casbin  Casbin  `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	Minio   Minio   `mapstructure:"minio" json:"minio" yaml:"minio"`
	Zinx    Zinx    `mapstructure:"zinx" json:"zinx" yaml:"zinx"`
}
