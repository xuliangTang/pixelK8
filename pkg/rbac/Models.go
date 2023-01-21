package rbac

import rbacV1 "k8s.io/api/rbac/v1"

type RoleListModel struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	CreatedAt string `json:"created_at"`
}

type RoleDetailModel struct {
	Name      string              `json:"name"`
	Namespace string              `json:"namespace"`
	Rules     []rbacV1.PolicyRule `json:"rules"`
	CreatedAt string              `json:"created_at"`
}

type RoleBindingListModel struct {
	Name      string           `json:"name"`
	Namespace string           `json:"namespace"`
	RoleRef   string           `json:"role_ref"`
	Subjects  []rbacV1.Subject `json:"subjects"`
	CreatedAt string           `json:"created_at"`
}

type ServiceAccountListModel struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	CreatedAt string `json:"created_at"`
}
