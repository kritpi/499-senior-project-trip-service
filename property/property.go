package property

type Property struct {
	App      AppConfig      `mapstructure:"app"`
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Auth     Auth           `mapstructure:"auth"`
	Table    TableSchema    `mapstructure:"table_schema"`
	Storage  StorageConfig  `mapstructure:"storage"`
	Redis    RedisConfig    `mapstructure:"redis"`
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
	MemberTable          string `mapstructure:"member_table"`
	TripTable            string `mapstructure:"trip_table"`
	TripMembersTable     string `mapstructure:"trip_memebers_table"`
	ActivityTable        string `mapstructure:"activity_table"`
	ActivityMembersTable string `mapstructure:"activity_members_table"`
	ExpenseTable         string `mapstructure:"expense_table"`
	ExpenseMemberTable   string `mapstructure:"expense_members_table"`
}

type Auth struct {
	ClientId      string `mapstructure:"client_id"`
	JWTSecrets    string `mapstructure:"jwt_secrets"`
	ClientSecrets string `mapstructure:"client_secret"`
}

type StorageConfig struct {
	Url    string `mapstructure:"url"`
	ApiKey string `mapstructure:"api_key"`
	Bucket string `mapstructure:"bucket"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}
