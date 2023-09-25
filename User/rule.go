package User

import (
	_ "embed"
	"io/fs"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

var (
	//go:embed rule.casbin.txt
	CasbinModel  string
	AuthEnforcer *casbin.Enforcer

	//go:embed rule.casbin.txt
	Rules fs.FS
)

// func init() {
// 	// 加载 casbin 鉴权规则
// 	if enforcer, err := casbin.NewEnforcerSafe(
// 		casbin.NewModel(CasbinModel),
// 		gormadapter.NewAdapter(
// 			"mysql",
// 			"root:123456@tcp(127.0.0.1:3306)/",
// 		),
// 	); err != nil {
// 		panic(err)
// 	} else {
// 		AuthEnforcer = enforcer
// 	}

// 	if err := AuthEnforcer.LoadPolicy(); err != nil {
// 		panic(err)
// 	}

// 	insertPolicy()
// }

func init() {
	// 加载 casbin 鉴权规则
	a, err := gormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}

	e, err := casbin.NewEnforcer("./User/rule.casbin.txt", a)
	if err != nil {
		panic(err)
	}

	AuthEnforcer = e

	if err = AuthEnforcer.LoadPolicy(); err != nil {
		panic(err)
	}

	insertPolicy()
}

func insertPolicy() {
	AuthEnforcer.AddPolicy("admin", "/*", "*")
	AuthEnforcer.AddPolicy("anonymous", "/login", "*")
	AuthEnforcer.AddPolicy("member", "/logout", "*")
	AuthEnforcer.AddPolicy("member", "/member/*", "*")

	if err := AuthEnforcer.SavePolicy(); err != nil {
		panic(err)
	}
}
