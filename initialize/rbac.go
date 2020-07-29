package initialize

import (
	"gingo/internal/config"
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"go.uber.org/zap"
)

// Enforcer Enforcer
var (
	Enforcer *casbin.Enforcer
)

// InitRbac init rabc
func InitRbac() {
	modeltext := `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act)
`
	conf := config.Conf
	m, err := model.NewModelFromString(modeltext)
	if err != nil {
		log.Panic("NewModelFromString Err", zap.Error(err))
	}
	adapter, err := gormadapter.NewAdapter("mysql", conf.MysqlDsn, true)
	if err != nil {
		log.Panic("NewAdapter Err", zap.Error(err))
	}

	enforcer, err := casbin.NewEnforcer(m, adapter)
	if err != nil {
		log.Fatal("InitRabc failed", zap.Error(err))
	}
	Enforcer = enforcer
}

// // GetEnforcer get casbin auth
// func GetEnforcer() *casbin.Enforcer {
// 	return enforcer
// }
