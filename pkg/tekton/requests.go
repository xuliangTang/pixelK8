package tekton

type RsUri struct {
	Namespace string `uri:"ns" binding:"required"`
	Name      string `uri:"name" binding:"required"`
}
