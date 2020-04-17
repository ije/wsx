package rex

import (
	"io"

	"github.com/ije/rex/session"
	"golang.org/x/crypto/acme/autocert"
)

// Config contains context options.
type Config struct {
	SendError    bool
	ErrorType    string
	Logger       Logger
	AccessLogger Logger
	SIDStore     session.SIDStore
	SessionPool  session.Pool
}

// ServerConfig contains options to run the REX server.
type ServerConfig struct {
	Port           uint16    `json:"port"`
	TLS            TLSConfig `json:"tls"`
	ReadTimeout    uint32    `json:"readTimeout"`
	WriteTimeout   uint32    `json:"writeTimeout"`
	MaxHeaderBytes uint32    `json:"maxHeaderBytes"`
}

// TLSConfig contains options to support https.
type TLSConfig struct {
	Port         uint16        `json:"port"`
	CertFile     string        `json:"certFile"`
	KeyFile      string        `json:"keyFile"`
	AutoTLS      AutoTLSConfig `json:"autotls"`
	AutoRedirect bool          `json:"autoRedirect"`
}

// AutoTLSConfig contains options to support autocert by Let's Encrypto SSL.
type AutoTLSConfig struct {
	AcceptTOS bool           `json:"acceptTOS"`
	Hosts     []string       `json:"hosts"`
	CacheDir  string         `json:"cacheDir"`
	Cache     autocert.Cache `json:"-"`
}

// CORSOptions contains options to CORS.
type CORSOptions struct {
	AllowOrigin      string
	AllowMethods     []string
	AllowHeaders     []string
	ExposeHeaders    []string
	AllowCredentials bool
	MaxAge           int // in seconds
}

// Template is a template contains an Execute method.
type Template interface {
	Execute(wr io.Writer, data interface{}) error
}

// Logger is a logger contains Println and Printf methods.
type Logger interface {
	Println(v ...interface{})
	Printf(format string, v ...interface{})
}
