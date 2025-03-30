package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	//1 - Создаем логгер
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// 2 -  создать сервер с помощью вашей функции из пакета server, и запустить его.
	myserver := server.CreateRouter(errorLog)
	// запускаем сервер
	if err := http.ListenAndServe(myserver.Addr, myserver.Handler); err != nil {
		// 3-  при запуске сервера возникают ошибки, выведите её с помощью логгера на уровне Fatal

		errorLog.Fatal("Ошибка при запуске сервера: %s", err.Error())
		return
	}

}
