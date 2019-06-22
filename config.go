package rex

import (
	"golang.org/x/crypto/acme/autocert"
)

type Config struct {
	Port           uint16      `json:"port"`
	HTTPS          HTTPSConfig `json:"https"`
	ReadTimeout    uint32      `json:"readTimeout"`
	WriteTimeout   uint32      `json:"writeTimeout"`
	MaxHeaderBytes uint32      `json:"maxHeaderBytes"`
	Debug          bool        `json:"debug"`
	Logger         Logger      `json:"-"`
	AccessLogger   Logger      `json:"-"`
}

type HTTPSConfig struct {
	Port     uint16        `json:"port"`
	CertFile string        `json:"certFile"`
	KeyFile  string        `json:"keyFile"`
	AutoTLS  AutoTLSConfig `json:"autotls"`
}

type AutoTLSConfig struct {
	Enable   bool           `json:"enable"`
	Hosts    []string       `json:"hosts"`
	CacheDir string         `json:"cacheDir"`
	Cache    autocert.Cache `json:"-"`
}

type CORSOptions struct {
	AllowOrigin      string
	AllowMethods     []string
	AllowHeaders     []string
	ExposeHeaders    []string
	AllowCredentials bool
	MaxAge           int // in seconds
}

type Logger interface {
	Println(v ...interface{})
	Printf(format string, v ...interface{})
}
