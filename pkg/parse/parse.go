package xparse

import (
	"github.com/zeromicro/go-zero/core/validation"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"sync/atomic"
)

var (
	validator atomic.Value
)

func Parse(r *http.Request, v any) error {
	if err := httpx.ParsePath(r, v); err != nil {
		return err
	}

	if err := httpx.ParseForm(r, v); err != nil {
		return err
	}

	if err := httpx.ParseHeaders(r, v); err != nil {
		return err
	}

	if err := httpx.ParseJsonBody(r, v); err != nil {
		return err
	}

	if valid, ok := v.(validation.Validator); ok {
		return valid.Validate()
	} else if val := validator.Load(); val != nil {
		return val.(httpx.Validator).Validate(r, v)
	}

	return nil
}
