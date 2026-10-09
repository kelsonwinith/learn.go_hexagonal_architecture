package config

// ============================================================================
// Types
// ============================================================================

type Config struct {
	App        App
	PostgreSQL PostgreSQL
	JWT        JWT
}

type App struct {
	Port string `envconfig:"APP_PORT" default:"8080"`
}

type JWT struct {
	Secret      string `envconfig:"JWT_SECRET" required:"true"`
	ExpiryHours int    `envconfig:"JWT_EXPIRY_HOURS" default:"24"`
}

type PostgreSQL struct {
	DBHost     string `envconfig:"POSTGRESQL_HOST" default:"localhost"`
	DBPort     string `envconfig:"POSTGRESQL_PORT" default:"5432"`
	DBUser     string `envconfig:"POSTGRESQL_USER" required:"true"`
	DBPassword string `envconfig:"POSTGRESQL_PASSWORD" required:"true"`
	DBName     string `envconfig:"POSTGRESQL_NAME" required:"true"`
	DBSSLMode  string `envconfig:"POSTGRESQL_SSLMODE" default:"disable"`
}
