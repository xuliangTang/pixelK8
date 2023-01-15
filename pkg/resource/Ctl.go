package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	"k8s.io/client-go/kubernetes"
	"strings"
)

// ResourceCtl @Controller
type ResourceCtl struct {
	K8sClient *kubernetes.Clientset `inject:"-"`
}

func NewResourceCtl() *ResourceCtl {
	return &ResourceCtl{}
}

func (this *ResourceCtl) apiResources(ctx *gin.Context) any {
	_, resourceList, err := this.K8sClient.ServerGroupsAndResources()
	athena.Error(err)

	groupResourceList := make([]GroupResource, 0)

	for _, resource := range resourceList {
		group, version := this.getGroupAndVersion(resource.GroupVersion)
		gr := GroupResource{
			Group:   group,
			Version: version,
		}

		for _, ar := range resource.APIResources {
			rs := Resource{
				Name:  ar.Name,
				Verbs: ar.Verbs,
			}
			gr.Resources = append(gr.Resources, rs)
		}

		groupResourceList = append(groupResourceList, gr)
	}

	return groupResourceList
}

func (this *ResourceCtl) getGroupAndVersion(groupVersion string) (group, version string) {
	r := strings.Split(groupVersion, "/")
	if len(r) == 1 {
		return "core", r[0]
	} else if len(r) == 2 {
		return r[0], r[1]
	}

	panic("error groupVersion: " + groupVersion)
}

func (this *ResourceCtl) Build(athena *athena.Athena) {
	// api-resource列表
	athena.Handle("GET", "/resources", this.apiResources)
}
