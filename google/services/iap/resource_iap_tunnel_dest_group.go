// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/iap/TunnelDestGroup.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package iap

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceIapTunnelDestGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceIapTunnelDestGroupCreate,
		Read:   resourceIapTunnelDestGroupRead,
		Update: resourceIapTunnelDestGroupUpdate,
		Delete: resourceIapTunnelDestGroupDelete,

		Importer: &schema.ResourceImporter{
			State: resourceIapTunnelDestGroupImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"group_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Unique tunnel destination group name.`,
			},
			"cidrs": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `List of CIDRs that this group applies to.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"fqdns": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `List of FQDNs that this group applies to.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"region": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `The region of the tunnel group. Must be the same as the network resources in the group.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Full resource name.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceIapTunnelDestGroupCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	cidrsProp, err := expandIapTunnelDestGroupCidrs(d.Get("cidrs"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("cidrs"); !tpgresource.IsEmptyValue(reflect.ValueOf(cidrsProp)) && (ok || !reflect.DeepEqual(v, cidrsProp)) {
		obj["cidrs"] = cidrsProp
	}
	fqdnsProp, err := expandIapTunnelDestGroupFqdns(d.Get("fqdns"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fqdns"); !tpgresource.IsEmptyValue(reflect.ValueOf(fqdnsProp)) && (ok || !reflect.DeepEqual(v, fqdnsProp)) {
		obj["fqdns"] = fqdnsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IapBasePath}}projects/{{project}}/iap_tunnel/locations/{{region}}/destGroups?tunnelDestGroupId={{group_name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new TunnelDestGroup: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TunnelDestGroup: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating TunnelDestGroup: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/iap_tunnel/locations/{{region}}/destGroups/{{group_name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating TunnelDestGroup %q: %#v", d.Id(), res)

	return resourceIapTunnelDestGroupRead(d, meta)
}

func resourceIapTunnelDestGroupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IapBasePath}}projects/{{project}}/iap_tunnel/locations/{{region}}/destGroups/{{group_name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TunnelDestGroup: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("IapTunnelDestGroup %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading TunnelDestGroup: %s", err)
	}

	if err := d.Set("name", flattenIapTunnelDestGroupName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading TunnelDestGroup: %s", err)
	}
	if err := d.Set("cidrs", flattenIapTunnelDestGroupCidrs(res["cidrs"], d, config)); err != nil {
		return fmt.Errorf("Error reading TunnelDestGroup: %s", err)
	}
	if err := d.Set("fqdns", flattenIapTunnelDestGroupFqdns(res["fqdns"], d, config)); err != nil {
		return fmt.Errorf("Error reading TunnelDestGroup: %s", err)
	}

	return nil
}

func resourceIapTunnelDestGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TunnelDestGroup: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	cidrsProp, err := expandIapTunnelDestGroupCidrs(d.Get("cidrs"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("cidrs"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, cidrsProp)) {
		obj["cidrs"] = cidrsProp
	}
	fqdnsProp, err := expandIapTunnelDestGroupFqdns(d.Get("fqdns"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fqdns"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, fqdnsProp)) {
		obj["fqdns"] = fqdnsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IapBasePath}}projects/{{project}}/iap_tunnel/locations/{{region}}/destGroups/{{group_name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating TunnelDestGroup %q: %#v", d.Id(), obj)
	headers := make(http.Header)

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
		Headers:   headers,
	})

	if err != nil {
		return fmt.Errorf("Error updating TunnelDestGroup %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating TunnelDestGroup %q: %#v", d.Id(), res)
	}

	return resourceIapTunnelDestGroupRead(d, meta)
}

func resourceIapTunnelDestGroupDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TunnelDestGroup: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{IapBasePath}}projects/{{project}}/iap_tunnel/locations/{{region}}/destGroups/{{group_name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting TunnelDestGroup %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "TunnelDestGroup")
	}

	log.Printf("[DEBUG] Finished deleting TunnelDestGroup %q: %#v", d.Id(), res)
	return nil
}

func resourceIapTunnelDestGroupImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/iap_tunnel/locations/(?P<region>[^/]+)/destGroups/(?P<group_name>[^/]+)$",
		"^(?P<project>[^/]+)/iap_tunnel/locations/(?P<region>[^/]+)/destGroups/(?P<group_name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<group_name>[^/]+)$",
		"^(?P<region>[^/]+)/destGroups/(?P<group_name>[^/]+)$",
		"^(?P<region>[^/]+)/(?P<group_name>[^/]+)$",
		"^(?P<group_name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/iap_tunnel/locations/{{region}}/destGroups/{{group_name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenIapTunnelDestGroupName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIapTunnelDestGroupCidrs(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIapTunnelDestGroupFqdns(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandIapTunnelDestGroupCidrs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIapTunnelDestGroupFqdns(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
