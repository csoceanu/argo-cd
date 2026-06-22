package controller

import (
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

// SyncRateLimiter enforces a per-application minimum interval between sync
// operations. This prevents a single application from consuming excessive
// controller resources when it enters a rapid sync loop (e.g., due to a
// webhook storm or a flapping health check).
//
// Configuration:
//   - MinSyncInterval: the minimum duration between consecutive syncs for the
//     same application. If a sync is requested before this interval has elapsed,
//     it is deferred until the interval expires. Set to 0 to disable (default).
//   - BurstSize: the number of syncs allowed in rapid succession before rate
//     limiting kicks in. This accommodates legitimate bursts (e.g., initial
//     deployment of multiple resources). Default is 1.
//
// The rate limiter is safe for concurrent use from multiple goroutines.
type SyncRateLimiter struct {
	mu              sync.Mutex
	minInterval     time.Duration
	burstSize       int
	appLastSync     map[string]time.Time
	appBurstCounter map[string]int
}

// NewSyncRateLimiter creates a rate limiter with the given minimum interval
// between syncs and burst size. A zero interval disables rate limiting.
func NewSyncRateLimiter(minInterval time.Duration, burstSize int) *SyncRateLimiter {
	if burstSize < 1 {
		burstSize = 1
	}
	return &SyncRateLimiter{
		minInterval:     minInterval,
		burstSize:       burstSize,
		appLastSync:     make(map[string]time.Time),
		appBurstCounter: make(map[string]int),
	}
}

// Allow checks whether a sync operation for the given application should
// proceed. It returns true if the sync is allowed, false if it should be
// deferred. When false is returned, the caller should skip the sync and
// allow the next resync cycle to pick it up.
func (r *SyncRateLimiter) Allow(appKey string) bool {
	if r == nil || r.minInterval <= 0 {
		return true
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	lastSync, exists := r.appLastSync[appKey]

	if !exists {
		r.appLastSync[appKey] = now
		r.appBurstCounter[appKey] = 1
		return true
	}

	elapsed := now.Sub(lastSync)

	// Reset burst counter if enough time has passed
	if elapsed >= r.minInterval {
		r.appLastSync[appKey] = now
		r.appBurstCounter[appKey] = 1
		return true
	}

	// Within the interval — check burst allowance
	counter := r.appBurstCounter[appKey]
	if counter < r.burstSize {
		r.appBurstCounter[appKey] = counter + 1
		r.appLastSync[appKey] = now
		return true
	}

	log.WithField("application", appKey).
		WithField("elapsed", elapsed).
		WithField("minInterval", r.minInterval).
		Info("Sync rate limited: too many syncs in short interval")

	return false
}

// Reset removes rate limiting state for an application. Call this when an
// application is deleted to prevent memory leaks.
func (r *SyncRateLimiter) Reset(appKey string) {
	if r == nil {
		return
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.appLastSync, appKey)
	delete(r.appBurstCounter, appKey)
}

// Stats returns the number of applications currently tracked by the rate limiter.
func (r *SyncRateLimiter) Stats() int {
	if r == nil {
		return 0
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	return len(r.appLastSync)
}
