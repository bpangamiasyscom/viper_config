package main

import (
        "fmt"
        "os"
        "strings"
        "github.com/spf13/viper"
)
func viper_load(s string) {
        viper.SetConfigName(s) // config file name without extension
        viper.SetConfigType("yaml")
        viper.AddConfigPath(".")
        viper.AddConfigPath("./config/") // config file path
        viper.AutomaticEnv()             // read value ENV variable
        viper.SetEnvPrefix(viper.GetString("ENV"))
        viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

        err := viper.ReadInConfig()
        if err != nil {
                fmt.Println("fatal error config file: default \n", err)
                os.Exit(1)
        }
}
