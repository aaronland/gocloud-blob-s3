package main

import (
	"context"
	"log"

	s3blob "github.com/aaronland/gocloud-blob-s3"
	"github.com/aaronland/gocloud-blob/app/copy"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/fileblob"
	_ "gocloud.dev/blob/memblob"
)

var str_acl string

func main() {

	ctx := context.Background()

	fs := copy.DefaultFlagSet()

	fs.StringVar(&str_acl, "acl", "", "...")

	opts, err := copy.RunOptionsFromFlagSet(fs)

	if err != nil {
		log.Fatal(err)
	}

	if str_acl != "" {

		acl, err := s3blob.StringACLToObjectCannedACL(str_acl)

		if err != nil {
			log.Fatal(err)
		}

		before := func(asFunc func(interface{}) bool) error {

			req := &s3.PutObjectInput{}
			ok := asFunc(&req)

			if ok {
				req.ACL = acl
			}

			return nil
		}

		opts.WriterOptions = &blob.WriterOptions{
			BeforeWrite: before,
		}
	}

	err = copy.RunWithOptions(ctx, opts)

	if err != nil {
		log.Fatal(err)
	}
}
