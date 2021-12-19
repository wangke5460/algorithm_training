package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

func main() {
	eg, ctx := errgroup.WithContext(context.Background())
	eg.Go(func() error {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
		})
		http.ListenAndServe(":8080", nil)
		return errors.New("test1")
	})

	c := make(chan os.Signal, 1)

	eg.Go(func() error {
		signal.Notify(c)
		return errors.New("test2")
	})
	if err := eg.Wait(); err != nil {
		fmt.Println(err)
	}

}
