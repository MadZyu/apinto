package upstream_http_anonymous

import (
	"reflect"

	"github.com/eolinker/eosc"
)

const (
	driverName = "http_proxy_anonymous"
)

//driver 实现github.com/eolinker/eosc.eosc.IProfessionDriver接口
type driver struct {
	profession string
	name       string
	driver     string
	label      string
	desc       string
	configType reflect.Type
	params     map[string]string
}

//ConfigType 返回匿名http_proxy驱动配置的反射类型
func (d *driver) ConfigType() reflect.Type {
	return d.configType
}

//Create 创建匿名http_proxy驱动的实例
func (d *driver) Create(id, name string, v interface{}, workers map[eosc.RequireId]interface{}) (eosc.IWorker, error) {
	w := &httpUpstream{
		id:     id,
		name:   name,
		driver: driverName,
	}
	return w, nil
}
