package slidingWindow

import (
	"testing"
	"time"
)

func BenchmarkLimiter_Allow(b *testing.B) {
	limiter := NewLimiter(time.Second, time.Millisecond, 100)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = limiter.Allow()
	}
}

func BenchmarkLimiter_Take(b *testing.B) {
	limiter := NewLimiter(time.Second, time.Millisecond, 100)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		limiter.Take()
	}
}

func BenchmarkLimiter_AllowTake(b *testing.B) {
	limiter := NewLimiter(time.Second, time.Millisecond, 1000)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if limiter.Allow() {
			limiter.Take()
		}
	}
}

func BenchmarkLimiter_moveToNextPosition(b *testing.B) {
	limiter := NewLimiter(time.Millisecond, time.Millisecond, 10)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		limiter.moveToNextPosition()
	}
}
