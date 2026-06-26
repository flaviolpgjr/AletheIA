package captcha

import "context"

type Validator interface {
	Validate(ctx context.Context, token string) error
}