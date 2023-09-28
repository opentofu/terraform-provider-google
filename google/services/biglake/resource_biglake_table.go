// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package biglake

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
	"github.com/hashicorp/terraform-provider-google/google/verify"
)

func ResourceBiglakeTable() *schema.Resource {
	return &schema.Resource{
		Create: resourceBiglakeTableCreate,
		Read:   resourceBiglakeTableRead,
		Update: resourceBiglakeTableUpdate,
		Delete: resourceBiglakeTableDelete,

		Importer: &schema.ResourceImporter{
			State: resourceBiglakeTableImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Output only. The name of the Table. Format:
projects/{project_id_or_number}/locations/{locationId}/catalogs/{catalogId}/databases/{databaseId}/tables/{tableId}`,
			},
			"database": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The id of the parent database.`,
			},
			"hive_options": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Options of a Hive table.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"parameters": {
							Type:     schema.TypeMap,
							Optional: true,
							Description: `Stores user supplied Hive table parameters. An object containing a
list of "key": value pairs.
Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.`,
							Elem: &schema.Schema{Type: schema.TypeString},
						},
						"storage_descriptor": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Stores physical storage information on the data.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"input_format": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `The fully qualified Java class name of the input format.`,
									},
									"location_uri": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `Cloud Storage folder URI where the table data is stored, starting with "gs://".`,
									},
									"output_format": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `The fully qualified Java class name of the output format.`,
									},
								},
							},
						},
						"table_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Hive table type. For example, MANAGED_TABLE, EXTERNAL_TABLE.`,
						},
					},
				},
			},
			"type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"HIVE", ""}),
				Description:  `The database type. Possible values: ["HIVE"]`,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Output only. The creation time of the table. A timestamp in RFC3339 UTC
"Zulu" format, with nanosecond resolution and up to nine fractional
digits. Examples: "2014-10-02T15:01:23Z" and
"2014-10-02T15:01:23.045123456Z".`,
			},
			"delete_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Output only. The deletion time of the table. Only set after the
table is deleted. A timestamp in RFC3339 UTC "Zulu" format, with
nanosecond resolution and up to nine fractional digits. Examples:
"2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The checksum of a table object computed by the server based on the value
of other fields. It may be sent on update requests to ensure the client
has an up-to-date value before proceeding. It is only checked for update
table operations.`,
			},
			"expire_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Output only. The time when this table is considered expired. Only set
after the table is deleted. A timestamp in RFC3339 UTC "Zulu" format,
with nanosecond resolution and up to nine fractional digits. Examples:
"2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
			},
			"update_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Output only. The last modification time of the table. A timestamp in
RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
fractional digits. Examples: "2014-10-02T15:01:23Z" and
"2014-10-02T15:01:23.045123456Z".`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceBiglakeTableCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	typeProp, err := expandBiglakeTableType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	hiveOptionsProp, err := expandBiglakeTableHiveOptions(d.Get("hive_options"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("hive_options"); !tpgresource.IsEmptyValue(reflect.ValueOf(hiveOptionsProp)) && (ok || !reflect.DeepEqual(v, hiveOptionsProp)) {
		obj["hiveOptions"] = hiveOptionsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BiglakeBasePath}}{{database}}/tables?tableId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Table: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating Table: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{database}}/tables/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Table %q: %#v", d.Id(), res)

	return resourceBiglakeTableRead(d, meta)
}

func resourceBiglakeTableRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BiglakeBasePath}}{{database}}/tables/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("BiglakeTable %q", d.Id()))
	}

	if err := d.Set("create_time", flattenBiglakeTableCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Table: %s", err)
	}
	if err := d.Set("update_time", flattenBiglakeTableUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Table: %s", err)
	}
	if err := d.Set("delete_time", flattenBiglakeTableDeleteTime(res["deleteTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Table: %s", err)
	}
	if err := d.Set("expire_time", flattenBiglakeTableExpireTime(res["expireTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Table: %s", err)
	}
	if err := d.Set("etag", flattenBiglakeTableEtag(res["etag"], d, config)); err != nil {
		return fmt.Errorf("Error reading Table: %s", err)
	}
	if err := d.Set("type", flattenBiglakeTableType(res["type"], d, config)); err != nil {
		return fmt.Errorf("Error reading Table: %s", err)
	}
	if err := d.Set("hive_options", flattenBiglakeTableHiveOptions(res["hiveOptions"], d, config)); err != nil {
		return fmt.Errorf("Error reading Table: %s", err)
	}

	return nil
}

func resourceBiglakeTableUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	typeProp, err := expandBiglakeTableType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	hiveOptionsProp, err := expandBiglakeTableHiveOptions(d.Get("hive_options"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("hive_options"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, hiveOptionsProp)) {
		obj["hiveOptions"] = hiveOptionsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BiglakeBasePath}}{{database}}/tables/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Table %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("type") {
		updateMask = append(updateMask, "type")
	}

	if d.HasChange("hive_options") {
		updateMask = append(updateMask, "hiveOptions")
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

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
	})

	if err != nil {
		return fmt.Errorf("Error updating Table %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Table %q: %#v", d.Id(), res)
	}

	return resourceBiglakeTableRead(d, meta)
}

func resourceBiglakeTableDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{BiglakeBasePath}}{{database}}/tables/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Table %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Table")
	}

	log.Printf("[DEBUG] Finished deleting Table %q: %#v", d.Id(), res)
	return nil
}

func resourceBiglakeTableImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^(?P<database>.+)/tables/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{database}}/tables/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenBiglakeTableCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBiglakeTableUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBiglakeTableDeleteTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBiglakeTableExpireTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBiglakeTableEtag(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBiglakeTableType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBiglakeTableHiveOptions(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["parameters"] =
		flattenBiglakeTableHiveOptionsParameters(original["parameters"], d, config)
	transformed["table_type"] =
		flattenBiglakeTableHiveOptionsTableType(original["tableType"], d, config)
	transformed["storage_descriptor"] =
		flattenBiglakeTableHiveOptionsStorageDescriptor(original["storageDescriptor"], d, config)
	return []interface{}{transformed}
}
func flattenBiglakeTableHiveOptionsParameters(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBiglakeTableHiveOptionsTableType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBiglakeTableHiveOptionsStorageDescriptor(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["location_uri"] =
		flattenBiglakeTableHiveOptionsStorageDescriptorLocationUri(original["locationUri"], d, config)
	transformed["input_format"] =
		flattenBiglakeTableHiveOptionsStorageDescriptorInputFormat(original["inputFormat"], d, config)
	transformed["output_format"] =
		flattenBiglakeTableHiveOptionsStorageDescriptorOutputFormat(original["outputFormat"], d, config)
	return []interface{}{transformed}
}
func flattenBiglakeTableHiveOptionsStorageDescriptorLocationUri(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBiglakeTableHiveOptionsStorageDescriptorInputFormat(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBiglakeTableHiveOptionsStorageDescriptorOutputFormat(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandBiglakeTableType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBiglakeTableHiveOptions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedParameters, err := expandBiglakeTableHiveOptionsParameters(original["parameters"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedParameters); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["parameters"] = transformedParameters
	}

	transformedTableType, err := expandBiglakeTableHiveOptionsTableType(original["table_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTableType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["tableType"] = transformedTableType
	}

	transformedStorageDescriptor, err := expandBiglakeTableHiveOptionsStorageDescriptor(original["storage_descriptor"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStorageDescriptor); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["storageDescriptor"] = transformedStorageDescriptor
	}

	return transformed, nil
}

func expandBiglakeTableHiveOptionsParameters(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandBiglakeTableHiveOptionsTableType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBiglakeTableHiveOptionsStorageDescriptor(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLocationUri, err := expandBiglakeTableHiveOptionsStorageDescriptorLocationUri(original["location_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocationUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["locationUri"] = transformedLocationUri
	}

	transformedInputFormat, err := expandBiglakeTableHiveOptionsStorageDescriptorInputFormat(original["input_format"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInputFormat); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["inputFormat"] = transformedInputFormat
	}

	transformedOutputFormat, err := expandBiglakeTableHiveOptionsStorageDescriptorOutputFormat(original["output_format"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOutputFormat); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["outputFormat"] = transformedOutputFormat
	}

	return transformed, nil
}

func expandBiglakeTableHiveOptionsStorageDescriptorLocationUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBiglakeTableHiveOptionsStorageDescriptorInputFormat(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBiglakeTableHiveOptionsStorageDescriptorOutputFormat(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}