package list_ints

import "testing"

func BenchmarkList_Push_Parallel(b *testing.B) {
	l := &List{}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Push(1)
		}
	})
}

func BenchmarkList_Push(b *testing.B) {
	l := &List{}

	for n := 0; n < b.N; n++ {
		l.Push(1)
	}
}

func BenchmarkList_PushChannel(b *testing.B) {
	l := New()

	for n := 0; n < b.N; n++ {
		l.PushChannel(1)
	}
}

func BenchmarkList_PushChannel_Parallel(b *testing.B) {
	l := New()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Push(1)
		}
	})
}
