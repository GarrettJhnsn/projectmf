package render

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/garrettjhnsn/projectmf/internal/config"
	"github.com/garrettjhnsn/projectmf/internal/models"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {
	//What I am putting in the session
	gob.Register(models.ConsultationRequest{})

	testApp.InProduction = false

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	testApp.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	testApp.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}

type customWriter struct{}

func (tw *customWriter) Header() http.Header {
	var h http.Header
	return h
}

func (tw *customWriter) WriteHeader(i int) {}

func (tw *customWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}
