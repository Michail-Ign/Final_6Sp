package server

import (
	"log"
	"net/http"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type ServerStr struct {
	Addr         string
	Handler      *http.ServeMux
	ErrorLog     *log.Logger
	ReadTimeout  int16
	WriteTimeout int16
	IdleTimeout  int16
}

func CreateRouter(log *log.Logger) ServerStr {

	//2 HHTP роутер
	mux := http.NewServeMux()
	//3. Регистрация хендлеров в http-роутере.
	mux.HandleFunc("/", handlers.HandleRoot)
	mux.HandleFunc("/upload", handlers.HandleUpload)

	//4.Создание ссылки на сервер
	var pServer ServerStr
	pServer.Addr = ":8080"
	pServer.Handler = mux
	pServer.ErrorLog = log
	pServer.ReadTimeout = 5
	pServer.WriteTimeout = 10
	pServer.IdleTimeout = 15

	return pServer
}
