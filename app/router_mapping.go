package app

var (
	HostRouterMapping   map[string]string        //key : domain  value : service id
	PrefixRouterMapping map[string]*PrefixRouter // key : url prefix value: prefixRouter
)
