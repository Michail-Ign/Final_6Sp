package handlers

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

var tpl = template.Must(template.ParseFiles("index.html"))

// handleRoot хендлер(эндпоинт - "/"), который возвращает HTML из файла index.html
func HandleRoot(w http.ResponseWriter, r *http.Request) {

	tpl.Execute(w, nil)
}

// handleUpload обрабатывает POST эндпоинт - /upload
func HandleUpload(w http.ResponseWriter, r *http.Request) {

	//Получение файла из формы
	myfile, _, err := r.FormFile("myFile")
	if err != nil {

		str := fmt.Sprintf("ошибка(%v) при получении файла", err)
		http.Error(w, str, http.StatusInternalServerError)

		return
	}
	defer myfile.Close()

	str_res := ""

	scanner := bufio.NewScanner(myfile)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}
		//Переконвертация каждой строки файла и зпоминае в мапе
		str_res = str_res + service.ToConvert(line) + "\n"
	}

	// создаём локальный файл для записи данных
	name_file := time.Now().UTC().String() + ".txt"
	name_file = strings.Replace(name_file, ":", "_", -1)

	resFile, err := os.Create(name_file)
	if err != nil {

		str := fmt.Sprintf("ошибка(%v) при cоздании локального файла", err)
		http.Error(w, str, http.StatusInternalServerError)

		return
	}
	defer resFile.Close()

	_, err = fmt.Fprintf(resFile, "%s", str_res)
	if err != nil {

		str := fmt.Sprintf("ошибка(%v) при сохранении результата в локальный файл", err)
		http.Error(w, str, http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(str_res))
}
