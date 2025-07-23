package main

import (
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	config, err := loadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Ошибка загрузки конфиг файла: %v", err)
	}

	listenerAddr := config.Listener.Addr
	cacheSize := config.Cache.Size

	cache := NewCache(cacheSize)

	// Создаем сервер Echo
	e := echo.New()

	// регистрируем маршрут
	e.GET("/:domain", handleRequest(cache))

	// Запускаем сервер
	log.Printf("Starting proxy on %s", listenerAddr)
	log.Fatal(e.Start(listenerAddr))
}

// 1. добавить авторизацию на сервере (бейсик) или JWT (Json Web Token) через gRPC
// (попробовать сделать со временем жизни токена refresh token)
// 2. ДОП. изменить хандлер, сделать его объектом, не функцией
// 3. ДОП. добавить тесты
// 4. ДОП. graceful-shutdown
// 5. ДОП. banchmark тесты для checkQueueSize
// 6. ДОП. в гитигнор .idea, .pkg
// 7. ДОП. завернуть в docker-compose
// 8. ДОП. реализовать аннотациями свагер (посмотреть на кодогенерацию)
// 9. ДОП. распихать по папкам
// 10. ДОП. в кеше, вместо мапы найти сторонее key-value хранилище (elastic, tarantool, redis)
