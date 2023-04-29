// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagebuckets

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/boundary/api"
)

// Option is a func that sets optional attributes for a call. This does not need
// to be used directly, but instead option arguments are built from the
// functions in this package. WithX options set a value to that given in the
// argument; DefaultX options indicate that the value should be set to its
// default. When an API call is made options are processed in ther order they
// appear in the function call, so for a given argument X, a succession of WithX
// or DefaultX calls will result in the last call taking effect.
type Option func(*options)

type options struct {
	postMap                 map[string]interface{}
	queryMap                map[string]string
	withAutomaticVersioning bool
	withSkipCurlOutput      bool
	withFilter              string
	withRecursive           bool
}

func getDefaultOptions() options {
	return options{
		postMap:  make(map[string]interface{}),
		queryMap: make(map[string]string),
	}
}

func getOpts(opt ...Option) (options, []api.Option) {
	opts := getDefaultOptions()
	for _, o := range opt {
		if o != nil {
			o(&opts)
		}
	}
	var apiOpts []api.Option
	if opts.withSkipCurlOutput {
		apiOpts = append(apiOpts, api.WithSkipCurlOutput(true))
	}
	if opts.withFilter != "" {
		opts.queryMap["filter"] = opts.withFilter
	}
	if opts.withRecursive {
		opts.queryMap["recursive"] = strconv.FormatBool(opts.withRecursive)
	}
	return opts, apiOpts
}

// If set, and if the version is zero during an update, the API will perform a
// fetch to get the current version of the resource and populate it during the
// update call. This is convenient but opens up the possibility for subtle
// order-of-modification issues, so use carefully.
func WithAutomaticVersioning(enable bool) Option {
	return func(o *options) {
		o.withAutomaticVersioning = enable
	}
}

// WithSkipCurlOutput tells the API to not use the current call for cURL output.
// Useful for when we need to look up versions.
func WithSkipCurlOutput(skip bool) Option {
	return func(o *options) {
		o.withSkipCurlOutput = true
	}
}

// WithFilter tells the API to filter the items returned using the provided
// filter term.  The filter should be in a format supported by
// hashicorp/go-bexpr.
func WithFilter(filter string) Option {
	return func(o *options) {
		o.withFilter = strings.TrimSpace(filter)
	}
}

// WithRecursive tells the API to use recursion for listing operations on this
// resource
func WithRecursive(recurse bool) Option {
	return func(o *options) {
		o.withRecursive = true
	}
}

func WithAttributes(inAttributes map[string]interface{}) Option {
	return func(o *options) {
		o.postMap["attributes"] = inAttributes
	}
}

func DefaultAttributes() Option {
	return func(o *options) {
		o.postMap["attributes"] = nil
	}
}

func WithBucketName(inBucketName string) Option {
	return func(o *options) {
		o.postMap["bucket_name"] = inBucketName
	}
}

func DefaultBucketName() Option {
	return func(o *options) {
		o.postMap["bucket_name"] = nil
	}
}

func WithBucketPrefix(inBucketPrefix string) Option {
	return func(o *options) {
		o.postMap["bucket_prefix"] = inBucketPrefix
	}
}

func DefaultBucketPrefix() Option {
	return func(o *options) {
		o.postMap["bucket_prefix"] = nil
	}
}

func WithDescription(inDescription string) Option {
	return func(o *options) {
		o.postMap["description"] = inDescription
	}
}

func DefaultDescription() Option {
	return func(o *options) {
		o.postMap["description"] = nil
	}
}

func WithName(inName string) Option {
	return func(o *options) {
		o.postMap["name"] = inName
	}
}

func DefaultName() Option {
	return func(o *options) {
		o.postMap["name"] = nil
	}
}

func WithPluginId(inPluginId string) Option {
	return func(o *options) {
		o.postMap["plugin_id"] = inPluginId
	}
}

func WithPluginName(inPluginName string) Option {
	return func(o *options) {
		o.queryMap["plugin_name"] = fmt.Sprintf("%v", inPluginName)
	}
}

func WithSecrets(inSecrets map[string]interface{}) Option {
	return func(o *options) {
		o.postMap["secrets"] = inSecrets
	}
}

func DefaultSecrets() Option {
	return func(o *options) {
		o.postMap["secrets"] = nil
	}
}

func WithWorkerFilter(inWorkerFilter string) Option {
	return func(o *options) {
		o.postMap["worker_filter"] = inWorkerFilter
	}
}

func DefaultWorkerFilter() Option {
	return func(o *options) {
		o.postMap["worker_filter"] = nil
	}
}
