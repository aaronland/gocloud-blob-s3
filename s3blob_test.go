package s3blob

import (
	"context"
	"flag"
	_ "github.com/aaronland/go-cloud-s3blob"
	"gocloud.dev/blob"
	"io"
	"testing"
)

var uri = flag.String("uri", "", "A valid s3blob:// URI.")

func TestOpenBucket(t *testing.T) {

	ctx := context.Background()

	bucket, err := blob.OpenBucket(ctx, *uri)

	if err != nil {
		t.Fatal(err)
	}

	defer bucket.Close()

	iter := bucket.List(nil)

	for {
		obj, err := iter.Next(ctx)

		if err == io.EOF {
			break
		}

		if err != nil {
			t.Fatal(err)
		}

		t.Log(obj.Key)
	}
}
