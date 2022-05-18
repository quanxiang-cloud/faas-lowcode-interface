package lowcode

import (
	lf "github.com/quanxiang-cloud/faas-lowcode-interface/lowcode/form"
	lu "github.com/quanxiang-cloud/faas-lowcode-interface/lowcode/user"
)

const (
	LOWCODE = "lowcode-interface"
)

type Lowcode interface {
	lu.User
	lf.Form
}
