package cache

import "time"

func Get(key string) (string, error) {
	return Client.Get(Ctx, key).Result()
}

func Set(key string, value string) error {
	return Client.Set(
		Ctx,
		key,
		value,
		30*time.Minute,
	).Err()
}