package config

type MSSQLDBConfig struct {
	SQLHost        string `split_words:"true" default:"sql-rafqkaoj4kdlc.database.windows.net:1433"`
	SQLDatabase    string `split_words:"true" default:"rizz-and-comp-db"`
	SQLUsername    string `split_words:"true" default:"sa"`
	SQLPassword    string `split_words:"true" default:"SuperSecret1337"`
	LocalSQLHost   string
	LocalSQLPort   string
	LocalSQLDbName string
	MSSQLConnectionConfig
}

type MSSQLConnectionConfig struct {
	SQLDialTimeout               uint  `split_words:"true" default:"2"`
	SQLConnectionTimeout         uint  `split_words:"true" default:"2"`
	SQLMaxOpenConnections        int   `split_words:"true" default:"50"`
	SQLMaxIdleConnections        int   `split_words:"true" default:"15"`
	SQLMaxConnectionLifetimeSecs int64 `split_words:"true" default:"300"`
}

type Config struct {
	IsLocal  func() bool
	DBConfig MSSQLDBConfig
}
