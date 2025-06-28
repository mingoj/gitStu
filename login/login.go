package login

import "fmt"

type UserMod struct {
}

func (*UserMod) Login() {

	/*
		初始化代码
		line 1
		line 2

		line 3

		main分支新增了代码 1 2
		新增了代码 3 4 5   这里的逻辑会影响个人信息修

		初始化好友配置 好友权限
		可能影响 好友相关代码
		代码1
		代码2

	*/
	fmt.Println("user login")

	/*
		dmj 分支新增功能
		新增认证功能 1
		新增认证功能 2
		新增认真功能 3
	*/

}

func NewUserMod() *UserMod {
	return &UserMod{}
}
