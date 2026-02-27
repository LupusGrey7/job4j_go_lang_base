package base_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_Lru_Cache(t *testing.T) {
	t.Parallel()

	t.Run("PUT LruCache - true", func(t *testing.T) {

		lru := base.NewLruCache(10)
		key := "some cache key"
		value := "1"
		lru.Put(key, value)
		fmt.Printf("LRU size is %d\n", lru.GetSize())

		assert.Equal(t, value, lru.Get(key))
	})

	t.Run("PUT LruCache more then have - true", func(t *testing.T) {

		lru := base.NewLruCache(3)
		key := "some cache key1"
		value := "1"
		lru.Put(key, value)
		key2 := "some cache key2"
		value2 := "2"
		lru.Put(key2, value2)
		key3 := "some cache key3"
		value3 := "3"
		lru.Put(key3, value3)

		fmt.Printf("LRU size is %d\n", lru.GetSize())

		assert.Equal(t, 3, lru.GetSize())
	})

	t.Run("PUT LruCache more then have - false", func(t *testing.T) {
		lru := base.NewLruCache(3)
		key := "some cache key1"
		value := "1"
		lru.Put(key, value)
		key2 := "some cache key2"
		value2 := "2"
		lru.Put(key2, value2)
		key3 := "some cache key3"
		value3 := "3"
		lru.Put(key3, value3)
		key4 := "some cache key4"
		value4 := "value4"
		lru.Put(key4, value4)

		fmt.Printf("LRU size is %d\n", lru.GetSize())

		assert.Equal(t, value4, lru.Get(key4))
	})
}
