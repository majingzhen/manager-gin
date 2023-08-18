// Package view 自动生成模板 SysUser
// @description <TODO description class purpose>
// @author
// @File: sys_user
// @version 1.0.0
// @create 2023-08-18 13:41:26
package view

// LoginUserView 结构体
type LoginUserView struct {
	UserName   string `json:"userName"`
	Password   string `json:"password"`
	VerifyCode string `json:"verifyCode"`
	VerifyKey  string `json:"verifyKey"`
}

// SysUserInfoView 用户信息 view
type SysUserInfoView struct {
	Id          string `json:"id"`
	OrgId       string `json:"orgId"`
	UserName    string `json:"userName"`
	NickName    string `json:"nickName"`
	UserType    string `json:"userType"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Sex         string `json:"sex"`
	Avatar      string `json:"avatar"`
}

type LoginUserResView struct {
	Token       string          `json:"token"`
	SysUserInfo SysUserInfoView `json:"userInfo"`
}
