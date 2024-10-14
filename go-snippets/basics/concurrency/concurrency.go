package concurrency

import "time"

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		go func(u string) {
			//from the inside the goroutine has access to "urls" variable, but
			//that would make all of them read/write the variable at the same time
			results[u] = wc(u)
		}(url)
	}

	time.Sleep(2 * time.Second)

	return results
}
