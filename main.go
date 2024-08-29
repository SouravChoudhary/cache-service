package main

import (
	"cache-service/cache/inmemorycache"
	"fmt"
	"time"
)

func main() {

	c := inmemorycache.NewCache()

	// Set some values in the cache
	c.Set("key1", "value1", 3*time.Second)
	c.Set("key2", "value2", 6*time.Second)

	// Retrieve the values
	val1, err := c.Get("key1")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("key1:", val1)
	}

	// Wait for key1 to expire
	time.Sleep(4 * time.Second)

	// Try to retrieve the expired key
	val1, err = c.Get("key1")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("key1:", val1)
	}

	// Delete key2
	err = c.Delete("key2")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("key2 deleted")
	}

	// Try to retrieve the deleted key
	val2, err := c.Get("key2")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("key2:", val2)
	}
}
