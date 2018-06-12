package manager

import (
	"github.com/go-xorm/xorm"
	"github.com/urfave/cli"
	"github.com/xuyuntech/usercenter/model"
)

// 操作数据库
// 操作缓存
type Manager interface {
	SaveUser(user *model.User) error
}

type DefaultManager struct {
	engine *xorm.Engine
}

func NewManager(c *cli.Context) (Manager, error) {
	engine, err := model.NewEngine(c.String("database-datasource"), []interface{}{new(model.User)})
	if err != nil {
		return nil, err
	}
	return &DefaultManager{
		engine: engine,
	}, nil
}

func (m *DefaultManager) SaveUser(user *model.User) error {
	return nil
}
