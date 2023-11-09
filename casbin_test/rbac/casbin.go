package simple

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
)

type permission struct {
	id     string
	path   string
	method string
}

type feature struct {
	id           string
	permissionId string
}

type role struct {
	id        string
	featureId string
}

func geRBACModel(role []role, feature []feature, permission []permission) (*casbin.Enforcer, error) {
	m, _ := model.NewModelFromString(`
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act || r.sub == "admin"
`)

	e, _ := casbin.NewEnforcer(m)

	for _, r := range role {
		e.AddGroupingPolicy(r.id, r.featureId)
	}

	for _, f := range feature {
		e.AddGroupingPolicy(f.id, f.permissionId)
	}

	for _, p := range permission {
		e.AddPolicy(p.id, p.path, p.method)
	}

	return e, nil
}
