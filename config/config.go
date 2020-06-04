package config

const DefaultDocPath = "dbdoc"

type DSN struct {
	URL     string            `yaml:"url"`
	Headers map[string]string `yaml:"headers,omitempty"`
}
