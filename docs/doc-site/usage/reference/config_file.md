---
title: "Configuration file"
description: "Use a configuration file for go-swagger"
weight: 10
---

The `swagger` command may be configured from a file for simpler reuse of predefined knobs.

You may use `swagger init config` to seed an initial default configuration.

The supported formats are json, yaml and toml.

{{% notice style="warning" title="Limitations" %}}
At this moment, only code generation commands support a config file.
{{% /notice %}}


{{< tabs groupid="config-files" >}}
{{% tab title="YAML" %}}
{{% include file="config_yaml.md" %}}
{{% /tab %}}

{{% tab title="TOML" %}}
{{% include file="config_toml.md" %}}
{{% /tab %}}

{{% tab title="JSON" %}}
{{% include file="config_json.md" %}}
{{% /tab %}}
{{< /tabs >}}
