package cracker

import (
	"crypto/md5"
	"encoding/hex"
  "fmt"
)

func getMD5Hash(text string) string {
  hash := md5.Sum([]byte(text)) 
  return hex.EncodeToString(hash[:])
}


func CrackPassword(lenght int, hashed_value string) string{
  // The easiest way to iterate through all possible variants is... Use nested loops :D
  // I don't really care about performance yet.
  startChar := 33
  endChar := 126

  for firstChar := startChar; firstChar <= endChar; firstChar++ {
    for secondChar := startChar; secondChar <= endChar; secondChar++ {
      for thirdChar := startChar; thirdChar <= endChar; thirdChar++ {
        for fourthChar := startChar; fourthChar <= endChar; fourthChar++ {
          password := string([]byte{byte(firstChar), byte(secondChar), byte(thirdChar), byte(fourthChar)})

          hash := getMD5Hash(password)
          fmt.Printf("Plain text: %s | Hash: %s | %t \n", password, hash, hash == hashed_value)
          if hash == hashed_value{
            return password
          }
        }
      }
    }
  }
  return ""
}
