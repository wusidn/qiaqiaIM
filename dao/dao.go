package dao

type Dao interface{
	GetTableUsers() TableUsers
	Close()
}