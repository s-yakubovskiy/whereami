package metrics

var globalLabels = map[string]string{
	// "env": "production",
}

// SetGlobalLabel sets a key-value pair for all metrics
func SetGlobalLabel(key, value string) {
	globalLabels[key] = value
}

// GetGlobalLabels returns a copy of the global labels map
func GetGlobalLabels() map[string]string {
	labels := make(map[string]string, len(globalLabels))
	for k, v := range globalLabels {
		labels[k] = v
	}
	return labels
}
