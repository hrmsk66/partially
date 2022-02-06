package partially

import (
	"context"
	"strconv"
	"time"

	"github.com/fastly/go-fastly/v6/fastly"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDataCenters() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDataCentersRead,
		Schema: map[string]*schema.Schema{
			"datacenters": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"group": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"coordinates": &schema.Schema{
							Type: schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema {
								Type: schema.TypeFloat,
							},
						},
						"shield": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDataCentersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tflog.Info(ctx, "Getting /datacenters...")

	c := m.(*fastly.Client)
	var diags diag.Diagnostics

	r, err := c.AllDatacenters()
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("datacenters", r); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}