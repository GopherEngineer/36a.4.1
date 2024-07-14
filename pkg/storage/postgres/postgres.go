// Пакет для работы с БД приложения "Новостной агрегатор".
package postgres

import (
	"aggregator/pkg/storage"
	"context"
	"errors"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Хранилище данных.
type Storage struct {
	db *pgxpool.Pool
}

// Конструктор, принимает строку подключения к БД.
func New() (*Storage, error) {
	url := os.Getenv("DB_URL")
	if url == "" {
		return nil, errors.New("переменная окружения DB_URL не задана")
	}
	db, err := pgxpool.New(context.Background(), url)
	if err != nil {
		return nil, err
	}
	s := Storage{
		db: db,
	}
	return &s, nil
}

// News возвращает список публикаций из БД.
// Принимает один параметр n - сколько публикаций нужно вернуть.
func (s *Storage) News(n int) ([]storage.Post, error) {
	// в случае n равным нулю считаем, что нужно вернуть 10 публикаций
	if n == 0 {
		n = 10
	}

	// получаем публицации из базы данных с указанным ограничение
	rows, err := s.db.Query(context.Background(), `
		SELECT 
			id,
			title,
			content,
			pub_time,
			link
		FROM news
		ORDER BY pub_time DESC
		LIMIT $1;
	`, n)
	// обязательно проверяем на ошибку
	if err != nil {
		return nil, err
	}

	// объявляем слайс для публикаций
	var news []storage.Post

	// итерирование по результату выполнения запроса
	// и сканирование каждой строки в переменную
	for rows.Next() {
		var p storage.Post

		// сохраняем полученные значение публикации в переменную
		err = rows.Scan(
			&p.ID,
			&p.Title,
			&p.Content,
			&p.PubTime,
			&p.Link,
		)
		if err != nil {
			return nil, err
		}

		// добавляем переменной в массив результатов
		news = append(news, p)
	}

	// ВАЖНО не забыть проверить rows.Err()
	return news, rows.Err()
}

// SaveNews сохраняет массив публикаций в базу данных.
// Уникальность публикации проверяется по полю "Link".
func (s *Storage) SaveNews(news []storage.Post) error {
	// проходим по слайсу публикаций и создаем новую публикацию в базе данных
	for _, post := range news {
		_, err := s.db.Exec(context.Background(), `
			INSERT INTO news (title, content, pub_time, link)
			VALUES ($1, $2, $3, $4);
		`,
			post.Title,
			post.Content,
			post.PubTime,
			post.Link,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
