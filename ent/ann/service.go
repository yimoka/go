// Package ann Service 服务配置
package ann

import "entgo.io/ent/entc/gen"

// 当前服务的注解

// Service _

// Service 服务配置
type Service struct {
	// 服务名
	ServiceName string `yaml:"serviceName"`
	// 类别，在微服务中，分给不同的服务进行分类 如 base、app
	ServiceType string `yaml:"serviceType"`
	// proto 目录路径
	ProtoPath string `yaml:"protoPath"`
	// proto 的引用目录
	ProtoImportPath string `yaml:"protoImportPath"`
	// 权限配置的路径
	PermissionsPath string `yaml:"permissionsPath"`
	// 权限配置的父级ID
	PermissionsParentID string `yaml:"permissionsParentID"`
}

const serviceNameKey = "Service"

// Name _
func (Service) Name() string {
	return serviceNameKey
}

// GetService 获取服务配置
func GetService(graph *gen.Graph) *Service {
	ann := graph.Annotations[serviceNameKey]
	if ann == nil {
		return &Service{}
	}
	conf, _ := ann.(Service)
	return &conf
}
