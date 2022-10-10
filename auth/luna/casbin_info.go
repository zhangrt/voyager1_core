package luna

// Casbin info structure
type CasbinInfo struct {
	Path   string `json:"path"`   // 路径
	Method string `json:"method"` // 方法
}

// Casbin structure for input parameters
type CasbinInReceive struct {
	AuthorityId string       `json:"authorityId"` // 权限id
	CasbinInfos []CasbinInfo `json:"casbinInfos"`
}

// 默认的Casbin policy, 实际规则可由使用者根据需要存入数据库表中
func DefaultCasbin() []CasbinInfo {
	return []CasbinInfo{
		{Path: "/menu/getMenu", Method: "POST"},
		{Path: "/jwt/jsonInBlacklist", Method: "POST"},
		{Path: "/base/login", Method: "POST"},
		{Path: "/user/register", Method: "POST"},
		{Path: "/user/changePassword", Method: "POST"},
		{Path: "/user/setUserRole", Method: "POST"},
		{Path: "/user/setUserInfo", Method: "PUT"},
		{Path: "/user/getUserInfo", Method: "GET"},
	}
}
