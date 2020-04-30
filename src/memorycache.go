/*
	Copyright (c) 2020 be|ys - MIT License
	For more informations, please refer to the LICENSE file.

	Original code by goenning - MIT License. Please refer to the README file for a direct link.
*/

package main

import (
	"sync"
	"time"
)

// Item is a cached reference
type Item struct {
	Content    []byte
	Expiration int64
}

// Expired returns true if the item has expired.
func (item Item) Expired() bool {
	if item.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > item.Expiration
}

//Storage mecanism for caching strings in memory
type Storage struct {
	items map[string]Item
	mu    *sync.RWMutex
}

//NewStorage creates a new in memory storage
func NewStorage() *Storage {
	return &Storage{
		items: make(map[string]Item),
		mu:    &sync.RWMutex{},
	}
}

//Get a cached content by key
func (s Storage) Get(key string) []byte {
	s.mu.RLock()
	defer s.mu.RUnlock()

	item := s.items[key]
	if item.Expired() {
		delete(s.items, key)
		return nil
	}
	return item.Content
}

//Set a cached content by key
func (s Storage) Set(key string, content []byte, duration time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.items[key] = Item{
		Content:    content,
		Expiration: time.Now().Add(duration).UnixNano(),
	}
}
