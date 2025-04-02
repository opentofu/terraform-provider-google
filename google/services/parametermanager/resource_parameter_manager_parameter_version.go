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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/parametermanager/ParameterVersion.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package parametermanager

import (
	"encoding/base64"
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

	"google.golang.org/api/googleapi"
)

func ResourceParameterManagerParameterVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceParameterManagerParameterVersionCreate,
		Read:   resourceParameterManagerParameterVersionRead,
		Update: resourceParameterManagerParameterVersionUpdate,
		Delete: resourceParameterManagerParameterVersionDelete,

		Importer: &schema.ResourceImporter{
			State: resourceParameterManagerParameterVersionImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"parameter": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `Parameter Manager Parameter resource.`,
			},
			"parameter_version_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Version ID of the Parameter Version Resource. This must be unique within the Parameter.`,
			},
			"parameter_data": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The Parameter data.`,
				Sensitive:   true,
			},

			"disabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `The current state of Parameter Version. This field is only applicable for updating Parameter Version.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time at which the Parameter Version was created.`,
			},
			"kms_key_version": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The resource name of the Cloud KMS CryptoKeyVersion used to decrypt parameter version payload. Format
'projects/{{project}}/locations/global/keyRings/{{key_ring}}/cryptoKeys/{{crypto_key}}/cryptoKeyVersions/{{crypto_key_version}}'`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The resource name of the Parameter Version. Format:
'projects/{{project}}/locations/global/parameters/{{parameter_id}}/versions/{{parameter_version_id}}'`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time at which the Parameter Version was updated.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceParameterManagerParameterVersionCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	disabledProp, err := expandParameterManagerParameterVersionDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(disabledProp)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}
	payloadProp, err := expandParameterManagerParameterVersionPayload(nil, d, config)
	if err != nil {
		return err
	} else if !tpgresource.IsEmptyValue(reflect.ValueOf(payloadProp)) {
		obj["payload"] = payloadProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ParameterManagerBasePath}}{{parameter}}/versions?parameter_version_id={{parameter_version_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ParameterVersion: %#v", obj)
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
		return fmt.Errorf("Error creating ParameterVersion: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{parameter}}/versions/{{parameter_version_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating ParameterVersion %q: %#v", d.Id(), res)

	return resourceParameterManagerParameterVersionRead(d, meta)
}

func resourceParameterManagerParameterVersionRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ParameterManagerBasePath}}{{parameter}}/versions/{{parameter_version_id}}")
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ParameterManagerParameterVersion %q", d.Id()))
	}

	if err := d.Set("name", flattenParameterManagerParameterVersionName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading ParameterVersion: %s", err)
	}
	if err := d.Set("create_time", flattenParameterManagerParameterVersionCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ParameterVersion: %s", err)
	}
	if err := d.Set("update_time", flattenParameterManagerParameterVersionUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ParameterVersion: %s", err)
	}
	if err := d.Set("disabled", flattenParameterManagerParameterVersionDisabled(res["disabled"], d, config)); err != nil {
		return fmt.Errorf("Error reading ParameterVersion: %s", err)
	}
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if flattenedProp := flattenParameterManagerParameterVersionPayload(res["payload"], d, config); flattenedProp != nil {
		if gerr, ok := flattenedProp.(*googleapi.Error); ok {
			return fmt.Errorf("Error reading ParameterVersion: %s", gerr)
		}
		casted := flattenedProp.([]interface{})[0]
		if casted != nil {
			for k, v := range casted.(map[string]interface{}) {
				if err := d.Set(k, v); err != nil {
					return fmt.Errorf("Error setting %s: %s", k, err)
				}
			}
		}
	}
	if err := d.Set("kms_key_version", flattenParameterManagerParameterVersionKmsKeyVersion(res["kmsKeyVersion"], d, config)); err != nil {
		return fmt.Errorf("Error reading ParameterVersion: %s", err)
	}

	return nil
}

func resourceParameterManagerParameterVersionUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	disabledProp, err := expandParameterManagerParameterVersionDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ParameterManagerBasePath}}{{parameter}}/versions/{{parameter_version_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ParameterVersion %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("disabled") {
		updateMask = append(updateMask, "disabled")
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
			return fmt.Errorf("Error updating ParameterVersion %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating ParameterVersion %q: %#v", d.Id(), res)
		}

	}

	return resourceParameterManagerParameterVersionRead(d, meta)
}

func resourceParameterManagerParameterVersionDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{ParameterManagerBasePath}}{{parameter}}/versions/{{parameter_version_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting ParameterVersion %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "ParameterVersion")
	}

	log.Printf("[DEBUG] Finished deleting ParameterVersion %q: %#v", d.Id(), res)
	return nil
}

func resourceParameterManagerParameterVersionImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := tpgresource.ParseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	name := d.Get("name").(string)
	parameterRegex := regexp.MustCompile("(projects/.+/locations/global/parameters/.+)/versions/.+$")
	versionRegex := regexp.MustCompile("projects/(.+)/locations/global/parameters/(.+)/versions/(.+)$")

	parts := parameterRegex.FindStringSubmatch(name)
	if len(parts) != 2 {
		return nil, fmt.Errorf("Version name does not fit the format `projects/{{project}}/locations/global/parameters/{{parameter_id}}/versions/{{parameter_version_id}}`")
	}
	if err := d.Set("parameter", parts[1]); err != nil {
		return nil, fmt.Errorf("Error setting parameter: %s", err)
	}

	parts = versionRegex.FindStringSubmatch(name)

	if err := d.Set("parameter_version_id", parts[3]); err != nil {
		return nil, fmt.Errorf("Error setting parameter_version_id: %s", err)
	}

	return []*schema.ResourceData{d}, nil
}

func flattenParameterManagerParameterVersionName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenParameterManagerParameterVersionCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenParameterManagerParameterVersionUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenParameterManagerParameterVersionDisabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenParameterManagerParameterVersionPayload(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	data, err := base64.StdEncoding.DecodeString(original["data"].(string))
	if err != nil {
		return err
	}
	transformed["parameter_data"] = string(data)
	return []interface{}{transformed}
}

func flattenParameterManagerParameterVersionKmsKeyVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandParameterManagerParameterVersionDisabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandParameterManagerParameterVersionPayload(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	transformed := make(map[string]interface{})
	transformedParameterData, err := expandParameterManagerParameterVersionPayloadParameterData(d.Get("parameter_data"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedParameterData); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["data"] = transformedParameterData
	}

	return transformed, nil
}

func expandParameterManagerParameterVersionPayloadParameterData(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == nil {
		return nil, nil
	}

	return base64.StdEncoding.EncodeToString([]byte(v.(string))), nil
}
