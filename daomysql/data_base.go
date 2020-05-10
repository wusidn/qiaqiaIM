package daomysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/wusidn/qiaqia/dao"
)

type daoMysql struct{
	dataBase *sqlx.DB
}

type DataBaseConnectionInfo struct{
	UserName 	string
	Password 	string
	IpAddress 	string
	Port		int
	DBName		string
	Charset		string
}

func Default()(dao.Dao, error){
	connectionInfo := &DataBaseConnectionInfo{userName, 
		password, ipAddress, port, dbName, charset,}

	return initDataBase(connectionInfo)
}

func New(connectionInfo *DataBaseConnectionInfo)(dao.Dao, error){
	return initDataBase(connectionInfo)
}

func initDataBase(connectionInfo *DataBaseConnectionInfo)(dao.Dao, error){
	var dao daoMysql
	var err error
	sqlConnectStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&multiStatements=true", 
		connectionInfo.UserName, connectionInfo.Password, connectionInfo.IpAddress,
		connectionInfo.Port, connectionInfo.DBName, connectionInfo.Charset)

	for{
		var db *sqlx.DB
		db, err = sqlx.Open("mysql", sqlConnectStr)
		if err != nil{
			break
		}

		err = db.Ping()
		if err != nil{
			break
		} 
		dao.dataBase = db
		break
	}
	
	return dao, err
}

func (dao daoMysql) GetTableUsers() dao.TableUsers{
	return tableUsers{dao.dataBase}
}

func (dao daoMysql) Close(){
	dao.dataBase.Close()
	dao.dataBase = nil
}