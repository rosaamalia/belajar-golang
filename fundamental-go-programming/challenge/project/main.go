package main

import (
	"fmt"
	"sync"
)

type array interface {
	getValue() string
}

type data struct {
	value [3]string
}

func (d data) getValue() string {
    return fmt.Sprintf("%v", d.value)
}

func printGoroutine (index int, value array, wg *sync.WaitGroup) {
	fmt.Println(value.getValue(), index)
	wg.Done()
}

/**
 * Goroutine acak
 */
func goroutineAcak(bisa array, coba array, wg *sync.WaitGroup) {
	for i := 1; i <= 4; i++ {
		wg.Add(2)
		go printGoroutine(i, bisa, wg)
		go printGoroutine(i, coba, wg)
	}

	wg.Wait()
}

/**
 * Goroutine rapih
 */
func goroutineRapih(bisa array, coba array, wg *sync.WaitGroup) {
	var mtx sync.Mutex

    for i := 1; i <= 8; i++ {
        wg.Add(2)
		
		go func(i int) {
			if i%2==0 {
				mtx.Lock()
				printGoroutine(i, bisa, wg)
				mtx.Unlock()
			} else {
				mtx.Lock()
				printGoroutine(i, coba, wg)
				mtx.Unlock()
			}

		    wg.Done()
		}(i)
    }

    wg.Wait()
}

func main() {
	var wg sync.WaitGroup

	var bisa array = data {value: [3]string{"bisa1", "bisa2", "bisa3"}}
	var coba array = data {value: [3]string{"coba1", "coba2", "coba3"}}
    
	// goroutineAcak(bisa, coba, &wg)
	goroutineRapih(bisa, coba, &wg)
}