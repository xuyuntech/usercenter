package manager

import (
	"github.com/xuyuntech/usercenter/model"
)

// 操作数据库
// 操作缓存
type Manager interface {
	SaveUser(user *model.User) error
}

type DefaultManager struct {
}

func NewManager() Manager {
	return &DefaultManager{}
}

func (m *DefaultManager) SaveUser(user *model.User) error {
	return nil
}
