package client

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/suzuki-shunsuke/go-graylog"
)

// GetInputs returns all inputs.
func (client *Client) GetInputs(ctx context.Context) (
	[]graylog.Input, int, *ErrorInfo, error,
) {
	inputs := &graylog.InputsBody{}
	ei, err := client.callGet(ctx, client.Endpoints().Inputs(), nil, inputs)
	return inputs.Inputs, inputs.Total, ei, err
}

// GetInput returns a given input.
func (client *Client) GetInput(
	ctx context.Context, id string,
) (*graylog.Input, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}
	u, err := client.Endpoints().Input(id)
	if err != nil {
		return nil, nil, err
	}
	input := &graylog.Input{}
	ei, err := client.callGet(ctx, u.String(), nil, input)
	return input, ei, err
}

// CreateInput creates an input.
func (client *Client) CreateInput(
	ctx context.Context, input *graylog.Input,
) (ei *ErrorInfo, err error) {
	if input == nil {
		return nil, fmt.Errorf("input is nil")
	}
	if input.ID != "" {
		return nil, fmt.Errorf("input id should be empty")
	}
	// change attributes to configuration
	// https://github.com/Graylog2/graylog2-server/issues/3480
	d := map[string]interface{}{
		"title":         input.Title,
		"type":          input.Type(),
		"configuration": input.Attrs,
		"global":        input.Global,
	}
	if input.Node != "" {
		d["node"] = input.Node
	}

	return client.callPost(ctx, client.Endpoints().Inputs(), &d, input)
}

// UpdateInput updates an given input.
func (client *Client) UpdateInput(
	ctx context.Context, prms *graylog.InputUpdateParams,
) (*graylog.Input, *ErrorInfo, error) {
	if prms == nil {
		return nil, nil, fmt.Errorf("input is nil")
	}
	if prms.ID == "" {
		return nil, nil, errors.New("id is empty")
	}
	u, err := client.Endpoints().Input(prms.ID)
	if err != nil {
		return nil, nil, err
	}
	// change attributes to configuration
	// https://github.com/Graylog2/graylog2-server/issues/3480
	d := map[string]interface{}{
		"title":         prms.Title,
		"type":          prms.Type,
		"configuration": prms.Attrs,
		"global":        prms.Global,
	}
	if prms.Node != "" {
		d["node"] = prms.Node
	}
	input := &graylog.Input{}
	ei, err := client.callPut(ctx, u.String(), &d, input)
	return input, ei, err
}

// DeleteInput deletes an given input.
func (client *Client) DeleteInput(
	ctx context.Context, id string,
) (*ErrorInfo, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	u, err := client.Endpoints().Input(id)
	if err != nil {
		return nil, err
	}
	return client.callDelete(ctx, u.String(), nil, nil)
}
