package pg

import "fmt"

type ConnCfg struct {
	DbHost     string
	DbPort     int
	DbUser     string
	DbPassword string
	DbName   string
}

func MakeDbUrl(cfg ConnCfg) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName,
	)
}
