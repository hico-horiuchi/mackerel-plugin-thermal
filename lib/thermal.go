package mpthermal

import (
	"flag"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"

	mp "github.com/mackerelio/go-mackerel-plugin"
)

type ThermalPlugin struct {
	prefix string
}

func (t ThermalPlugin) MetricKeyPrefix() string {
	if t.prefix == "" {
		t.prefix = "thermal"
	}
	return t.prefix
}

func (t ThermalPlugin) GraphDefinition() map[string]mp.Graphs {
	labelPrefix := strings.Title(t.prefix)
	return map[string]mp.Graphs{
		"#": {
			Label: labelPrefix,
			Unit:  mp.UnitFloat,
			Metrics: []mp.Metrics{
				{Name: "temp", Label: "Temperature"},
			},
		},
	}
}

func (t ThermalPlugin) FetchMetrics() (map[string]float64, error) {
	result := make(map[string]float64)

	files, err := filepath.Glob("/sys/class/thermal/thermal_zone*")
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		thermal_type, err := ioutil.ReadFile(filepath.Join(file, "type"))
		if err != nil {
			return nil, err
		}

		temperature, err := ioutil.ReadFile(filepath.Join(file, "temp"))
		if err != nil {
			return nil, err
		}
		temp, err := strconv.ParseFloat(strings.TrimRight(string(temperature), "\n"), 64)
		if err != nil {
			return nil, err
		}

		key := strings.TrimRight(string(thermal_type), "\n") + ".temp"
		result[key] = temp / 1000.0
	}

	return result, nil
}

func Do() {
	optPrefix := flag.String("metric-key-prefix", "thermal", "Metric key prefix")
	flag.Parse()

	mp.NewMackerelPlugin(&ThermalPlugin{
		prefix: *optPrefix,
	}).Run()
}
