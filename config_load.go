package main

import (
        "fmt"
        "os"
        "strings"
        "github.com/spf13/viper"
)
func viper_load() {
// Config
        viper.SetConfigName("default") // config file name without extension
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
// // Set default value
//         viper.SetDefault("app.linetoken", "DefaultLineTokenValue")

// // Declare var
//         mode := viper.GetString("eventor.mode")
//         server_port :=  viper.GetString("services.server.port")

// // Print
//         fmt.Println("---------- Example ----------")
//         fmt.Println("mode :",mode)
//         fmt.Println("server port :",server_port)
}
