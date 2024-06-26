---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "xray_artifacts_scan Data Source - terraform-provider-xray"
subcategory: ""
description: |-
  Get a list of artifacts scanned by Xray for a specific repository. See JFrog Scans List - Get Artifacts API documentation https://jfrog.com/help/r/xray-rest-apis/scans-list-get-artifacts for more details.
---

# xray_artifacts_scan (Data Source)

Get a list of artifacts scanned by Xray for a specific repository. See JFrog [Scans List - Get Artifacts API documentation](https://jfrog.com/help/r/xray-rest-apis/scans-list-get-artifacts) for more details.

## Example Usage

```terraform
data "xray_artifacts_scan" "my_artifacts_scan" {
  repo = "my-docker-local"
  order_by = "repo_path"
  offset = 15
}

output "my_artifacts_scan" {
  value = data.xray_artifacts_scan.my_artifacts_scan.results
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `repo` (String) The repository key for which to get artifacts.

### Optional

- `created_end` (String) Return only records created before the specified time (in RFC 3339 format).
- `created_start` (String) Return only records created after the specified time (in RFC 3339 format).
- `direction` (String) The direction by which to order the results (either ascending or descending). Allowed value: `asc` or `desc`. Default is `asc`.
- `num_of_rows` (Number) The number of entries to return. Default is 15.
- `offset` (Number) A value returned by the API. It needs to be passed to the API to get the next page. A value of -1 means that the last page was reached.
- `order_by` (String) By which column to order the results. Allowed value: `created`, `size`, `name`, or `repo_path`.
- `repo_path` (String)

### Read-Only

- `results` (Attributes List) Result of artifacts scan. (see [below for nested schema](#nestedatt--results))

<a id="nestedatt--results"></a>
### Nested Schema for `results`

Read-Only:

- `created` (String)
- `deployed_by` (String)
- `exposures_issues` (Attributes) (see [below for nested schema](#nestedatt--results--exposures_issues))
- `malicious_packages` (Set of String)
- `name` (String)
- `package_id` (String)
- `repo_full_path` (String)
- `repo_path` (String)
- `sec_issues` (Attributes) (see [below for nested schema](#nestedatt--results--sec_issues))
- `size` (String)
- `version` (String)
- `violations` (Number)

<a id="nestedatt--results--exposures_issues"></a>
### Nested Schema for `results.exposures_issues`

Read-Only:

- `categories` (Attributes) (see [below for nested schema](#nestedatt--results--exposures_issues--categories))
- `last_scanned` (String)

<a id="nestedatt--results--exposures_issues--categories"></a>
### Nested Schema for `results.exposures_issues.categories`

Read-Only:

- `applications` (Attributes) (see [below for nested schema](#nestedatt--results--exposures_issues--categories--applications))
- `iac` (Attributes) (see [below for nested schema](#nestedatt--results--exposures_issues--categories--iac))
- `secrets` (Attributes) (see [below for nested schema](#nestedatt--results--exposures_issues--categories--secrets))
- `services` (Attributes) (see [below for nested schema](#nestedatt--results--exposures_issues--categories--services))

<a id="nestedatt--results--exposures_issues--categories--applications"></a>
### Nested Schema for `results.exposures_issues.categories.applications`

Read-Only:

- `critical` (Number)
- `high` (Number)
- `information` (Number)
- `low` (Number)
- `medium` (Number)
- `total` (Number)
- `unknown` (Number)


<a id="nestedatt--results--exposures_issues--categories--iac"></a>
### Nested Schema for `results.exposures_issues.categories.iac`

Read-Only:

- `critical` (Number)
- `high` (Number)
- `information` (Number)
- `low` (Number)
- `medium` (Number)
- `total` (Number)
- `unknown` (Number)


<a id="nestedatt--results--exposures_issues--categories--secrets"></a>
### Nested Schema for `results.exposures_issues.categories.secrets`

Read-Only:

- `critical` (Number)
- `high` (Number)
- `information` (Number)
- `low` (Number)
- `medium` (Number)
- `total` (Number)
- `unknown` (Number)


<a id="nestedatt--results--exposures_issues--categories--services"></a>
### Nested Schema for `results.exposures_issues.categories.services`

Read-Only:

- `critical` (Number)
- `high` (Number)
- `information` (Number)
- `low` (Number)
- `medium` (Number)
- `total` (Number)
- `unknown` (Number)




<a id="nestedatt--results--sec_issues"></a>
### Nested Schema for `results.sec_issues`

Read-Only:

- `critical` (Number)
- `high` (Number)
- `information` (Number)
- `low` (Number)
- `medium` (Number)
- `total` (Number)
- `unknown` (Number)
