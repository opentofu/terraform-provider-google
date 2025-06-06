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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/securitycenter/MuteConfig.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package securitycenter

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
	"github.com/hashicorp/terraform-provider-google/google/verify"
)

func ResourceSecurityCenterMuteConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityCenterMuteConfigCreate,
		Read:   resourceSecurityCenterMuteConfigRead,
		Update: resourceSecurityCenterMuteConfigUpdate,
		Delete: resourceSecurityCenterMuteConfigDelete,

		Importer: &schema.ResourceImporter{
			State: resourceSecurityCenterMuteConfigImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"filter": {
				Type:     schema.TypeString,
				Required: true,
				Description: `An expression that defines the filter to apply across create/update
events of findings. While creating a filter string, be mindful of
the scope in which the mute configuration is being created. E.g.,
If a filter contains project = X but is created under the
project = Y scope, it might not match any findings.`,
			},
			"mute_config_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Unique identifier provided by the client within the parent scope.`,
			},
			"parent": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Resource name of the new mute configs's parent. Its format is
"organizations/[organization_id]", "folders/[folder_id]", or
"projects/[project_id]".`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A description of the mute config.`,
			},
			"expiry_time": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `Optional. The expiry of the mute config. Only applicable for dynamic configs.
If the expiry is set, when the config expires, it is removed from all findings.

A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
			},
			"type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"MUTE_CONFIG_TYPE_UNSPECIFIED", "STATIC", "DYNAMIC", ""}),
				Description:  `The type of the mute config, which determines what type of mute state the config affects. Default value: "DYNAMIC" Possible values: ["MUTE_CONFIG_TYPE_UNSPECIFIED", "STATIC", "DYNAMIC"]`,
				Default:      "DYNAMIC",
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The time at which the mute config was created. This field is set by
the server and will be ignored if provided on config creation.`,
			},
			"most_recent_editor": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Email address of the user who last edited the mute config. This
field is set by the server and will be ignored if provided on
config creation or update.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Name of the mute config. Its format is
organizations/{organization}/muteConfigs/{configId},
folders/{folder}/muteConfigs/{configId},
or projects/{project}/muteConfigs/{configId}`,
			},
			"update_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Output only. The most recent time at which the mute config was
updated. This field is set by the server and will be ignored if
provided on config creation or update.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceSecurityCenterMuteConfigCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandSecurityCenterMuteConfigDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	filterProp, err := expandSecurityCenterMuteConfigFilter(d.Get("filter"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("filter"); !tpgresource.IsEmptyValue(reflect.ValueOf(filterProp)) && (ok || !reflect.DeepEqual(v, filterProp)) {
		obj["filter"] = filterProp
	}
	typeProp, err := expandSecurityCenterMuteConfigType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	expiryTimeProp, err := expandSecurityCenterMuteConfigExpiryTime(d.Get("expiry_time"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("expiry_time"); !tpgresource.IsEmptyValue(reflect.ValueOf(expiryTimeProp)) && (ok || !reflect.DeepEqual(v, expiryTimeProp)) {
		obj["expiryTime"] = expiryTimeProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecurityCenterBasePath}}{{parent}}/muteConfigs?muteConfigId={{mute_config_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new MuteConfig: %#v", obj)
	billingProject := ""

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
		return fmt.Errorf("Error creating MuteConfig: %s", err)
	}
	// Set computed resource properties from create API response so that they're available on the subsequent Read
	// call.
	err = resourceSecurityCenterMuteConfigPostCreateSetComputedFields(d, meta, res)
	if err != nil {
		return fmt.Errorf("setting computed ID format fields: %w", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating MuteConfig %q: %#v", d.Id(), res)

	return resourceSecurityCenterMuteConfigRead(d, meta)
}

func resourceSecurityCenterMuteConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecurityCenterBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("SecurityCenterMuteConfig %q", d.Id()))
	}

	if err := d.Set("name", flattenSecurityCenterMuteConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading MuteConfig: %s", err)
	}
	if err := d.Set("description", flattenSecurityCenterMuteConfigDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading MuteConfig: %s", err)
	}
	if err := d.Set("filter", flattenSecurityCenterMuteConfigFilter(res["filter"], d, config)); err != nil {
		return fmt.Errorf("Error reading MuteConfig: %s", err)
	}
	if err := d.Set("create_time", flattenSecurityCenterMuteConfigCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading MuteConfig: %s", err)
	}
	if err := d.Set("update_time", flattenSecurityCenterMuteConfigUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading MuteConfig: %s", err)
	}
	if err := d.Set("most_recent_editor", flattenSecurityCenterMuteConfigMostRecentEditor(res["mostRecentEditor"], d, config)); err != nil {
		return fmt.Errorf("Error reading MuteConfig: %s", err)
	}
	if err := d.Set("type", flattenSecurityCenterMuteConfigType(res["type"], d, config)); err != nil {
		return fmt.Errorf("Error reading MuteConfig: %s", err)
	}
	if err := d.Set("expiry_time", flattenSecurityCenterMuteConfigExpiryTime(res["expiryTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading MuteConfig: %s", err)
	}

	return nil
}

func resourceSecurityCenterMuteConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	descriptionProp, err := expandSecurityCenterMuteConfigDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	filterProp, err := expandSecurityCenterMuteConfigFilter(d.Get("filter"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("filter"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, filterProp)) {
		obj["filter"] = filterProp
	}
	typeProp, err := expandSecurityCenterMuteConfigType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	expiryTimeProp, err := expandSecurityCenterMuteConfigExpiryTime(d.Get("expiry_time"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("expiry_time"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, expiryTimeProp)) {
		obj["expiryTime"] = expiryTimeProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecurityCenterBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating MuteConfig %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("filter") {
		updateMask = append(updateMask, "filter")
	}

	if d.HasChange("type") {
		updateMask = append(updateMask, "type")
	}

	if d.HasChange("expiry_time") {
		updateMask = append(updateMask, "expiryTime")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
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
			return fmt.Errorf("Error updating MuteConfig %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating MuteConfig %q: %#v", d.Id(), res)
		}

	}

	return resourceSecurityCenterMuteConfigRead(d, meta)
}

func resourceSecurityCenterMuteConfigDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{SecurityCenterBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting MuteConfig %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "MuteConfig")
	}

	log.Printf("[DEBUG] Finished deleting MuteConfig %q: %#v", d.Id(), res)
	return nil
}

func resourceSecurityCenterMuteConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	if err := tpgresource.ParseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	// current import_formats can't import fields with forward slashes in their value
	name := d.Get("name").(string)

	matched, err := regexp.MatchString("(organizations|folders|projects)/.+/muteConfigs/.+", name)
	if err != nil {
		return nil, fmt.Errorf("error validating import name: %s", err)
	}

	if !matched {
		return nil, fmt.Errorf("error validating import name: %s does not fit naming for muteConfigs. Expected %s",
			name, "organizations/{organization}/muteConfigs/{configId}, folders/{folder}/muteConfigs/{configId} or projects/{project}/muteConfigs/{configId}")
	}

	if err := d.Set("name", name); err != nil {
		return nil, fmt.Errorf("Error setting name: %s", err)
	}

	// mute_config_id and parent are not returned by the API and therefore need to be set manually
	stringParts := strings.Split(d.Get("name").(string), "/")
	if err := d.Set("mute_config_id", stringParts[3]); err != nil {
		return nil, fmt.Errorf("Error setting mute_config_id: %s", err)
	}

	if err := d.Set("parent", fmt.Sprintf("%s/%s", stringParts[0], stringParts[1])); err != nil {
		return nil, fmt.Errorf("Error setting mute_config_id: %s", err)
	}

	return []*schema.ResourceData{d}, nil
}

func flattenSecurityCenterMuteConfigName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterMuteConfigDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterMuteConfigFilter(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterMuteConfigCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterMuteConfigUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterMuteConfigMostRecentEditor(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterMuteConfigType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterMuteConfigExpiryTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandSecurityCenterMuteConfigDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterMuteConfigFilter(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterMuteConfigType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterMuteConfigExpiryTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceSecurityCenterMuteConfigPostCreateSetComputedFields(d *schema.ResourceData, meta interface{}, res map[string]interface{}) error {
	config := meta.(*transport_tpg.Config)
	if err := d.Set("name", flattenSecurityCenterMuteConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}
	return nil
}
