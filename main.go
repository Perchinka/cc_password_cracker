package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
)

func getMD5Hash(text string) string {
  hash := md5.Sum([]byte(text)) 
  return hex.EncodeToString(hash[:])
}

func main(){
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Password to hash > ")
  text, _ := reader.ReadString('\n')
  hash := getMD5Hash(text)
  fmt.Print("Hashed password: ", hash, "\n")
}

