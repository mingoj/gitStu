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


	*/
	fmt.Println("user login")

}

func NewUserMod() *UserMod {
	return &UserMod{}
}
