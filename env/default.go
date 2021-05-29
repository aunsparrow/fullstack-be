package env

import (
	"os"
)

func SetDefaultEnv() {
	os.Setenv("ENV", "Default")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_NAME", "paiduay")
	os.Setenv("DB_USERNAME", "postgres")
	os.Setenv("DB_PASSWORD", "Password#01")
}
