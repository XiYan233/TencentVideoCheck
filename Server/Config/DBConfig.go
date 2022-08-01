package Config

func GetDsn() string {
	var config Conf
	config.getConf()
	//MySQL
	//数据库用户名
	dbUser := config.DBUser
	//数据库密码
	dbPassword := config.DBPassword
	//数据库地址
	dbHost := config.DBHost
	//数据库端口
	dbPort := config.DBPort
	//数据库名
	dbName := config.DBName
	return dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
}
