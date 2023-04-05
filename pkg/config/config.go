package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// import cycle -> app not compile
// this package is imported by other parts of the app
// but it doesn't need to import anything else
// only using things from library

// App config holds the application config
type AppConfig struct {
	//tamplate cache
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
