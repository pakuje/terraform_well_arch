---
subcategory: "Data Factory"
layout: "azurerm"
page_title: "Azure Resource Manager: azurerm_data_factory_pipeline"
description: |-
  Manages a Pipeline inside a Azure Data Factory.
---

# azurerm_data_factory_pipeline

Manages a Pipeline inside a Azure Data Factory.

## Example Usage

```hcl
resource "azurerm_resource_group" "example" {
  name     = "example-resources"
  location = "West Europe"
}

resource "azurerm_data_factory" "example" {
  name                = "example"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
}

resource "azurerm_data_factory_pipeline" "example" {
  name                = "example"
  resource_group_name = azurerm_resource_group.example.name
  data_factory_id     = azurerm_data_factory.example.id
}
```

## Example Usage with Activities

```hcl
resource "azurerm_data_factory_pipeline" "test" {
  name                = "acctest%d"
  resource_group_name = azurerm_resource_group.test.name
  data_factory_id     = azurerm_data_factory.test.id
  variables = {
    "bob" = "item1"
  }
  activities_json = <<JSON
[
	{
		"name": "Append variable1",
		"type": "AppendVariable",
		"dependsOn": [],
		"userProperties": [],
		"typeProperties": {
			"variableName": "bob",
			"value": "something"
		}
	}
]
  JSON
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Specifies the name of the Data Factory Pipeline. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.

* `resource_group_name` - (Required) The name of the resource group in which to create the Data Factory Pipeline. Changing this forces a new resource

* `data_factory_id` - (Optional) The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.

* `data_factory_name` - (Optional) The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.

-> **Note:** This property has been deprecated in favour of the `data_factory_id` property and will be removed in version 3.0 of the provider.

-> **Note:** At least one of `data_factory_id` or `data_factory_name` must be set.

* `description` - (Optional) The description for the Data Factory Pipeline.

* `annotations` - (Optional) List of tags that can be used for describing the Data Factory Pipeline.

* `concurrency` - (Optional) The max number of concurrent runs for the Data Factory Pipeline. Must be between `1` and `50`.

* `folder` - (Optional) The folder that this Pipeline is in. If not specified, the Pipeline will appear at the root level.

* `moniter_metrics_after_duration` - (Optional) The TimeSpan value after which an Azure Monitoring Metric is fired.

* `parameters` - (Optional) A map of parameters to associate with the Data Factory Pipeline.

* `variables` - (Optional) A map of variables to associate with the Data Factory Pipeline.

* `activities_json` - (Optional) A JSON object that contains the activities that will be associated with the Data Factory Pipeline.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Data Factory Pipeline.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 30 minutes) Used when creating the Data Factory Pipeline.
* `update` - (Defaults to 30 minutes) Used when updating the Data Factory Pipeline.
* `read` - (Defaults to 5 minutes) Used when retrieving the Data Factory Pipeline.
* `delete` - (Defaults to 30 minutes) Used when deleting the Data Factory Pipeline.

## Import

Data Factory Pipeline's can be imported using the `resource id`, e.g.

```shell
terraform import azurerm_data_factory_pipeline.example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/pipelines/example
```
