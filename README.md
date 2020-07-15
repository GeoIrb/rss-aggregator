# tochka-test
----

Сервис Агрегатор-новостей
Функционал:

1. Добавление отслеживания RSS ленты
2. Остановка отслеживания RSS ленты
3. Получение новостей

[API](cmd/API.md)

Создание базы данных

```
CREATE TABLE public.news (
	id serial NOT NULL,
	title varchar NOT NULL,
	pubdate timestamp NOT NULL,
	CONSTRAINT news_pkey PRIMARY KEY (id)
);
```