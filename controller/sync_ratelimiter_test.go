package controller

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSyncRateLimiter_AllowFirstSync(t *testing.T) {
	rl := NewSyncRateLimiter(10*time.Second, 1)
	assert.True(t, rl.Allow("default/my-app"))
}

func TestSyncRateLimiter_BlocksRapidSyncs(t *testing.T) {
	rl := NewSyncRateLimiter(10*time.Second, 1)

	assert.True(t, rl.Allow("default/my-app"))
	assert.False(t, rl.Allow("default/my-app"))
}

func TestSyncRateLimiter_AllowsBurst(t *testing.T) {
	rl := NewSyncRateLimiter(10*time.Second, 3)

	assert.True(t, rl.Allow("default/my-app"))
	assert.True(t, rl.Allow("default/my-app"))
	assert.True(t, rl.Allow("default/my-app"))
	assert.False(t, rl.Allow("default/my-app"))
}

func TestSyncRateLimiter_IndependentApps(t *testing.T) {
	rl := NewSyncRateLimiter(10*time.Second, 1)

	assert.True(t, rl.Allow("default/app-a"))
	assert.True(t, rl.Allow("default/app-b"))
	assert.False(t, rl.Allow("default/app-a"))
	assert.True(t, rl.Allow("default/app-c"))
}

func TestSyncRateLimiter_AllowsAfterInterval(t *testing.T) {
	rl := NewSyncRateLimiter(50*time.Millisecond, 1)

	assert.True(t, rl.Allow("default/my-app"))
	assert.False(t, rl.Allow("default/my-app"))

	time.Sleep(60 * time.Millisecond)
	assert.True(t, rl.Allow("default/my-app"))
}

func TestSyncRateLimiter_DisabledWhenZeroInterval(t *testing.T) {
	rl := NewSyncRateLimiter(0, 1)

	for i := 0; i < 100; i++ {
		assert.True(t, rl.Allow("default/my-app"))
	}
}

func TestSyncRateLimiter_NilSafe(t *testing.T) {
	var rl *SyncRateLimiter

	assert.True(t, rl.Allow("default/my-app"))
	rl.Reset("default/my-app")
	assert.Equal(t, 0, rl.Stats())
}

func TestSyncRateLimiter_Reset(t *testing.T) {
	rl := NewSyncRateLimiter(10*time.Second, 1)

	assert.True(t, rl.Allow("default/my-app"))
	assert.False(t, rl.Allow("default/my-app"))

	rl.Reset("default/my-app")
	assert.True(t, rl.Allow("default/my-app"))
}

func TestSyncRateLimiter_Stats(t *testing.T) {
	rl := NewSyncRateLimiter(10*time.Second, 1)

	assert.Equal(t, 0, rl.Stats())
	rl.Allow("default/app-a")
	rl.Allow("default/app-b")
	assert.Equal(t, 2, rl.Stats())
	rl.Reset("default/app-a")
	assert.Equal(t, 1, rl.Stats())
}

func TestSyncRateLimiter_ConcurrentAccess(t *testing.T) {
	rl := NewSyncRateLimiter(10*time.Second, 5)

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			rl.Allow(fmt.Sprintf("default/app-%d", id%5))
		}(i)
	}
	wg.Wait()

	assert.LessOrEqual(t, rl.Stats(), 5)
}

func TestSyncRateLimiter_MinBurstSize(t *testing.T) {
	rl := NewSyncRateLimiter(10*time.Second, 0)
	assert.Equal(t, 1, rl.burstSize)

	rl = NewSyncRateLimiter(10*time.Second, -5)
	assert.Equal(t, 1, rl.burstSize)
}
