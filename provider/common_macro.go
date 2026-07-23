package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/makuartur/go-zabbix-api"
)

// macroListSchema models macros as a set because Zabbix does not guarantee
// their order in API responses. A macro name is unique within a host/template,
// so it is also the stable identity of a set element.
var macroListSchema = &schema.Schema{
	Type:     schema.TypeSet,
	Optional: true,
	Set: func(v interface{}) int {
		return schema.HashString(v.(map[string]interface{})["name"])
	},
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
				Description:  "Macro Name (key)",
			},
			"value": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
				Description:  "Macro Value",
			},
		},
	},
}

// macroGenerate build macro structs from terraform inputs
func macroGenerate(d *schema.ResourceData) (macros zabbix.Macros) {
	list := d.Get("macro").(*schema.Set).List()
	macros = make(zabbix.Macros, len(list))
	for i, raw := range list {
		item := raw.(map[string]interface{})
		macros[i] = zabbix.Macro{
			MacroName: item["name"].(string),
			Value:     item["value"].(string),
			MacroID:   item["id"].(string),
		}
	}

	return
}

// flattenMacros convert response to terraform input
func flattenMacros(list zabbix.Macros) []interface{} {
	val := make([]interface{}, len(list))
	for i := 0; i < len(list); i++ {
		val[i] = map[string]interface{}{
			"name":  list[i].MacroName,
			"value": list[i].Value,
			"id":    list[i].MacroID,
		}
	}
	return val
}
