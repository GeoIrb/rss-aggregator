# API

### Добавление отслеживания RSS ленты

Метод: `POST`

URI: `/tracking`

Body:
```json
{
    "url": "",
    "format": ""
}
```

> url - url RSS ленты
>
> format - формат даты в ленте

### Остановка отслеживания RSS ленты

Метод: `DELETE`

URI: `/tracking`

Body:
```json
{
    "url": ""
}
```

> url - url RSS ленты

### Получение новостей

Метод: `GET`

URI: `/news?title={title}`

> title - название новости или его часть