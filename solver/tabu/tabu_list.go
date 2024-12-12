package tabu

import (
	"container/list"
	"fmt"
	"hash/fnv"

	"github.com/kruily/go-timefold-solver/solver/api"
)

type TabuList struct {
	minSize     int
	maxSize     int
	currentSize int
	items       *list.List
	lookup      map[string]struct{}
	hash        func(api.IMove) (string, error)
}

func NewTabuList(minSize, maxSize int, hash func(api.IMove) (string, error)) *TabuList {
	if minSize < 0 {
		minSize = 1
	}
	if maxSize < minSize {
		maxSize = minSize
	}
	return &TabuList{
		minSize: minSize,
		maxSize: maxSize,
		items:   list.New(),
		lookup:  make(map[string]struct{}),
		hash:    hash,
	}
}

func (t *TabuList) Add(move api.IMove, iteration int, score api.IScore) error {
	if t.items == nil {
		t.items = list.New()
	}
	if t.lookup == nil {
		t.lookup = make(map[string]struct{})
	}
	hash, err := t.hash(move)
	if err != nil {
		return fmt.Errorf("failed to hash move: %w", err)
	}
	if _, ok := t.lookup[hash]; ok {
		t.updateItem(hash, iteration, score)
		return nil
	}
	if t.items.Len() >= t.currentSize {
		oldest := t.items.Remove(t.items.Front()).(*TabuItem)
		delete(t.lookup, oldest.hash)
	}
	item := &TabuItem{hash: hash, iteration: iteration, score: score}
	t.items.PushBack(item)
	t.lookup[hash] = struct{}{}
	return nil
}

func (t *TabuList) adjustSize(improvement float64) {
	if improvement > 0.3 {
		t.currentSize = max(t.minSize, t.currentSize-1)
	} else if improvement < -0.3 {
		t.currentSize = min(t.maxSize, t.currentSize+1)
	}
}

func (t *TabuList) Contains(move api.IMove) (bool, error) {
	hash, err := t.hash(move)
	if err != nil {
		return false, fmt.Errorf("failed to hash move: %w", err)
	}
	_, ok := t.lookup[hash]
	return ok, nil
}

func DefaultMoveHash(move api.IMove) (string, error) {
	h := fnv.New64a()
	switch m := move.(type) {
	case interface{ HashString() string }:
		return m.HashString(), nil
	default:
		return fmt.Sprintf("%x", h.Sum64()), nil
	}
}

func (t *TabuList) updateItem(hash string, iteration int, score api.IScore) {
	for e := t.items.Front(); e != nil; e = e.Next() {
		item := e.Value.(*TabuItem)
		if item.hash == hash {
			item.iteration = iteration
			item.score = score
			return
		}
	}
}

func (t *TabuList) Clear() {
	t.items.Init()
	t.lookup = make(map[string]struct{})
	t.currentSize = t.minSize
}
