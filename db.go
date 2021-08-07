package main

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

func (a *application) initRedis() {
	a.pool = &redis.Pool{
		MaxIdle:     30,
		MaxActive:   30,
		IdleTimeout: 200,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", "redis:6379")
			if err != nil {
				log.Fatalf("Unable to connect to redis: %v", err)
			}
			return conn, err
		},
	}

	app = &application{
		pool: a.pool,
	}
}

func (a *application) Get(key string) ([]byte, error) {
	conn := app.pool.Get()
	defer conn.Close()

	exists, err := app.Exists(key)

	var data []byte

	if exists {
		data, err = redis.Bytes(conn.Do("GET", key))
		if err != nil {
			log.Printf("Unable to GET key: %v", err)
		}
	} else {
		log.Printf("Key doesn't exist: %v", key)
	}

	return data, err
}

func (a *application) Exists(key string) (bool, error) {
	conn := a.pool.Get()
	defer conn.Close()

	ok, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		log.Printf("Failed to check if key exists: %v", err)
	}

	return ok, err
}
