package main

import (
  "github.com/cc_password_cracker/internal/cracker"
)



func main(){
  // cracker.PureBruteforce("08054846bbc9933fd0395f8be516a9f9", 4, whitelist)
  // cracker.WordlistBruteForce("2bdb742fc3d075ec6b73ea414f27819a")
  cracker.SaveRainbowTableUpToNLenghtPassword(1, 1)
}

