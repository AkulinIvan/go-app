package book

import (
    "fmt"
    "io"
    "net/http"
    "os"
)

// Book представляет структуру для хранения информации о книге
type Book struct {
	Title   string // Название книги
	Author  string // Автор книги
	Pages   int    // Количество страниц
	Rating  float32 // Рейтинг книги
}

func describeBook(book Book) {
	fmt.Printf("Название: %s\nАвтор: %s\nКоличество страниц: %d\nРейтинг: %.1f\n",
	book.Title,
	book.Author,
	book.Pages,
	book.Rating)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
    // Проверяем метод запроса
    if r.Method != "POST" {
        http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
        return
    }
    // Получаем файл из формы
    file, handler, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Ошибка при получении файла", http.StatusInternalServerError)
        return
    }
    defer file.Close()
    // Создаём файл на сервере
    dst, err := os.Create("uploaded_files/" + handler.Filename)
    if err != nil {
        http.Error(w, "Ошибка при создании файла на сервере", http.StatusInternalServerError)
        return
    }
    defer dst.Close()
    // Копируем содержимое файла на сервер
    if _, err = io.Copy(dst, file); err != nil {
        http.Error(w, "Ошибка при копировании файла", http.StatusInternalServerError)
        return
    }
    // Возвращаем сообщение об успешной загрузке файла
    fmt.Fprintf(w, "Файл %s успешно загружен", handler.Filename)
}
