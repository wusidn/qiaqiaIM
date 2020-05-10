package daomysql

import(
	"log"
	"errors"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/wusidn/qiaqia/dao"
)

type tableUsers struct{
	db *sqlx.DB
}

type userInfoQuery struct{
	Id 				int 	`db:"UserId"` 
	Locked 			uint8 	`db:"Locked"`
	Email 			sql.NullString	`db:"Email"`
	PhoneNumber 	sql.NullString	`db:"PhoneNumber"`
	EncodePassword 	sql.NullString	`db:"EncodePassWD"`
	IsSystem 		uint8	`db:"IsSystem"`
}

func (this userInfoQuery)toUserInfo() dao.UserInfo{
	return dao.UserInfo{ this.Id, 
		this.Locked,
		If(this.Email.Valid, this.Email.String, "").(string),
		If(this.PhoneNumber.Valid, this.PhoneNumber.String, "").(string),
		If(this.EncodePassword.Valid, this.EncodePassword.String, "").(string),
		this.IsSystem,
	}
}

func (table tableUsers) InsertUserInfo(userInfo dao.UserInfo) error{
	var err error
	for{

		var result sql.Result
		result, err = table.db.Exec("insert into users(CreateDate, ModifyDate, Locked, LockDate) values(now(),now(), false, null)")
		if err != nil{
			log.Printf("insert UserInfo fail 1 error[%v]", err.Error())
			break
		}

		var userId int64
		userId, err = result.LastInsertId()
		if err != nil{
			break
		}

		_, err = table.db.Exec("insert into ADMIN(UserId, Email, PhoneNumber,EncodePassWD, IsSystem) values(?, ?, ?, ?, ?)",
		userId, userInfo.Email, userInfo.PhoneNumber, userInfo.EncodePassword, userInfo.IsSystem)
		if err != nil{
			log.Printf("insert UserInfo fail 2 error[%v]", err.Error())
			break
		}

		break
	}

	return err
}

func (table tableUsers) QueryUserInfo(userId int)(dao.UserInfo, error){
	var result userInfoQuery
	var err error
	for{

		err = table.db.QueryRowx("select UserId, Locked, Email, PhoneNumber,EncodePassWD, IsSystem from users, ADMIN " + 
		"where users.Id = ? and users.Id = ADMIN.UserId", userId).StructScan(&result)

		if err != nil{
			log.Printf("QueryUserInfo fail error[%v]", err.Error())
			break
		}

		break
	}

	return result.toUserInfo(), err
}

func (table tableUsers) QueryUserInfoByEmail(email string)(dao.UserInfo, error){
	var result userInfoQuery
	var err error
	for{

		err = table.db.QueryRowx("select UserId, Locked, Email, PhoneNumber,EncodePassWD, IsSystem from users, ADMIN " + 
		"where ADMIN.Email = ? and users.Id = ADMIN.UserId", email).StructScan(&result)

		if err != nil{
			log.Printf("QueryUserInfoByEmail fail error[%v]", err.Error())
			break
		}

		break
	}

	return result.toUserInfo(), err
}
func (table tableUsers) QueryUserInfoByPhoneNumber(phoneNumber string)(dao.UserInfo, error){
	var result userInfoQuery
	var err error
	for{

		err = table.db.QueryRowx("select UserId, Locked, Email, PhoneNumber,EncodePassWD, IsSystem from users, ADMIN " +
		"where ADMIN.PhoneNumber = ? and users.Id = ADMIN.UserId", phoneNumber).StructScan(&result)
		if err != nil{
			log.Printf("QueryUserInfoByPhoneNumber fail error[%v]", err.Error())
			break
		}
		break
	}
 
	return result.toUserInfo(), err
}
func (table tableUsers) QueryUserInfoByEmailOrPhoneNumber(email string, phoneNumber string)(dao.UserInfo, error){
	var result userInfoQuery
	var err error
	for{

		err = table.db.QueryRowx("select UserId, Locked, Email, PhoneNumber,EncodePassWD, IsSystem from users, ADMIN " +
		"where (ADMIN.Email = ? or ADMIN.PhoneNumber = ?) and users.Id = ADMIN.UserId", email, phoneNumber).StructScan(&result)
		if err != nil{
			log.Printf("QueryUserInfoByPhoneNumber fail error[%v]", err.Error())
			break
		}
		break
	}
 
	return result.toUserInfo(), err
}
func (table tableUsers) UpdateUserInfo(userInfo dao.UserInfo) error{
	var err error
	for{
		var result sql.Result
		result, err = table.db.Exec("update ADMIN set Email=?, PhoneNumber=?, EncodePassWD=? where UserId=?", userInfo.Email, userInfo.PhoneNumber, userInfo.EncodePassword, userInfo.Id)

		if err != nil{
			break
		}

		var rowsAffected int64
		rowsAffected, err = result.RowsAffected()

		if err != nil{
			break
		}

		if rowsAffected != 1 {
			err = errors.New("update fial error[not find UserId]")
			break
		}

		break
	}

	return err
}