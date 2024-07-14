// Пакет для соблюдения контракта работы с базой данных.
package storage

// Публикация, получаемая из RSS.
type Post struct {
	ID      int    `json:"id"`       // номер записи
	Title   string `json:"title"`    // заголовок публикации
	Content string `json:"content"`  // содержание публикации
	PubTime int64  `json:"pub_time"` // время публикации
	Link    string `json:"link"`     // ссылка на источник
}

// Interface задаёт контракт на работу с БД.
type Interface interface {
	News(n int) ([]Post, error)
	SaveNews([]Post) error
}
