package storage

import (
	"errors"
	"sort"
	"sync"

	"github.com/adrg/strutil/metrics"
	"github.com/akl-infra/slf/v2"
	"golang.org/x/exp/maps"
)

var jaroWinkler = metrics.NewJaroWinkler()

type SyncCache struct {
	inner map[string]slf.Layout
	lock  sync.RWMutex
}

func NewSyncCache() SyncCache {
	return SyncCache{
		inner: make(map[string]slf.Layout),
		lock:  sync.RWMutex{},
	}
}

func (c *SyncCache) Get(key string) (slf.Layout, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	layout, found := c.inner[key]
	return layout, found
}

func (c *SyncCache) GetFuzzy(target string) (slf.Layout, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	if len(c.inner) == 0 {
		return slf.Layout{}, errors.New("cache not initialized")
	}

	bestMatch := ""
	bestScore := -1.0

	for name := range c.inner {
		score := jaroWinkler.Compare(name, target)
		if score > bestScore {
			bestMatch = name
			bestScore = score
		}
	}
	layout, _ := c.inner[bestMatch]
	return layout, nil
}

func (c *SyncCache) Put(key string, value slf.Layout) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.inner[key] = value
}

func (c *SyncCache) List() []string {
	c.lock.RLock()
	defer c.lock.RUnlock()

	layouts := maps.Keys(c.inner)
	sort.Sort(sort.StringSlice(layouts))
	return layouts
}
