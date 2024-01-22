package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World")
}

type Cache struct {
	Map         map[string]string
	MaxCapacity int
}

// skip the constructor

func (c *Cache) Get(key string) *string {
	val, ok := c.Map[key]

	fmt.Println(c.Map)

	if !ok {
		return nil
	}

	return &val
}

func (c *Cache) Set(key string, val string) {
	_, alreadyExists := c.Map[key]

	if alreadyExists {
		c.Map[key] = val
		return
	}

	currCap := c.CurrCapacity()
	maxCap := c.MaxCapacity

	// fmt.Println("curr", currCap, "max", maxCap)

	if currCap == maxCap {
		keys := []string{}

		for existingKey := range c.Map {
			keys = append(keys, existingKey)
		}

		keyToRm := keys[0]

		// delete a key
		delete(c.Map, keyToRm)

	}

	// add in a new key
	c.Map[key] = val

	// implement kicking
}

func (c *Cache) CurrCapacity() int {
	return len(c.Map)
}
