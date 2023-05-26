package config

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

func (cfg *Config) DbUrl() string {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	return "postgres://" + cfg.Database.User + ":" + cfg.Database.Password + "@" + cfg.Database.Host + ":" + cfg.Database.Port + "/" + cfg.Database.Name
}
