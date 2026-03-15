package concurrency

type WebSiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebSites(check WebSiteChecker, websites []string) map[string]bool {

	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range websites {
		go func() {
			resultChannel <- result{url, check(url)}
		}()
	}

	for i := 0; i < len(websites); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results

}