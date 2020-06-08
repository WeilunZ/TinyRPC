package selector

type Selector interface {
	Select(string) (string, error)
}

type defaultSelector struct {
}

type Options struct {
}

type Option func(*Options)

func init() {
	RegisterSelector("default", DefaultSelector)
}

var DefaultSelector = &defaultSelector{}

var selectorMap = make(map[string]Selector)

func RegisterSelector(name string, selector Selector) {
	if selectorMap == nil {
		selectorMap = make(map[string]Selector)
	}
	selectorMap[name] = selector
}

func (d *defaultSelector) Select(serviceName string) (string, error) {
	// 会基于服务发现的第三方库如zookeeper,consul,etcd等实现，此处忽略
	return "", nil
}

func GetSelector(name string) Selector {
	if selector, ok := selectorMap[name]; ok {
		return selector
	}
	return DefaultSelector
}
