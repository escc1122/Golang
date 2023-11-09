package simple

import (
	"github.com/casbin/casbin/v2/model"
)

func geACLModel() (model.Model, error) {
	m, _ := model.NewModelFromString(`
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
`)

	return m, nil
}
