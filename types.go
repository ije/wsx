package rex

import (
	"bytes"
	"html/template"
	"io"
	"os"
	"path"
	"time"

	"golang.org/x/crypto/acme/autocert"
)

// ServerConfig contains options to run the REX server.
type ServerConfig struct {
	Host           string    `json:"host"`
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

// CORS contains options to CORS.
type CORS struct {
	AllowAllOrigins  bool
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	ExposeHeaders    []string
	AllowCredentials bool
	MaxAge           int // in seconds
}

// A ACLUser interface contains the Permissions method that returns the permission IDs
type ACLUser interface {
	Permissions() []string
}

// A Logger interface contains the Printf method.
type Logger interface {
	Printf(format string, v ...interface{})
}

// Error defines an error with status.
type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type recoverError struct {
	status  int
	message string
}

type redirecting struct {
	status int
	url    string
}

// Redirect replies to the request with a redirect to url,
// which may be a path relative to the request path.
func Redirect(url string, status int) interface{} {
	return &redirecting{status, url}
}

type contentful struct {
	name    string
	status  int
	mtime   time.Time
	content io.ReadSeeker
}

// Content replies to the request using the content in the provided ReadSeeker.
func Content(name string, mtime time.Time, r io.ReadSeeker) *contentful {
	return &contentful{name, 200, mtime, r}
}

// File replies to the request using the file content.
func File(name string) *contentful {
	fi, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			panic(&recoverError{404, "file not found"})
		}
		panic(&recoverError{500, err.Error()})
	}
	if fi.IsDir() {
		panic(&recoverError{400, "is a directory"})
	}

	file, err := os.Open(name)
	if err != nil {
		panic(&recoverError{500, err.Error()})
	}

	return &contentful{path.Base(name), 200, fi.ModTime(), file}
}

// HTML replies to the request with a html content.
func HTML(status int, html string, data interface{}) *contentful {
	if data == nil {
		return &contentful{
			name:    "index.html",
			mtime:   time.Now(),
			content: bytes.NewReader([]byte(html)),
		}
	}
	t, err := template.New("").Parse(html)
	if err != nil {
		panic(&recoverError{500, err.Error()})
	}
	buf := bytes.NewBuffer(nil)
	if err := t.Execute(buf, data); err != nil {
		panic(&recoverError{500, err.Error()})
	}
	return &contentful{
		name:    "index.html",
		status:  status,
		mtime:   time.Now(),
		content: bytes.NewReader(buf.Bytes()),
	}
}

type fs struct {
	root     string
	fallback string
}

// FS replies to the request with the contents of the file system rooted at root.
func FS(root string, fallback string) interface{} {
	fi, err := os.Lstat(root)
	if err != nil {
		panic(&recoverError{500, err.Error()})
	}
	if !fi.IsDir() {
		panic(&recoverError{500, "FS root is not a directory"})
	}
	return &fs{root, fallback}
}
