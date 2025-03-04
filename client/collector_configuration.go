package client

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/suzuki-shunsuke/go-set"

	"github.com/suzuki-shunsuke/go-graylog"
)

// CreateCollectorConfiguration creates a collector configuration.
func (client *Client) CreateCollectorConfiguration(
	ctx context.Context, cfg *graylog.CollectorConfiguration,
) (*ErrorInfo, error) {
	// POST /plugins/org.graylog.plugins.collector/configurations Create new collector configuration
	if cfg == nil {
		return nil, fmt.Errorf("collector configuration is nil")
	}
	if cfg.Inputs == nil {
		cfg.Inputs = []graylog.CollectorConfigurationInput{}
	}
	if cfg.Outputs == nil {
		cfg.Outputs = []graylog.CollectorConfigurationOutput{}
	}
	if cfg.Snippets == nil {
		cfg.Snippets = []graylog.CollectorConfigurationSnippet{}
	}
	if cfg.Tags == nil {
		cfg.Tags = set.StrSet{}
	}
	return client.callPost(ctx, client.Endpoints().CollectorConfigurations(), cfg, cfg)
}

// GetCollectorConfigurations returns all collector configurations.
func (client *Client) GetCollectorConfigurations(ctx context.Context) ([]graylog.CollectorConfiguration, int, *ErrorInfo, error) {
	cfgs := &graylog.CollectorConfigurationsBody{}
	ei, err := client.callGet(
		ctx, client.Endpoints().CollectorConfigurations(), nil, cfgs)
	return cfgs.Configurations, cfgs.Total, ei, err
}

// GetCollectorConfiguration returns a given user.
func (client *Client) GetCollectorConfiguration(
	ctx context.Context, id string,
) (*graylog.CollectorConfiguration, *ErrorInfo, error) {
	// GET /api/plugins/org.graylog.plugins.collector/configurations/:id
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}
	u, err := client.Endpoints().CollectorConfiguration(id)
	if err != nil {
		return nil, nil, err
	}
	cfg := &graylog.CollectorConfiguration{}
	ei, err := client.callGet(ctx, u.String(), nil, cfg)
	return cfg, ei, err
}

// RenameCollectorConfiguration renames a collector configuration.
func (client *Client) RenameCollectorConfiguration(
	ctx context.Context, id, name string,
) (*graylog.CollectorConfiguration, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is nil")
	}
	if name == "" {
		return nil, nil, fmt.Errorf("name is nil")
	}
	input := graylog.CollectorConfiguration{
		Name:     name,
		Inputs:   []graylog.CollectorConfigurationInput{},
		Outputs:  []graylog.CollectorConfigurationOutput{},
		Snippets: []graylog.CollectorConfigurationSnippet{},
	}
	u, err := client.Endpoints().CollectorConfigurationName(id)
	if err != nil {
		return nil, nil, err
	}
	cfg := graylog.CollectorConfiguration{Name: name}
	ei, err := client.callPut(ctx, u.String(), &input, &cfg)
	return &cfg, ei, err
}

// DeleteCollectorConfiguration deletes a collector configuration.
func (client *Client) DeleteCollectorConfiguration(
	ctx context.Context, id string,
) (*ErrorInfo, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	u, err := client.Endpoints().CollectorConfiguration(id)
	if err != nil {
		return nil, err
	}
	return client.callDelete(ctx, u.String(), nil, nil)
}
