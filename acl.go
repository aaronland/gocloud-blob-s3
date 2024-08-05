package s3blob

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// SetACLWriterOptionsWithContext return a new context.Context instance with a gocloud.dev/blob.WriterOptions
// instance used to assign 'acl' permissions for all S3 blob writes. The WriterOptions instance is assigned
// to the new context with key 'key' and is assumed to be retrieved later by code using blob.NewWriter instances.
// This method is DEPRECATED. Please use SetWriterOptionsWithContext() instead.
func SetACLWriterOptionsWithContext(ctx context.Context, key interface{}, acl string) context.Context {
	ctx, _ = SetWriterOptionsWithContext(ctx, key, "ACL", acl)
	return ctx
}

// StringACLToObjectCannedACL resolves a subset of the string values for S3 ACLs (those specific to objects) to
// their corresponding `github.com/aws/aws-sdk-go-v2/service/s3/types.ObjectCannedACL` instance.
func StringACLToObjectCannedACL(str_acl string) (types.ObjectCannedACL, error) {

	switch str_acl {
	case "private":
		return types.ObjectCannedACLPrivate, nil
	case "public-read":
		return types.ObjectCannedACLPublicRead, nil
	case "public-read-write":
		return types.ObjectCannedACLPublicReadWrite, nil
	case "authenticated-read":
		return types.ObjectCannedACLAuthenticatedRead, nil
	default:
		return "", fmt.Errorf("Invalid or unsupported ACL")
	}

}
