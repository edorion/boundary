// Code generated by "make api"; DO NOT EDIT.
package storagebuckets

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/hashicorp/boundary/api"
	"github.com/hashicorp/boundary/api/plugins"
	"github.com/hashicorp/boundary/api/scopes"
)

type StorageBucket struct {
	Id                string                 `json:"id,omitempty"`
	ScopeId           string                 `json:"scope_id,omitempty"`
	Scope             *scopes.ScopeInfo      `json:"scope,omitempty"`
	PluginId          string                 `json:"plugin_id,omitempty"`
	Plugin            *plugins.PluginInfo    `json:"plugin,omitempty"`
	Name              string                 `json:"name,omitempty"`
	Description       string                 `json:"description,omitempty"`
	BucketName        string                 `json:"bucket_name,omitempty"`
	BucketPrefix      string                 `json:"bucket_prefix,omitempty"`
	CreatedTime       time.Time              `json:"created_time,omitempty"`
	UpdatedTime       time.Time              `json:"updated_time,omitempty"`
	Version           uint32                 `json:"version,omitempty"`
	Attributes        map[string]interface{} `json:"attributes,omitempty"`
	Secrets           map[string]interface{} `json:"secrets,omitempty"`
	SecretsHmac       string                 `json:"secrets_hmac,omitempty"`
	WorkerFilter      string                 `json:"worker_filter,omitempty"`
	AuthorizedActions []string               `json:"authorized_actions,omitempty"`

	response *api.Response
}

type StorageBucketReadResult struct {
	Item     *StorageBucket
	response *api.Response
}

func (n StorageBucketReadResult) GetItem() *StorageBucket {
	return n.Item
}

func (n StorageBucketReadResult) GetResponse() *api.Response {
	return n.response
}

type StorageBucketCreateResult = StorageBucketReadResult
type StorageBucketUpdateResult = StorageBucketReadResult

type StorageBucketDeleteResult struct {
	response *api.Response
}

// GetItem will always be nil for StorageBucketDeleteResult
func (n StorageBucketDeleteResult) GetItem() interface{} {
	return nil
}

func (n StorageBucketDeleteResult) GetResponse() *api.Response {
	return n.response
}

type StorageBucketListResult struct {
	Items    []*StorageBucket
	response *api.Response
}

func (n StorageBucketListResult) GetItems() []*StorageBucket {
	return n.Items
}

func (n StorageBucketListResult) GetResponse() *api.Response {
	return n.response
}

// Client is a client for this collection
type Client struct {
	client *api.Client
}

// Creates a new client for this collection. The submitted API client is cloned;
// modifications to it after generating this client will not have effect. If you
// need to make changes to the underlying API client, use ApiClient() to access
// it.
func NewClient(c *api.Client) *Client {
	return &Client{client: c.Clone()}
}

// ApiClient returns the underlying API client
func (c *Client) ApiClient() *api.Client {
	return c.client
}

func (c *Client) Create(ctx context.Context, scopeId string, opt ...Option) (*StorageBucketCreateResult, error) {
	if scopeId == "" {
		return nil, fmt.Errorf("empty scopeId value passed into Create request")
	}

	opts, apiOpts := getOpts(opt...)

	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts.postMap["scope_id"] = scopeId

	req, err := c.client.NewRequest(ctx, "POST", "storage-buckets", opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating Create request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during Create call: %w", err)
	}

	target := new(StorageBucketCreateResult)
	target.Item = new(StorageBucket)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding Create response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.response = resp
	return target, nil
}

func (c *Client) Read(ctx context.Context, id string, opt ...Option) (*StorageBucketReadResult, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id value passed into Read request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("storage-buckets/%s", url.PathEscape(id)), nil, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating Read request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during Read call: %w", err)
	}

	target := new(StorageBucketReadResult)
	target.Item = new(StorageBucket)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding Read response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.response = resp
	return target, nil
}

func (c *Client) Update(ctx context.Context, id string, version uint32, opt ...Option) (*StorageBucketUpdateResult, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id value passed into Update request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, errors.New("zero version number passed into Update request and automatic versioning not specified")
		}
		existingTarget, existingErr := c.Read(ctx, id, append([]Option{WithSkipCurlOutput(true)}, opt...)...)
		if existingErr != nil {
			if api.AsServerError(existingErr) != nil {
				return nil, fmt.Errorf("error from controller when performing initial check-and-set read: %w", existingErr)
			}
			return nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingTarget == nil {
			return nil, errors.New("nil resource response found when performing initial check-and-set read")
		}
		if existingTarget.Item == nil {
			return nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Item.Version
	}

	opts.postMap["version"] = version

	req, err := c.client.NewRequest(ctx, "PATCH", fmt.Sprintf("storage-buckets/%s", url.PathEscape(id)), opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating Update request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during Update call: %w", err)
	}

	target := new(StorageBucketUpdateResult)
	target.Item = new(StorageBucket)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding Update response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.response = resp
	return target, nil
}

func (c *Client) Delete(ctx context.Context, id string, opt ...Option) (*StorageBucketDeleteResult, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id value passed into Delete request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	req, err := c.client.NewRequest(ctx, "DELETE", fmt.Sprintf("storage-buckets/%s", url.PathEscape(id)), nil, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating Delete request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during Delete call: %w", err)
	}

	apiErr, err := resp.Decode(nil)
	if err != nil {
		return nil, fmt.Errorf("error decoding Delete response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}

	target := &StorageBucketDeleteResult{
		response: resp,
	}
	return target, nil
}

func (c *Client) List(ctx context.Context, scopeId string, opt ...Option) (*StorageBucketListResult, error) {
	if scopeId == "" {
		return nil, fmt.Errorf("empty scopeId value passed into List request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)
	opts.queryMap["scope_id"] = scopeId

	req, err := c.client.NewRequest(ctx, "GET", "storage-buckets", nil, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating List request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during List call: %w", err)
	}

	target := new(StorageBucketListResult)
	apiErr, err := resp.Decode(target)
	if err != nil {
		return nil, fmt.Errorf("error decoding List response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.response = resp
	return target, nil
}
