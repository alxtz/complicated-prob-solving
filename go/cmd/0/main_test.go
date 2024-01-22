package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldGetValWorkAfterSetting(t *testing.T) {
	var c = Cache{Map: make(map[string]string), MaxCapacity: 1}

	c.Set("Oranges", "juicy")

	assert.Exactly(
		t,
		"juicy",
		*(c.Get("Oranges")),
	)
}

func Test_ShouldWorkAfterSetting2Times(t *testing.T) {
	var c = Cache{Map: make(map[string]string), MaxCapacity: 1}
	c.Set("Oranges", "juicy")
	c.Set("Oranges", "not-juicy")

	assert.Exactly(
		t,
		"not-juicy",
		*(c.Get("Oranges")),
	)
}

func Test_ShouldNilWhenGettingNonExistKey(t *testing.T) {
	var c = Cache{Map: make(map[string]string), MaxCapacity: 1}
	c.Set("Apple", "nah")

	assert.Nil(
		t,
		c.Get("Oranges"),
	)
}

// should the max capacity work
func Test_ShouldHaveAFunctionalMaxCapacity(t *testing.T) {
	var c = Cache{Map: make(map[string]string), MaxCapacity: 10}
	c.Set("Apple", "nah")

	assert.Exactly(
		t,
		10,
		c.MaxCapacity,
	)
}

// should the current map capacity itself by respected
func Test_ShouldHaveAFunctionalMapCapacity(t *testing.T) {
	var c = Cache{Map: make(map[string]string), MaxCapacity: 10}
	c.Set("Apple", "nah")
	c.Set("Apple2", "nah")

	assert.Exactly(
		t,
		2,
		c.CurrCapacity(),
	)
}

// there should not be a chance for the CurrCap to exceed MaxCap
func Test_ShouldNotExceedMaxCap(t *testing.T) {
	var c = Cache{Map: make(map[string]string), MaxCapacity: 3}
	c.Set("Apple", "nah")
	c.Set("Apple2", "nah")
	c.Set("Apple3", "nah")
	c.Set("Apple4", "nah")

	assert.Exactly(
		t,
		3,
		c.CurrCapacity(),
	)
}

// when MaxCap is exceeded, kickout the leasted used item
func Test_ShouldKickLeastUsed_WhenExceedsCap(t *testing.T) {
	var c = Cache{Map: make(map[string]string), MaxCapacity: 3}
	c.Set("Apple", "nah0")
	c.Set("Apple2", "nah1")
	c.Set("Apple3", "nah2")

	c.Get("Apple2")
	c.Get("Apple3")

	// apple should exist
	assert.Exactly(
		t,
		"nah0",
		c.Get("Apple"),
	)

	c.Set("Apple4", "nah3") // should remove apple

	// apple4 should exist
	assert.Exactly(
		t,
		"nah3",
		c.Get("Apple4"),
	)
	// apple should be missing
	assert.Nil(
		t,
		c.Get("Apple"),
	)
}
