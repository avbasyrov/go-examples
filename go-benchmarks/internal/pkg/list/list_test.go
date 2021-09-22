package list

import (
	"strconv"
	"testing"
)

func BenchmarkMap_Push(b *testing.B) {
	var keys []string

	// Заранее подготовим N ключей, чтобы при замере скорости не учитывалось время вызова strconv.Itoa(n)
	for n := 0; n < b.N; n++ {
		key := strconv.Itoa(n)
		keys = append(keys, key)
	}

	m := NewMapList()

	b.ReportAllocs() // Включаем опцию показа аллокаций памяти
	b.ResetTimer()   // Сбрасываем счетчик длительности тесты (то есть отсчет пойдет с текущего места)

	for _, key := range keys {
		m.Push(key, "value")
	}
}

func BenchmarkList_Push(b *testing.B) {
	var keys []string

	// Заранее подготовим N ключей, чтобы при замере скорости не учитывалось время вызова strconv.Itoa(n)
	for n := 0; n < b.N; n++ {
		key := strconv.Itoa(n)
		keys = append(keys, key)
	}

	s := NewSliceList()

	b.ReportAllocs() // Включаем опцию показа аллокаций памяти
	b.ResetTimer()   // Сбрасываем счетчик длительности тесты (то есть отсчет пойдет с текущего места)

	for _, key := range keys {
		s.Push(key, "value")
	}
}
