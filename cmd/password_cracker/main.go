package main

import (
  "github.com/cc_password_cracker/internal/cracker"
)



func main(){
  whitelist := []rune("ABCDEFGHIJKLMOPQRSTUVWXYZ")
  cracker.CrackPassword("7a95bf926a0333f57705aeac07a362a2", 4, whitelist)
}

