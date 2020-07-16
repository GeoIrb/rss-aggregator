# rss-aggregator

Сервис Rss-агрегатор
Функционал:

1. Добавление отслеживания RSS ленты
2. Остановка отслеживания RSS ленты
3. Получение новостей по части заголовка

[API](cmd/API.md)

Создание базы данных

```
CREATE TABLE public.news (
	id serial NOT NULL,
	title varchar NOT NULL,
	pubDate timestamp NOT NULL,
	CONSTRAINT news_pkey PRIMARY KEY (id)
);
```
# ToDo
 * [x] Добавление отслеживания RSS ленты
 * [x] Остановка отслеживания RSS ленты
 * [x] Получение новостей по части заголовка
 * [ ] Клиент
 * [ ] Тест
