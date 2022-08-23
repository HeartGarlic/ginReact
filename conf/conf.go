package conf

type Conf struct {
}

// InitConf 初始化配置
func InitConf() {
	// 初始化xOrm 引擎
	InitMysql()
}
