// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package auth

import (
	"context"

	"go.uber.org/fx/config"

	"github.com/uber-go/tally"
)

var (
	// NopClient is used for testing and no-op integration
	NopClient = nopClient(nil, tally.NoopScope)

	_ Client = &nop{}
)

type nop struct {
}

func nopClient(config config.Provider, scope tally.Scope) Client {
	return &nop{}
}

func (*nop) Name() string {
	return "nop"
}

func (*nop) Authenticate(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

func (*nop) Authorize(ctx context.Context) error {
	return nil
}

func (*nop) SetAttribute(ctx context.Context, key, value string) context.Context {
	return ctx
}
