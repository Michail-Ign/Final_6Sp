package handlers

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
	"github.com/go-delve/delve/service"
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
		http.Error(w, "ошибка при получении файла", http.StatusInternalServerError)
		return
	}
	defer myfile.Close()

	// создаем мапу для хранения преобразованных строк из файла
	str_res := ""

	// создаем сканер для построчного чтения файла
	scanner := bufio.NewScanner(myfile)

	// читаем файл построчно
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		//Переконвертация каждой строки файла и зпоминае в мапе
		str_res = str_res + service.ToConvert(line) + "\n"
		// декодируем JSON-строку в структуру LogEntry

		//if err != nil {
		//	log.Printf("ошибка парсинга: %v. Строка: %s\n", err, line)
		//	continue
		//}
		//i++
	}

	// создаём локальный файл для записи данных
	name_file := time.Now().UTC().String()
	resFile, err := os.Create(name_file)
	if err != nil {
		//log.Fatal(err)
		http.Error(w, "ошибка при cоздании локального файла", http.StatusInternalServerError)
		return
	}
	defer resFile.Close()

	//for line_w, _ := range map_data {
	_, err = fmt.Fprintf(resFile, "%s", str_res)
	if err != nil {
		//log.Fatal(err)
		http.Error(w, "ошибка при сохранении результата в локальный файл", http.StatusInternalServerError)
		return
	}
	//}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(str_res))
}
