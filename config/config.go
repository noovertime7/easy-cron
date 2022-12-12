package config

var SysConfig *Config

type Config struct {
	Cron []CronOptions `mapstructure:"cron"`
}

type CronOptions struct {
	Name  string `mapstructure:"name"`
	Cron  string `mapstructure:"cron"`
	Shell string `mapstructure:"shell"`
}
