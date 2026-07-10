module github.com/tpretz/terraform-provider-zabbix

go 1.11

require (
	github.com/hashicorp/terraform v0.12.23
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/makuartur/go-zabbix-api v0.17.1-0.20260710162921-2a862330fff3
)

//replace github.com/makuartur/go-zabbix-api => ../go-zabbix-api
