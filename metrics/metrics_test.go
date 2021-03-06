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

package metrics

import (
	"errors"
	"io"
	"testing"

	"go.uber.org/fx/config"
	"go.uber.org/fx/testutils"

	"github.com/stretchr/testify/assert"
	"github.com/uber-go/tally"
)

func TestRegisterReporter_OK(t *testing.T) {
	defer cleanup()

	scope, reporter, closer := getScope()
	assert.Equal(t, scope, tally.NoopScope)
	assert.Equal(t, reporter, NopCachedStatsReporter)
	assert.NoError(t, closer.Close())

	RegisterRootScope(goodScope)
	scope, reporter, closer = getScope()
	defer closer.Close()
	assert.NotNil(t, scope)
	assert.NotNil(t, reporter)
	assert.NotNil(t, closer)
}

func TestRegisterReporterPanics(t *testing.T) {
	defer cleanup()

	RegisterRootScope(goodScope)
	assert.Panics(t, func() {
		RegisterRootScope(goodScope)
	})
}

func TestRegisterReporterFrozen(t *testing.T) {
	defer cleanup()

	Freeze()
	assert.Panics(t, func() {
		RegisterRootScope(goodScope)
	})
}

func TestRegisterBadReporterPanics(t *testing.T) {
	defer cleanup()

	RegisterRootScope(badScope)
	assert.Panics(t, func() {
		getScope()
	})
}

func goodScope(
	name string,
	cfg config.Provider,
) (tally.Scope, tally.CachedStatsReporter, io.Closer, error) {
	return tally.NoopScope, NopCachedStatsReporter, testutils.NopCloser{}, nil
}

func badScope(
	name string,
	cfg config.Provider,
) (tally.Scope, tally.CachedStatsReporter, io.Closer, error) {
	return nil, nil, nil, errors.New("fake error")
}

func getScope() (tally.Scope, tally.CachedStatsReporter, io.Closer) {
	return RootScope(scopeInit())
}

func scopeInit() config.Provider {
	return config.NewStaticProvider(
		map[string]interface{}{
			config.ServiceNameKey: "somename",
			"foo": "bar",
		},
	)
}

func cleanup() {
	_scopeFunc = nil
	_frozen = false
}

func configData(data map[string]interface{}) config.Provider {
	return config.NewStaticProvider(data)
}
