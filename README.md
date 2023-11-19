# LRU Cache Implementation in Go

This repository contains an implementation of a Least Recently Used (LRU) Cache in Go. The LRU cache is a popular caching algorithm that evicts the least recently accessed items first. This implementation is thread-safe and can be used in various applications where caching of data is required to improve performance.

## Features

- **LRU Eviction Policy**: Implements the Least Recently Used (LRU) eviction policy.
- **Thread-Safety**: Ensures that the cache can be safely used in a multi-threaded environment.
- **Flexible Key-Value Storage**: Allows storing any type of data as a value with an integer key.
- **Capacity Limit**: The cache has a fixed capacity set during initialization, beyond which older items are evicted.

## Getting Started

### Prerequisites

- Go (version 1.15 or later)

### Installing

Clone the repository to your local machine:

```bash
git clone https://github.com/rafaelmgr12/lru-cache.git
cd lru-cache
````

### Usage
To use the LRU cache, import the package and create a new cache instance:

```go
package main

import (
    lru "github.com/rafaelmgr12/lru-cache/pkg"
)

func main() {
    // Create a new cache with a capacity of 2 items
    cache := lru.NewLRUCache(2)

    // Set values in the cache
    cache.Set(1, "item1")
    cache.Set(2, "item2")

    // Retrieve item from the cache
    value := cache.Get(1) // returns "item1"
    // ...
}
```

### Running the Tests
To run the tests, use the following command:

```bash
go test ./pkg/...

```
## Contributing

Contributions are welcome! For major changes, please open an issue first to discuss what you would like to change.

## License 
This project is licensed under the MIT License - see the [LICENSE](LICENSE) fore more details
