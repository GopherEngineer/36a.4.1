// API приложения "Новостной агрегатор".
package api

import (
	"aggregator/pkg/storage"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type API struct {
	db     storage.Interface
	router *mux.Router
}

// Конструктор API.
func New(db storage.Interface) *API {
	api := API{
		db:     db,
		router: mux.NewRouter(),
	}
	api.endpoints()
	return &api
}

// Router возвращает маршрутизатор для использования
// в качестве аргумента HTTP-сервера.
func (api *API) Router() *mux.Router {
	return api.router
}

// регистрация методов API в маршрутизаторе запросов
func (api *API) endpoints() {
	// получить n последних публикаций
	api.router.HandleFunc("/news/{n}", api.handler).Methods(http.MethodGet)
	// веб-приложение
	api.router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./web/public"))))
}

// обработчик HTTP запросов получения публикаций
func (api *API) handler(w http.ResponseWriter, r *http.Request) {
	// подготовка HTTP заголовка ответа сервера, что будет возвращен JSON
	w.Header().Set("Content-Type", "application/json")
	// подготовка HTTP заголовка ответа сервера, что можно делать запросы с любых хостов
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// получение из запроса ограничения на количество публикаций
	n, _ := strconv.Atoi(mux.Vars(r)["n"])

	// обращние к базе данных для получения публикаций
	news, err := api.db.News(n)
	// если произошла ошибка чтения из базы данных, то завершаем
	// обращение к серверу и указываем, что произошла
	// внутренняя ошибка сервера с текстом ошибки
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// если публикаций нет, то возвращаем пустой массив в JSON спецификации
	if len(news) == 0 {
		w.Write([]byte("[]"))
		return
	}

	// пишем ответ кодирую публикации
	json.NewEncoder(w).Encode(news)
}
