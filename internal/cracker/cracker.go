package cracker

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
  "math"
)

func getMD5Hash(text string) string {
  hash := md5.Sum([]byte(text)) 
  return hex.EncodeToString(hash[:])
}

func powInt(x, y int) int {
    return int(math.Pow(float64(x), float64(y)))
}


func CrackPassword(hashed_value string, password_lenght int, startChar, endChar rune) string{
  if startChar == 0 {
    startChar = 33
  }
  if endChar == 0{
    endChar = 126
  }
  if password_lenght == 0{
    password_lenght = 4
  }
  
  charRange := int(endChar-startChar+1)
  totalCombinations := powInt(charRange, password_lenght)
  characters := make([]rune, password_lenght)

  // Counting from one cause math easier this way :)
  for i := 1; i < totalCombinations+1; i++{
    for j := 0; j < password_lenght; j++{
      characters[j] = rune(int(startChar)+int(i/powInt(26,j))%int(endChar))
    }
    hashed_pass := getMD5Hash(string(characters))
    fmt.Printf("Plain: %s | Hash: %s | same? %t \n", string(characters), hashed_pass, hashed_pass==hashed_value)

    if hashed_pass == hashed_value { return string(characters)}
  }

  return ""
}
