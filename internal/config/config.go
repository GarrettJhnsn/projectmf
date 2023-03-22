package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// AppConfig Holds Application Config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	ErrorLog      *log.Logger
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
