package main

import (
        "testing"
        "github.com/spf13/viper"
)

func Test1(t *testing.T){
        viper_load("default")
        mode := viper.GetString("eventor.mode")
        if mode != "dev" {
                t.Errorf("mode : %s", mode)
        }
        viper_load("default2")
        mode = viper.GetString("eventor.mode")
        if mode != "dev2" {
                t.Errorf("mode : %s", mode)
        }
}
