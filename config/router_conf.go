package config

// RouterConf ...
type RouterConf struct {
	FacebookCallbackURL string
}

// NewRouterConf ...
func NewRouterConf() (routerConf *RouterConf) {
	routerConf = &RouterConf{}
	Configure(routerConf, "router_conf")
	return
}
