docs/cmd/mavbake_upgrade.md## mavbake upgrade

Upgrades BB.

### Synopsis

Upgrades BB instance.

```
mavbake upgrade [flags]
```

### Options

```
  -h, --help              help for upgrade
      --node              Upgrade node.
      --pay               Upgrade pay.
      --peak              Upgrade peak.
  -a, --setup-ami         Install latest ami during the BB upgrade.
      --signer            Upgrade signer.
  -s, --upgrade-storage   Upgrade storage during the upgrade.
```

### Options inherited from parent commands

```
  -l, --log-level string       Sets output log format (json/text/auto) (default "info")
  -o, --output-format string   Sets output log format (json/text/auto) (default "auto")
  -p, --path string            Path to mavpay instance (default "/mavpay")
```

### SEE ALSO

* [mavbake](/mavbake/reference/cmd/mavbake)	 - mavbake CLI

###### Auto generated by spf13/cobra on 26-Sep-2024
