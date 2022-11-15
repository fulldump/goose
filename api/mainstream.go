package api

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"goose/inceptiondb"
)

func MainStream(ctx context.Context, w http.ResponseWriter) error {

	max := 10
	reader, err := GetInceptionClient(ctx).Find("tweets", inceptiondb.FindQuery{
		Index:   "by timestamp",
		Limit:   max,
		Reverse: true,
	})
	if err != nil {
		return fmt.Errorf("error reading from persistence layer")
	}

	io.Copy(w, reader)

	return nil
}
