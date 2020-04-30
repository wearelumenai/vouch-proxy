package conf

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"path/filepath"

	"github.com/jinzhu/configor"
)

// Clusauth configuration
type Clusauth struct {
	Endpoint string `env:"CLUSAUTH_ENDPOINT" default:"http://clusauth:8003"`
	Certfile string `env:"CLUSAUTH_CERTFILE"`
	Keyfile  string `env:"CLUSAUTH_KEYFILE"`
	Sentry   string `env:"CLUSAUTH_SENTRY"`
	Timeout  int    `default:"180" env:"CLUSAUTH_TIMEOUT"`
	Domains  []string
	Tokens   []string
	Secret   string `default:"clusAUTH" env:"CLUSAUTH_SECRET"`
	Debug    bool   `default:"false" env:"CLUSAUTH_DEBUG"`
}

// Vouch configuration
type Vouch struct {
	Endpoint string `env:"CLUSAUTH_VOUCH_ENDPOINT" default:"http://clusauth-vouch:9090"`
	Port     int    `env:"CLUSAUTH_VOUCH_PORT" default:"9090"`
	Domains  []string
}

// Conf global configuration
type Conf struct {
	Clusauth
	Vouch
}

// SetDefaultValues for vouch conf
func (vouch *Vouch) SetDefaultValues() (err error) {
	if vouch.Endpoint == "" {
		vouch.Endpoint = "http://clusauth-vouch:9090"
	}
	err = vouch.Ping()
	return
}

// Ping vouch
func (vouch Vouch) Ping() error {
	return Ping(vouch.Endpoint, "healthcheck")
}

// OnReload callback
func OnReload(config interface{}) {

}

// Load configuration from configs directory
func Load(configs string) (conf Conf, err error) {
	var exts = []string{
		"yml", "yaml", "json", "toml",
	}
	var found = false
	var config = configor.Config{
		AutoReloadCallback: OnReload,
	}
	var loader *configor.Configor
	for _, ext := range exts {
		var filename = fmt.Sprintf("config.%s", ext)
		var path = filepath.Join(configs, filename)
		if _, er := os.Stat(path); er == nil {
			loader = configor.New(&config)
			err = loader.Load(&conf, path)
			if err == nil {
				log.Printf("%s found\n", filename)
				found = true
			} else {
				log.Println(err)
			}
			break
		}
	}
	if !found {
		log.Println("config file not found")
	}
	if err == nil {
		conf.SetDefaultValues()
		if loader != nil {
			loader.Debug = conf.Clusauth.Debug
		}
	}
	return
}

// SetDefaultValues for conf
func (conf *Conf) SetDefaultValues() (err error) {
	conf.Clusauth.SetDefaultValues()
	conf.Vouch.SetDefaultValues()

	return
}

// ErrNoVouchProxy raised if vouch proxy is unavailable
var ErrNoVouchProxy = errors.New("No vouch proxy")

// SetDefaultValues for clusauth
func (conf *Clusauth) SetDefaultValues() (err error) {
	if conf.Timeout == 0 {
		conf.Timeout = 180
		var envTimeout = os.Getenv("CLUSAUTH_TIMEOUT")
		if envTimeout != "" {
			var timeout, errTimeout = strconv.Atoi(envTimeout)
			if errTimeout == nil {
				conf.Timeout = timeout
			}
		}
	}
	if conf.Tokens == nil {
		var tokens = os.Getenv("CLUSAUTH_TOKENS")
		if tokens != "" {
			conf.Tokens = strings.Split(tokens, ",")
		}
	}
	for i, token := range conf.Tokens {
		conf.Tokens[i] = strings.TrimSpace(token)
	}
	if conf.Domains == nil {
		var domains = os.Getenv("CLUSAUTH_DOMAINS")
		if domains != "" {
			conf.Domains = strings.Split(domains, ",")
		}
	}
	for i, domain := range conf.Domains {
		conf.Domains[i] = strings.TrimSpace(domain)
	}
	if conf.Endpoint == "" {
		conf.Endpoint = "http://clusauth:8003"
	}

	return
}
