package setting

type ServerSettings struct {
	NetWork      string
	Port         int
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
	Mode         string
	PIDFile      string
}

type LoggerSettings struct {
	RootPath    string
	Level       string
}

// MysqlSettings MysqlSettingS defines for connecting mysql.
type MysqlSettings struct {
	Host               string
	Port               int
	DataBase           string
	UserName           string
	Password           string
	Charset            string
	PoolNum            int
	Loc                string
	ParseTime          bool
	MultiStatements    bool
	ConnMaxLifeSecond  uint
	MaxIdleConns       int
}

// RedisSettings RedisSettingS defines for connecting redis.
type RedisSettings struct {
	Host        string
	Port        int
	Password    string
	DataBase    int
	PoolNum     int
}

