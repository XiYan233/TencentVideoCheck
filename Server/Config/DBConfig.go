package Config

func GetDsn() string {
	//MySQL
	//数据库用户名
	const USER = "root"
	//数据库密码
	const PASSWORD = ""
	//数据库地址
	const HOST = "127.0.0.1"
	//数据库端口
	const PORT = "3306"
	//数据库名
	const DBNAME = "tencentvideocheck"
	return USER + ":" + PASSWORD + "@tcp(" + HOST + ":" + PORT + ")/" + DBNAME
}
