package utils

import (
	"math/rand"
	"time"

	"github.com/spf13/viper"
)


func init() {
	rand.NewSource(time.Now().UnixNano())
}


func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}


func RandomOwner() string {
	return RandomString(12)
}

func RandomMoney() int64 {
	return int64(RandomInt(0, 1000))
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}



type Config struct {
    DBDriver      string `mapstructure:"DB_DRIVER"`
    DBSource      string `mapstructure:"DB_SOURCE"`
    ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}