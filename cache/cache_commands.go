package cache

type Cache struct {
	CacheType        string
	ConnectionString string
}

var Globle_Cache interface{}

func init() {

}

func (cache Cache) Conn() {
	switch cache.CacheType {
	case "redis":
		ConnRedis(cache.ConnectionString)
		Globle_Cache = redis_Cache
	default:
	}
}
