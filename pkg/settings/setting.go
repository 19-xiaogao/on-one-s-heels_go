package settings

import "time"
import "github.com/spf13/viper"

// ServerSetting use read to configs/config server struct
type ServerSetting struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type DataBaseSetting struct {
	DBType      string
	Username    string
	Password    string
	Host        string
	DBName      string
	TablePrefix string
	Charset     string
	ParseTime   bool
	MaxIdleCons int
	MaxOpenCons int
}

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.AddConfigPath("configs/")
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp: vp}, err
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
