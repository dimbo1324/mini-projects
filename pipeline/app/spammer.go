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
	batch := make([]User, 0, GetMessagesMaxUsersBatch)
	var wg sync.WaitGroup
	worker := func(usersBatch []User) {
		defer wg.Done()
		msgs, err := GetMessages(usersBatch...) //? что это
		if err != nil {
			_ = fmt.Sprintf("GetMessages error: %v", err)
			return
		}
		for _, m := range msgs {
			out <- m
		}
	}
	for raw := range in {
		u, ok := raw.(User)
		if !ok {
			continue
		}
		batch = append(batch, u)
		if len(batch) >= GetMessagesMaxUsersBatch {
			batchCopy := make([]User, len(batch))
			copy(batchCopy, batch)
			wg.Add(1)
			go worker(batchCopy)
			batch = batch[:0]
		}
	}
	if len(batch) > 0 {
		batchCopy := make([]User, len(batch))
		copy(batchCopy, batch)
		wg.Add(1)
		go worker(batchCopy)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
}

func CheckSpam(in, out chan any) {
}

func CombineResults(in, out chan any) {
}
