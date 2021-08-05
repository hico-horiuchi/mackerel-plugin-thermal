# mackerel-plugin-thermal

[thermal sysfs](https://www.kernel.org/doc/html/latest/driver-api/thermal/sysfs-api.html) custom metrics plugin for mackerel.io agent.

## Usage

```
mackerel-plugin-thermal
```

## Example

```
$ ./mackerel-plugin-thermal
thermal.cpu.temp        43.888000       1628155592
thermal.gpu.temp        44.444000       1628155592
```

## Install

```
mkr plugin install buty4649/mackerel-plugin-thermal
```

## Add mackerel-agent.conf

```
[plugin.metrics.thermal]
command = "/usr/bin/mackerel-plugin-thermal
```

## Author

[buty4649](https://github.com/buty4649/)
