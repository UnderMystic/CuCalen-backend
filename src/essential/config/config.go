package essential

import (
	"os"

	st "cucalen/essential/struct"

	"gopkg.in/yaml.v2"
)

type MyError struct{}

func (m *MyError) Error() string {
	return "boom"
}

func Config() (*st.Config, error) {
	cfg := &st.Config{}
	f, err := os.Open("config.yml")
	if err != nil {
		return cfg, (err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		return cfg, (err)
	}
	return cfg, nil
}
