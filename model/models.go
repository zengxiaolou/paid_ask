/**
AUTHOR:         zeng_xiao_yu
GITHUB:         https://github.com/zengxiaolou
EMAIL:          zengevent@gmail.com
TIME:           2020/11/6-14:31
INSTRUCTIONS:   数据库相关
**/

package model

import "database/sql"

var Models = []interface {}{
	&User{},
}

type Model struct {
	Id		int64		`gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
}

type User struct {
	Model
	Username			sql.NullString 	`gorm:"size:32;unique;" json:"username" form:"username"`					// 用户名
	Email 				sql.NullString	`gorm:"size:128;unique;" json:"email" form:"email"`							// 邮箱
	EmailVerified		bool			`gorm:"not null;default:false" json:"emailVerified" form:"emailVerified"`	// 邮箱是否验证
	Mobile 				string			`gorm:"size:11;unique;" json:"mobile" form:"mobile"`						// 手机号
	Nickname			string			`gorm:"size:16" json:"nickname" form:"nickname"`							// 昵称
	Avatar				string			`gorm:"type:text" json:"avatar" form:"avatar"`								// 头像
	Password 			string 			`gorm:"size:256" json:"password" form:"password"`							// 密码
	Roles 				string			`gorm:"type:text" json:"roles" form:"roles"`								// 角色
	Type				int				`gorm:"not null" json:"type" form:"type"`									// 用户类型
	CreateTime 			int64 			`json:"createTime" form:"createTime"`										// 创建时间
	LastLoginTime		int64			`json:"lastLoginTime" form:"lastLoginTime"`									// 最后登录时间
}