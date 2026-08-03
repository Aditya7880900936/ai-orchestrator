package cache

import "time"

var redisGet = func(key string) (string, error) {
	return Client.Get(Ctx, key).Result()
}

var redisSet = func(key, value string, ttl time.Duration) error {
	return Client.Set(Ctx, key, value, ttl).Err()
}

func Get(key string) (string, error) {
	return redisGet(key)
}

func Set(key string, value string) error {
	return redisSet(key, value, 30*time.Minute)
}