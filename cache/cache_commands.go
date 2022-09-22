package cache

type Cache struct {
}

var Globle_Cache interface{}

func init() {

}

func Conn(cachetype string, constr string) {
	switch cachetype {
	case "redis":
		ConnRedis(constr)
		Globle_Cache = Redis_Cache
	default:
	}
}
