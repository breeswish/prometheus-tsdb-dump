package main

import (
	"encoding/json"
	"os"

	"github.com/prometheus/prometheus/pkg/labels"
)

type LabelOnlyWriter struct {
}

func NewLabelOnlyWriter() *LabelOnlyWriter {
	return &LabelOnlyWriter{}
}

type line struct {
	Metric     map[string]string `json:"metric"`
	ValueCount int `json:"value_count"`
}

func (w *LabelOnlyWriter) Write(labels *labels.Labels, validTimestamps int) error {
	metric := map[string]string{}
	for _, l := range *labels {
		metric[l.Name] = l.Value
	}

	enc := json.NewEncoder(os.Stdout)
	err := enc.Encode(line{
		Metric:     metric,
		ValueCount: validTimestamps,
	})
	if err != nil {
		return err
	}
	return nil
}
