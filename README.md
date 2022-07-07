# learnCache

## Introduction

This git is to learn about caching and implementing a middleware that will check if the current request is cached, if so, we return cached data else we fetch data from [JSONPlaceholder](https://jsonplaceholder.typicode.com/) and return the data. The response has the key `data`, if fetched else `cached` when we get the data from the cache.

## Run and test

```bash
$ go run main.go
```

The server runs at `http://localhost:3000/`. One can hit the endpoint `http://localhost:3000/<id>`, where `id` is an integer and verify the caching.
