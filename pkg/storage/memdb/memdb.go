// Заглушка базы данных для тестирования.
package memdb

import (
	"aggregator/pkg/storage"
	"strconv"
)

// Хранилище данных.
type Storage struct{}

// Конструктор, принимает строку подключения к БД.
func New() (*Storage, error) {
	return new(Storage), nil
}

// News возвращает список публикаций из БД.
// Принимает один параметр n - сколько публикаций нужно вернуть.
func (s *Storage) News(n int) ([]storage.Post, error) {
	// в случае n равным нулю считаем, что нужно вернуть 10 публикаций
	if n == 0 {
		n = 10
	}

	var news []storage.Post

	for i := range n {
		news = append(news, storage.Post{
			ID:      i,
			Title:   "Title",
			Content: "Content",
			Link:    strconv.Itoa(i),
		})
	}

	return news, nil
}

// SaveNews сохраняет массив публикаций в базу данных.
// Уникальность публикации проверяется по полю "Link".
func (s *Storage) SaveNews(news []storage.Post) error {
	return nil
}
