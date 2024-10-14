package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string //anonymous values!
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			//from the inside the goroutine has access to "urls" variable, but
			//that would make all of them read/write the variable at the same time
			resultChannel <- result{u, wc(u)} //the arrow is a "send statement"
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel // "foo := <- baz" is a "receive expression"
		results[r.string] = r.bool
	}

	return results
}
