package main

import (
	"fmt"
	"sync"
)

func RunPipeline(cmds ...cmd) {
}

func SelectUsers(in, out chan any) {
	usersCh := make(chan User)
	var wg sync.WaitGroup
	go func() {
		for raw := range in {
			email := fmt.Sprintf("%v", raw)
			wg.Add(1)
			go func(e string) {
				defer wg.Done()
				user := GetUser(e)
				usersCh <- user
			}(email)
		}
		wg.Wait()
		close(usersCh)
	}()
	seen := make(map[uint64]struct{})
	for u := range usersCh {
		if _, ok := seen[u.ID]; !ok {
			seen[u.ID] = struct{}{}
			out <- u
		}
	}
	close(out)
}

func SelectMessages(in, out chan any) {
}

func CheckSpam(in, out chan any) {
}

func CombineResults(in, out chan any) {
}
