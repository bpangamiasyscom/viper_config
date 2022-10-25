package main

import (
        "testing"
        "github.com/spf13/viper"
)

func Test1(t *testing.T){
        viper_load()
        mode := viper.GetString("eventor.mode")
        if mode != "dev" {
                t.Errorf("mode : %s", mode)
        }
}
