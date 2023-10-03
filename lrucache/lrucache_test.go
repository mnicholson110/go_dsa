package lrucache

import (
	"testing"
)

func TestIntLRUCache(t *testing.T) {
	c := New[string, int](3)

	if _, r := c.Get("foo"); r {
		t.Error("Expected nil")
	}

	c.Update("foo", 69)
	if _, r := c.Get("foo"); !r {
		t.Error("Expected 69")
	}

	c.Update("bar", 420)
	if _, r := c.Get("bar"); !r {
		t.Error("Expected 420")
	}

	c.Update("baz", 1337)
	if _, r := c.Get("baz"); !r {
		t.Error("Expected 1337")
	}

	c.Update("ball", 69420)
	if _, r := c.Get("ball"); !r {
		t.Error("Expected 69420")
	}

	if _, r := c.Get("foo"); r {
		t.Error("Expected nil")
	}

	if _, r := c.Get("bar"); !r {
		t.Error("Expected 420")
	}

	c.Update("foo", 69)

	if _, r := c.Get("bar"); !r {
		t.Error("Expected 420")
	}

	if _, r := c.Get("foo"); !r {
		t.Error("Expected 69")
	}

	if _, r := c.Get("baz"); r {
		t.Error("Expected nil")
	}
}
