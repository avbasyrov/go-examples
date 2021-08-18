package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	// Форматирование даты-времени:
	fmt.Println("1)", now)                                     // 2021-08-18 06:04:48.703048978 +0300 MSK m=+0.000071073
	fmt.Println("2)", now.Format(time.RFC3339))                // 2021-08-18T06:04:48+03:00
	fmt.Println("3)", now.Format("02-01-2006 15:04:05Z07:00")) // 18-08-2021 06:04:48+03:00

	// Таймзона:
	timezone := now.Location()

	// Парсинг даты-времени в заданной таймзоне
	parsedTime, err := time.ParseInLocation("02-01-2006 15:04:05", "18-08-2021 06:06:30", timezone)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(parsedTime)

	adjusted := now.Add(time.Hour)            // Добавляем 1 час к текущему времени
	adjusted = adjusted.Add(-2 * time.Minute) // Отнимаем 2 минуты
	adjusted = adjusted.AddDate(0, 0, 2)      // Добавляем 2 суток

	// Получаем начало дня
	year, month, day := now.Date()
	fmt.Println("Начало дня:", time.Date(year, month, day, 0, 0, 0, 0, timezone))

	// Перебираем диапазон времени:
	tomorrow := now.Add(2 * time.Hour)
	for t := now; t.Before(tomorrow); t = t.Add(time.Hour) {
		fmt.Println(t)
	}
}
