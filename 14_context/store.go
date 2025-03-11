package main

//If you use ctx.Value in my (non-existent) company, you’re fired
//Not my quote :-)
//Michael Strba says passing context everywhere is a smell
//And should be fixed in Go 2

import (
	"context"
	"fmt"
	"net/http"
)

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return // TODO: log error however you like
		}

		fmt.Fprint(w, data)
	}
}

type Store interface {
	Fetch(ctx context.Context) (string, error)
}
