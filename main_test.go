package main

import (
       "testing"
)

func TestAdd(t *testing.T) {
       result := Add(1, 3)
       if result != 4 {
               t.Errorf("should be 4 but %d", result)
       }
}

func TestDiv(t *testing.T) {
       result,err := Div(4, 2)
       if err!=nil{
               t.Error(err.Error())
       }
       if result != 2 {
               t.Errorf("should be 4 but %d", result)
       }
}

func TestDivZero(t *testing.T) {
       _, err := Div(4, 0)
       if err==nil{
               t.Error("cannot pass")
       }
}
