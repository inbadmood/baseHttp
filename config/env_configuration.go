package config

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/spf13/viper"
	"log"
	"reflect"
	"strings"
)

var EnvConfig envConfiguration

type envConfiguration struct {
	Env               string            `env:"env" validate:"required"`
	Port              string            `env:"port"`
	ConnectionStrings ConnectionStrings `env:"connectionstrings" validate:"required"`
	API               ApiSettings       `env:"api" validate:"required"`
}

type App struct {
	Port string `env:"port"`
}

type ConnectionStrings struct {
	Mysql MysqlConfig `env:"Mysql" validate:"required"`
	Redis RedisConfig `env:"redis"`
}

type MysqlConfig struct {
	Host     string `env:"host" validate:"required"`
	Database string `env:"database" validate:"required"`
	User     string `env:"user" validate:"required"`
	Password string `env:"password" validate:"required"`
	Port     string `env:"port"`
}

type RedisConfig struct {
	Host     string `env:"host" validate:"required"`
	Database *int   `env:"database" validate:"required"`
}

type ApiSettings struct {
	Url   string `env:"url" validate:"required"`
	Token string `env:"token" validate:"required"`
}

func InitialEnvConfiguration() (err error) {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("app")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	bindEnvs(EnvConfig)
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		switch err.(type) {
		default:
			panic(fmt.Errorf("Fatal error loading config file: %s \n", err))
		case viper.ConfigFileNotFoundError:
			log.Print("No config file found. Using defaults and environment variables")
		}
	}
	err = viper.Unmarshal(&EnvConfig)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	validate := validator.New()
	if err = validate.Struct(&EnvConfig); err != nil {
		return err
	}
	fmt.Printf("%#v \n", EnvConfig)

	return nil
}

func bindEnvs(iFace interface{}, parts ...string) {
	ifv := reflect.ValueOf(iFace)
	ift := reflect.TypeOf(iFace)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("env")
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			bindEnvs(v.Interface(), append(parts, tv)...)
		default:
			_ = viper.BindEnv(strings.Join(append(parts, tv), "."))
		}
	}
}
