package preferences

func SetLabel(settings *Settings, key, value string) {
	if settings.Labels == nil {
		settings.Labels = make(map[string]string)
	}
	settings.Labels[key] = value
}
