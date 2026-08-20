package preferences

func SetLabel(settings *Settings, key, value string) {
	settings.Labels[key] = value
}
