---
title: Osm Pull
menu:
  product_osm_0.6.1:
    identifier: osm-pull
    name: Osm Pull
    parent: reference
product_name: osm
menu_name: product_osm_0.6.1
section_menu_id: reference
---
## osm pull

Pull item from container

### Synopsis

Pull item from container

```
osm pull <src> <dest> [flags]
```

### Examples

```
osm pull -c mybucket f1.txt /tmp/f1.txt
```

### Options

```
  -c, --container string   Name of container
      --context string     Name of osmconfig context to use
  -h, --help               help for pull
```

### Options inherited from parent commands

```
      --alsologtostderr                  log to standard error as well as files
      --analytics                        Send analytical events to Google Analytics (default true)
      --log_backtrace_at traceLocation   when logging hits line file:N, emit a stack trace (default :0)
      --log_dir string                   If non-empty, write log files in this directory
      --logtostderr                      log to standard error instead of files
      --osmconfig string                 Path to osm config (default "/home/tamal/.osm/config")
      --stderrthreshold severity         logs at or above this threshold go to stderr (default 2)
  -v, --v Level                          log level for V logs
      --vmodule moduleSpec               comma-separated list of pattern=N settings for file-filtered logging
```

### SEE ALSO

* [osm](/docs/reference/osm.md)	 - Object Store Manipulator by AppsCode

