package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"aggregator/pkg/storage"
	"aggregator/pkg/storage/memdb"
)

func request(api *API, method, url string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, url, nil)

	rr := httptest.NewRecorder()
	api.Router().ServeHTTP(rr, req)

	return rr
}

func Test(t *testing.T) {

	db, _ := memdb.New()
	api := New(db)

	t.Run("news_0", func(t *testing.T) {
		rr := request(api, "GET", "/news/0")

		if http.StatusOK != rr.Code {
			t.Errorf("Ожидалался код ответа %d. Получено %d\n", http.StatusOK, rr.Code)
		}

		body := rr.Body.Bytes()

		var news []storage.Post

		json.Unmarshal(body, &news)

		if len(news) != 10 {
			t.Errorf("Ожидалось получить 10 публикаций. Получили %d", len(news))
		}
	})

	t.Run("news_5", func(t *testing.T) {
		rr := request(api, "GET", "/news/5")

		if http.StatusOK != rr.Code {
			t.Errorf("Ожидалался код ответа %d. Получено %d\n", http.StatusOK, rr.Code)
		}

		body := rr.Body.Bytes()

		var news []storage.Post

		json.Unmarshal(body, &news)

		if len(news) != 5 {
			t.Errorf("Ожидалось получить 5 публикаций. Получили %d", len(news))
		}
	})

	t.Run("news_20", func(t *testing.T) {
		rr := request(api, "GET", "/news/20")

		if http.StatusOK != rr.Code {
			t.Errorf("Ожидалался код ответа %d. Получено %d\n", http.StatusOK, rr.Code)
		}

		body := rr.Body.Bytes()

		var news []storage.Post

		json.Unmarshal(body, &news)

		if len(news) != 20 {
			t.Errorf("Ожидалось получить 20 публикаций. Получили %d", len(news))
		}
	})
}
