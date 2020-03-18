package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//Config配置
var Config settings

func init() {
	readTimes := 0
readFile:
	f, err := ioutil.ReadFile("./config/" + "settings.yaml")
	if err != nil {
		if readTimes >= 3 {
			panic("读取配置文件错误，启动失败")
		} else {
			readTimes++
			goto readFile
		}
	}

	yaml.Unmarshal(f, &Config)
}

type env struct {
	Prod bool `yaml:"prod"`
}

// WssConf 环境设置细节
type wssConf struct {
	Address    string `yaml:"service_listen_ip"`
	Port       string `yaml:"service_listen_port"`
	SecureCert string `yaml:"service_listen_securt_cert"`
	SecureKey  string `yaml:"service_listen_securt_key"`
}

//ListendAddr 获取推送服务监听的地址
func (w wssConf) ListendAddr() string {
	return w.Address + ":" + w.Port
}

type httpConf struct {
	Port string `yaml:"service_listen_port"`
}

type mysqlDB struct {
	IP       string `yaml:"ip"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func (db *mysqlDB) Dsn() string {
	return db.Username + ":" + db.Password + "@tcp(" + db.IP + ":" + db.Port + ")/4s_wx_db?parseTime=true"
}

type mongoDB struct {
	IpPort   string `yaml:"ip_port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Options  string `yaml:"options"`
}

func (db *mongoDB) Dsn() string {
	dsn := "mongodb://"
	if len(db.Username) > 0 && len(db.Password) > 0 {
		dsn = dsn + db.Username + ":" + db.Password + "@"
	}

	dsn = dsn + db.IpPort

	if len(db.Options) > 0 {
		dsn = dsn + "?" + db.Options
	}

	return dsn
}

//settings 环境设置结构体
type settings struct {
	Wss     wssConf  `yaml:"wss_listen_service"`
	Http    httpConf `yaml:"http_listen_service"`
	MysqlDB mysqlDB  `yaml:"mysql_db"`
	MongoDB mongoDB  `yaml:"mongo_db"`
	Env     env      `yaml:"env"`
}
