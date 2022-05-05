package lowcode

import (
	lu "github.com/quanxiang-cloud/faas-lowcode-interface/lowcode/user"
)

const (
	LOWCODE = "lowcode-interface"
)

type Lowcode interface {
	lu.User
}
