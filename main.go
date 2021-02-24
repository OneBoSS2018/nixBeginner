package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i:=1;i<=100;i++{
		wg.Add(1)
		go func(i int) {
			getPost(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func getPost(s int)  {
	url:="https://jsonplaceholder.typicode.com/posts/"+ strconv.Itoa(s)
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)

	}
	fmt.Printf("%T", err)
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	result := string(body)
	fmt.Println(result)
}


