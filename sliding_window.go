package slidingWindow

import (
	"time"
)

type Limiter struct {
	size             time.Duration
	granularity      time.Duration
	limit            uint64
	currentPosition  int
	previousPosition int
	lastUsageTime    time.Time
	data             []uint64
	counter          uint64
}

func NewLimiter(size time.Duration, granularity time.Duration, limit uint64) *Limiter {
	return &Limiter{
		size:        size,
		granularity: granularity,
		limit:       limit,
		data:        make([]uint64, size/granularity),
	}
}

func (l *Limiter) Allow() bool {
	l.moveToNextPosition()
	if l.counter >= l.limit {
		return false
	}

	return true
}

func (l *Limiter) moveToNextPosition() {
	currentIndex := time.Now().UnixNano() % int64(l.size) / int64(l.granularity)
	if l.currentPosition == int(currentIndex) {
		return
	}
	now := time.Now()
	l.previousPosition = l.currentPosition
	l.currentPosition = int(currentIndex)

	sub := now.Sub(l.lastUsageTime)
	if sub > l.size {
		l.data = make([]uint64, l.size/l.granularity)
		l.counter = 0
	} else {
		if l.previousPosition < l.currentPosition {
			for i := l.previousPosition + 1; i <= l.currentPosition; i++ {
				l.counter -= l.data[i]
				l.data[i] = 0
			}
		} else {
			for i := 0; i <= l.currentPosition; i++ {
				l.counter -= l.data[i]
				l.data[i] = 0
			}
			for i := l.previousPosition + 1; i < len(l.data); i++ {
				l.counter -= l.data[i]
				l.data[i] = 0
			}
		}
	}
	l.lastUsageTime = now
}

func (l *Limiter) Take() {
	l.moveToNextPosition()
	l.data[l.currentPosition]++
	l.counter++
}

func (l *Limiter) count() (total uint64) {
	return l.counter
}
