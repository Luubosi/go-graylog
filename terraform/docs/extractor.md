# graylog_extractor

* https://docs.graylog.org/en/latest/pages/extractors.html
* https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/graylog/resource_extractor.go

## How to import

Specify `<input id>/<extractor id>` as ID.

```console
$ terraform import graylog_extractor.test 5bb1b4b5c9e77bbbbbbbbbbb/5c4acaefc9e77bbbbbbbbbbb
```

## Argument Reference

### Common Required Argument

name | type | description
--- | --- | ---
input_id | string |
type | string |
title | string |
cursor_strategy | string |
source_field | string |
condition_type | string |
extractor_config | object{} |
converters[].type | string |
converters[].config | object{} |
converters[].config.date_format | string | "" |

### Common Optional Argument

name | type | default | description
--- | --- | --- | ---
converters | list | [] |
target_field | string | "" |
condition_value | string | "" |
order | int | 0 |
converters[].config.time_zone | string | "" |
converters[].config.locale | string | "" |

## type: grok 

```hcl
resource "graylog_extractor" "test" {
  input_id        = "0000"
  title           = "test"
  type            = "grok"
  cursor_strategy = "copy"
  source_field    = "message"
  target_field    = "none"
  condition_type  = "none"
  condition_value = ""
  order           = 0

  grok_type_extractor_config = {
    grok_pattern = "%{DATA}"
  }
}
```

### Required Argument

name | type | description
--- | --- | ---
grok_type_extractor_config | object |
grok_type_extractor_config.grok_pattern | string |

### Optional Argument

None.

## type: json

```hcl
resource "graylog_extractor" "test" {
  input_id        = "0000"
  title           = "test"
  type            = "json"
  cursor_strategy = "copy"
  source_field    = "message"
  target_field    = "none"
  condition_type  = "none"
  condition_value = ""
  order           = 0

  json_type_extractor_config = {
    list_separator             = ", "
    kv_separator               = "="
    key_prefix                 = "visit_"
    key_separator              = "_"
    replace_key_whitespace     = false
    key_whitespace_replacement = "_"
  }
}
```

## Required Argument

name | type | description
--- | --- | ---
json_type_extractor_config | object |

### Optional Argument

name | default | type | description
--- | --- | --- | ---
json_type_extractor_config.list_separator | | string |
json_type_extractor_config.kv_separator | | string |
json_type_extractor_config.key_prefix | | string |
json_type_extractor_config.key_separator | | string |
json_type_extractor_config.replace_key_whitespace | | bool |
json_type_extractor_config.key_whitespace_replacement | | string |

## type: regex

```hcl
resource "graylog_extractor" "test" {
  input_id        = "0000"
  title           = "test"
  type            = "regex"
  cursor_strategy = "copy"
  source_field    = "message"
  target_field    = "none"
  condition_type  = "none"
  condition_value = ""
  order           = 0

  regex_type_extractor_config = {
    regex_value = ".*"
  }

  converters = {
    type = "date"

    config = {
      date_format = "yyyy/MM/ddTHH:mm:ss"
      time_zone   = "Japan"
      locale      = "en"
    }
  }
}
```

## Required Argument

name | type | description
--- | --- | ---
regex_type_extractor_config | object |
regex_type_extractor_config.regex_value | string |

### Optional Argument

None.

## other types

We provide some additional attributes.

* `general_int_extractor_config`
* `general_bool_extractor_config`
* `general_float_extractor_config`
* `general_string_extractor_config`

### Required Argument

None.

### Optional Argument

name | default | type | description
--- | --- | --- | ---
general_int_extractor_config | {} | map[string]int |
general_bool_extractor_config | {} | map[string]bool |
general_float_extractor_config | {} | map[string]float64 |
general_string_extractor_config | {} | map[string]string |
