package essential

type Config struct {
	Conf struct {
		Host   string `yaml:"host"`
		Port   string `yaml:"port"`
		Adpass string `yaml:"adpass"`
	} `yaml:"conf"`
	Db struct {
		Dbhost string `yaml:"dbhost"`
		User   string `yaml:"user"`
		Pass   string `yaml:"pass"`
		Dbname string `yaml:"dbname"`
		Port   string `yaml:"port"`
	} `yaml:"db"`
	Env struct {
		Typ string `yaml:"type"`
	} `yaml:"env"`
}

type GET struct{}
type POST struct{}
