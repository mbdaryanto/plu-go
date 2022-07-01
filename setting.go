package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/fernet/fernet-go"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
)

const ENV_FILE = ".env"

type Setting struct {
	SECRET_KEY  string
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASS     string
	DB_DATABASE string
}

func CreateConfig() {
	var setting Setting
	fmt.Println("Create Config...")
	fmt.Println("Enter Database Configuration")
	fmt.Print("Host: ")
	if _, err := fmt.Scanln(&setting.DB_HOST); err != nil {
		log.Fatal(err)
	}
	fmt.Print("Port: ")
	if _, err := fmt.Scanln(&setting.DB_PORT); err != nil {
		log.Fatal(err)
	}
	fmt.Print("User: ")
	if _, err := fmt.Scanln(&setting.DB_USER); err != nil {
		log.Fatal(err)
	}
	fmt.Print("Password: ")
	if _, err := fmt.Scanln(&setting.DB_PASS); err != nil {
		log.Fatal(err)
	}
	fmt.Print("Database: ")
	if _, err := fmt.Scanln(&setting.DB_DATABASE); err != nil {
		log.Fatal(err)
	}
	var secretKey fernet.Key
	if err := (&secretKey).Generate(); err != nil {
		log.Fatal(err)
	}
	setting.SECRET_KEY = secretKey.Encode()

	// encrypt password
	if cipher, err := fernet.EncryptAndSign([]byte(setting.DB_PASS), &secretKey); err != nil {
		log.Fatal(err)
	} else {
		setting.DB_PASS = string(cipher)
	}

	// convert struct to map
	dict := make(map[string]string, 6)
	if err := mapstructure.Decode(&setting, &dict); err != nil {
		log.Fatal(err)
	}
	// fmt.Println(dict)

	if err := godotenv.Write(dict, ENV_FILE); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Settings successfully saved")
	}
}

func GetSetting() (*Setting, error) {
	if dict, err := godotenv.Read(ENV_FILE); err != nil {
		return nil, err
	} else {
		var setting Setting
		if err = mapstructure.Decode(&dict, &setting); err != nil {
			return nil, err
		} else {
			return &setting, nil
		}
	}
}

func (setting *Setting) Decode(cipher string) (string, error) {
	if secretKey, err := fernet.DecodeKeys(setting.SECRET_KEY); err != nil {
		return "", err
	} else {
		if result := fernet.VerifyAndDecrypt([]byte(cipher), 0, secretKey); result == nil {
			return "", errors.New("decrypting invalid cipher")
		} else {
			return string(result), nil
		}
	}
}

func (setting *Setting) GetDsn() string {
	config := mysql.NewConfig()
	config.User = setting.DB_USER
	if plain, err := setting.Decode(setting.DB_PASS); err != nil {
		log.Fatal(err)
	} else {
		config.Passwd = plain
	}
	config.Net = "tcp"
	config.Addr = fmt.Sprintf("%v:%v", setting.DB_HOST, setting.DB_PORT)
	config.DBName = setting.DB_DATABASE
	config.AllowNativePasswords = true
	return config.FormatDSN()
}
