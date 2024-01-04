package models

type Config struct {
	App   *App   `yaml:"app"`
	Sever *Sever `yaml:"sever"`
}

type App struct {
	FirstSub string `yaml:"FirstSub"`
	OtherSub string `yaml:"OtherSub"`
	TokenSub string `yaml:"TokenSub"`
	UsedSub  string `yaml:"UsedSub"`
}

type Sever struct {
	RdsHost       string `yaml:"RdsHost"`
	RdsPort       int    `yaml:"RdsPort"`
	RdsUser       string `yaml:"RdsUser"`
	RdsDB         int    `yaml:"RdsDB"`
	RdsPass       string `yaml:"RdsPass"`
	Authorization string `yaml:"Authorization"`
	APIUrl        string `yaml:"APIUrl"`
	WxToken       string `yaml:"WxToken"`
	Limit         int    `yaml:"Limit"`
	Expire        int    `yaml:"Expire"`
}
