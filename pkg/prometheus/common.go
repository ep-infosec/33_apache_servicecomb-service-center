/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package prometheus

import (
	dto "github.com/prometheus/client_model/go"
)

func getValue(name string, labels map[string]string, getV func(m *dto.Metric) float64) float64 {
	f := Family(name)
	if f == nil {
		return 0
	}
	matchAll := len(labels) == 0
	var sum float64
	for _, m := range f.Metric {
		if !matchAll && !MatchLabels(m, labels) {
			continue
		}
		sum += getV(m)
	}
	return sum
}

func GaugeValue(name string, labels map[string]string) float64 {
	return getValue(name, labels, func(m *dto.Metric) float64 {
		return m.GetGauge().GetValue()
	})
}

func CounterValue(name string, labels map[string]string) float64 {
	return getValue(name, labels, func(m *dto.Metric) float64 {
		return m.GetCounter().GetValue()
	})
}

func MatchLabels(m *dto.Metric, labels map[string]string) bool {
	count := 0
	for _, label := range m.GetLabel() {
		v, ok := labels[label.GetName()]
		if ok && v != label.GetValue() {
			return false
		}
		if ok {
			count++
		}
	}
	return count == len(labels)
}

func Family(name string) *dto.MetricFamily {
	families, err := Gather()
	if err != nil {
		return nil
	}
	for _, f := range families {
		if f.GetName() == name {
			return f
		}
	}
	return nil
}
