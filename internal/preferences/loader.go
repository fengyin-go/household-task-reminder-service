package preferences

type Settings struct {
	Labels map[string]string
}

func LoadDefault() *Settings {
	return &Settings{Labels: make(map[string]string)}
}
