package retry

import "time"

func Execute[T any](
	attempts int,
	delay time.Duration,
	fn func() (T, error),
) (T, error) {

	var result T
	var err error

	for i := 0; i < attempts; i++ {

		result, err = fn()

		if err == nil {
			return result, nil
		}

		time.Sleep(delay)
	}

	return result, err
}