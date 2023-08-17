package microfiber

import "time"

// Retry 重试函数
func Retry(fn func() error, n uint, d time.Duration) (err error) {
	var count uint
	for {
		err = fn()
		if err != nil {
			time.Sleep(d)
			count++
		} else {
			err = nil
			break
		}
		if count > n {
			break
		}
	}
	return
}
