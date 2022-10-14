package config

type Cache struct {
	Addr       string   //单机地址
	Password   string   //密码
	Addrs      []string //集群地址
	MasterName string   //哨兵MasterName
	Options    string   //G_REDIS_STANDALONE(单机)、G_REDIS_CLUSTER(集群)
}
