package property

type Property struct {
	App      AppConfig      `mapstructure:"app"`
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Auth     Auth           `mapstructure:"auth"`
	Table    TableSchema    `mapstructure:"table_schema"`
}

type AppConfig struct {
	Name     string `mapstructure:"name"`
	Env      string `mapstructure:"env"`
	Timezone string `mapstructure:"timezone"`
}

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type DatabaseConfig struct {
	Url string `mapstructure:"url"`
}

type TableSchema struct {
	MemberTable string `mapstructure:"member_table"`
}

type Auth struct {
	ClientId   string `mapstructure:"client_id"`
	JWTSecrets string `mapstructure:"jwt_secrets"`
}
