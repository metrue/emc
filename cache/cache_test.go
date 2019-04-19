package cache

import "testing"

func TestCache(t *testing.T) {
	t.Run("Init/GetCacheClient", func(t *testing.T) {
		if cache != nil {
			t.Fatal("cache should be nil at first")
		}

		Init()

		if cache == nil {
			t.Fatal("cache should be initialized now")
		}

		c1 := GetCacheClient()
		c2 := GetCacheClient()
		if c1 != c2 {
			t.Fatal("cache client should be singleton")
		}
	})

	t.Run("Set/Get", func(t *testing.T) {
		key := "mock_k"
		if _, ok := Get(key); ok {
			t.Fatal("should not ok since no set yet")
		}
		val := "mock_v"
		Set(key, val)
		v, ok := Get(key)
		if !ok {
			t.Fatal("should not ok since no set yet")
		}
		if v != val {
			t.Fatalf("should get %v but got %v", val, v)
		}
	})
}
