package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

var (
	global *Config
)

// LoadGlobalConfig 加载全局配置
func LoadGlobalConfig(fpath string) error {
	c, err := ParseConfig(fpath)
	if err != nil {
		return err
	}
	global = c
	return nil
}

// GetGlobalConfig 获取全局配置
func GetGlobalConfig() *Config {
	if global == nil {
		return &Config{}
	}
	return global
}

// ParseConfig 解析配置文件
func ParseConfig(fpath string) (*Config, error) {
	var c Config
	_, err := toml.DecodeFile(fpath, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// Config 配置参数
type Config struct {
	RunMode         string   `toml:"run_mode"`
	CasbinModelConf string   `toml:"casbin_model_conf"`
	WWW             string   `toml:"www"`
	Swagger         string   `toml:"swagger"`
	Store           string   `toml:"store"`
	AllowInitMenu   bool     `toml:"allow_init_menu"`
	EnableCasbin    bool     `toml:"enable_casbin"`
	Log             Log      `toml:"log"`
	Mail            Mail     `toml:"mail"`
	Root            Root     `toml:"root"`
	JWTAuth         JWTAuth  `toml:"jwt_auth"`
	HTTP            HTTP     `toml:"http"`
	CORS            CORS     `toml:"cors"`
	Gorm            Gorm     `toml:"gorm"`
	MySQL           MySQL    `toml:"mysql"`
	Postgres        Postgres `toml:"postgres"`
	Sqlite3         Sqlite3  `toml:"sqlite3"`
}

// Log 日志配置参数
type Log struct {
	Level      int    `toml:"level"`
	Format     string `toml:"format"`
	Output     string `toml:"output"`
	OutputFile string `toml:"output_file"`
}

type Mail struct {
	Host        string `toml:"host"`
	Port        int    `toml:"port"`
	Addr        string `toml:"addr"`
	Password    string `toml:"password"`
	Mailform    string `toml:"mailform"`
	FrontendURL string `toml:"frontend_url"`
}

// Root root用户
type Root struct {
	UserName string `toml:"user_name"`
	Password string `toml:"password"`
	RealName string `toml:"real_name"`
}

// JWTAuth 用户认证
type JWTAuth struct {
	SigningMethod string `toml:"signing_method"`
	SigningKey    string `toml:"signing_key"`
	Expired       int    `toml:"expired"`
	Store         string `toml:"store"`
	FilePath      string `toml:"file_path"`
}

// HTTP http配置参数
type HTTP struct {
	Host            string `toml:"host"`
	Port            int    `toml:"port"`
	ShutdownTimeout int    `toml:"shutdown_timeout"`
}

// CORS 跨域请求配置参数
type CORS struct {
	Enable           bool     `toml:"enable"`
	AllowOrigins     []string `toml:"allow_origins"`
	AllowMethods     []string `toml:"allow_methods"`
	AllowHeaders     []string `toml:"allow_headers"`
	AllowCredentials bool     `toml:"allow_credentials"`
	MaxAge           int      `toml:"max_age"`
}

// Gorm gorm配置参数
type Gorm struct {
	Debug        bool   `toml:"debug"`
	DBType       string `toml:"db_type"`
	MaxLifetime  int    `toml:"max_lifetime"`
	MaxOpenConns int    `toml:"max_open_conns"`
	MaxIdleConns int    `toml:"max_idle_conns"`
	TablePrefix  string `toml:"table_prefix"`
}

// MySQL mysql配置参数
type MySQL struct {
	Host       string `toml:"host"`
	Port       int    `toml:"port"`
	User       string `toml:"user"`
	Password   string `toml:"password"`
	DBName     string `toml:"db_name"`
	Parameters string `toml:"parameters"`
}

// DSN 数据库连接串
func (a MySQL) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		a.User, a.Password, a.Host, a.Port, a.DBName, a.Parameters)
}

// Postgres postgres配置参数
type Postgres struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DBName   string `toml:"db_name"`
}

// DSN 数据库连接串
func (a Postgres) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
		a.Host, a.Port, a.User, a.DBName, a.Password)
}

// Sqlite3 sqlite3配置参数
type Sqlite3 struct {
	Path string `toml:"path"`
}

// DSN 数据库连接串
func (a Sqlite3) DSN() string {
	return a.Path
}
