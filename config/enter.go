package config

// 直接用成员体不加变量名的话会导致开发时遇到麻烦
// 查找成员域时会直接调用许多变量

type Config struct {
	Mysql    Mysql    `yaml:"mysql"`
	Logger   Logger   `yaml:"logger"`
	System   System   `yaml:"system"`
	SiteInfo SiteInfo `yaml:"site_info"`
	QQ       QQ       `yaml:"qq"`
	QiNiu    QiNiu    `yaml:"qi_niu"`
	Email    Email    `yaml:"email"`
	Jwt      Jwt      `yaml:"jwt"`
	Uploads  Uploads  `yaml:"uploads"`
	Redis    Redis    `yaml:"redis"`
	ES       ES       `yaml:"es"`
}
