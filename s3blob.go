package s3blob

import (
	"context"
	"fmt"
	"net/url"
	"sync"

	"github.com/aaronland/go-aws-auth"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"gocloud.dev/blob"
	gc_s3blob "gocloud.dev/blob/s3blob"
)

const Scheme = "s3blob"

func init() {
	blob.DefaultURLMux().RegisterBucket(Scheme, new(lazySessionOpener))
}

type URLOpener struct {
	Config *aws.Config
}

type lazySessionOpener struct {
	init   sync.Once
	opener *URLOpener
	err    error
}

func (o *lazySessionOpener) OpenBucketURL(ctx context.Context, u *url.URL) (*blob.Bucket, error) {

	o.init.Do(func() {

		cfg, err := auth.NewConfig(ctx, u.String())

		if err != nil {
			o.err = err
			return
		}

		o.opener = &URLOpener{
			Config: &cfg,
		}
	})

	if o.err != nil {
		return nil, fmt.Errorf("open bucket %v: %v", u, o.err)
	}

	return o.opener.OpenBucketURL(ctx, u)
}

func (o *URLOpener) OpenBucketURL(ctx context.Context, u *url.URL) (*blob.Bucket, error) {
	s3_client := s3.NewFromConfig(*o.Config)
	return gc_s3blob.OpenBucketV2(ctx, s3_client, u.Host, nil)
}
