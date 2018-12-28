package graylog

import (
	"encoding/json"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/util"
)

func resourceAlertCondition() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlertConditionCreate,
		Read:   resourceAlertConditionRead,
		Update: resourceAlertConditionUpdate,
		Delete: resourceAlertConditionDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// Required
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"stream_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"parameters": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
						"backlog": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"grace": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"query": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"repeat_notifications": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"in_grace": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func newAlertCondition(d *schema.ResourceData) (*graylog.AlertCondition, error) {
	params := &graylog.AlertConditionParameters{}
	prms := d.Get("parameters").([]interface{})[0].(map[string]interface{})
	if err := util.MSDecode(prms, params); err != nil {
		return nil, err
	}
	return &graylog.AlertCondition{
		Type:       d.Get("type").(string),
		Title:      d.Get("title").(string),
		InGrace:    d.Get("in_grace").(bool),
		ID:         d.Id(),
		Parameters: params,
	}, nil
}

func resourceAlertConditionCreate(d *schema.ResourceData, m interface{}) error {
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	cond, err := newAlertCondition(d)
	if err != nil {
		return err
	}

	if _, err := cl.CreateStreamAlertCondition(d.Get("stream_id").(string), cond); err != nil {
		return err
	}
	d.SetId(cond.ID)
	return nil
}

func resourceAlertConditionRead(d *schema.ResourceData, m interface{}) error {
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	streamID := d.Get("stream_id").(string)
	cond, _, err := cl.GetStreamAlertCondition(streamID, d.Id())
	if err != nil {
		return err
	}
	setStrToRD(d, "type", cond.Type)
	setStrToRD(d, "title", cond.Title)
	setStrToRD(d, "stream_id", streamID)
	setBoolToRD(d, "in_grace", cond.InGrace)
	if cond.Parameters != nil {
		b, err := json.Marshal(cond.Parameters)
		if err != nil {
			return err
		}
		dest := map[string]interface{}{}
		if err := json.Unmarshal(b, &dest); err != nil {
			return err
		}
		d.Set("parameters", []map[string]interface{}{dest})
	}
	return nil
}

func resourceAlertConditionUpdate(d *schema.ResourceData, m interface{}) error {
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	cond, err := newAlertCondition(d)
	if err != nil {
		return err
	}
	if _, err := cl.UpdateStreamAlertCondition(d.Get("stream_id").(string), cond); err != nil {
		return err
	}
	return nil
}

func resourceAlertConditionDelete(d *schema.ResourceData, m interface{}) error {
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	if _, err := cl.DeleteStreamAlertCondition(d.Get("stream_id").(string), d.Id()); err != nil {
		return err
	}
	return nil
}
