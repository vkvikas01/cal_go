package main

import (
	"fmt"
	"sync"
)

func worker(id int,jobs <-chan int,wg *sync.WaitGroup){
	defer wg.Done()

	for j:= range jobs{
		fmt.Println("Worker",id,"started job",j)
	}
}

func main() {

	jobs := make(chan int,10)

	var wg sync.WaitGroup

	for i := 0;i<=3;i++{
		wg.Add(1)
		go worker(i,jobs,&wg)
	}

	for j:=0;j<10;j++{
		jobs <- j
	}
	close(jobs)
	wg.Wait()
	fmt.Println("All workers finished their jobs")
}