package pkg

type Config struct {
	DatabaseHost     string `env:"DATABASE_HOST"`
	DatabaseUser     string `env:"DATABASE_USER"`
	DatabasePassword string `env:"DATABASE_PASSWORD"`
	DatabaseName     string `env:"DATABASE_NAME"`
}
