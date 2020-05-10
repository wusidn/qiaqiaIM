module github.com/wusidn/qiaqia

go 1.14

require (
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/protobuf v1.3.5 // indirect
	github.com/jmoiron/sqlx v1.2.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/wusidn/qiaqia/chat v0.0.0
	github.com/wusidn/qiaqia/daomysql v0.0.0
	github.com/wusidn/qiaqia/web v0.0.0
	golang.org/x/sys v0.0.0-20200331124033-c3d80250170d // indirect
)

replace (
	github.com/wusidn/qiaqia/chat => ./chat
	github.com/wusidn/qiaqia/dao => ./dao
	github.com/wusidn/qiaqia/daomysql => ./daomysql
	github.com/wusidn/qiaqia/mockdao => ./mockdao
	github.com/wusidn/qiaqia/web => ./web

)
