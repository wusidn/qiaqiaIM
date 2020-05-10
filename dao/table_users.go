package dao

type UserInfo struct{
	Id 				int 	`db:"UserId"` 
	Locked 			uint8 	`db:"Locked"`
	Email 			string	`db:"Email"`
	PhoneNumber 	string	`db:"PhoneNumber"`
	EncodePassword 	string	`db:"EncodePassWD"`
	IsSystem 		uint8	`db:"IsSystem"`
}

type TableUsers interface{
	InsertUserInfo(userInfo UserInfo) error
	QueryUserInfo(userId int)(UserInfo, error)
	QueryUserInfoByEmail(email string)(UserInfo, error)
	QueryUserInfoByPhoneNumber(phoneNumber string)(UserInfo, error)
	QueryUserInfoByEmailOrPhoneNumber(email string, phoneNumber string)(UserInfo, error)
	UpdateUserInfo(userInfo UserInfo) error
}



