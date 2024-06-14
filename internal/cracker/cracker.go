package cracker

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math"
	"os"
)

func getMD5Hash(text string) string {
  hash := md5.Sum([]byte(text)) 
  return hex.EncodeToString(hash[:])
}

func powInt(x, y int) int {
    return int(math.Pow(float64(x), float64(y)))
}


func PureBruteforce(hashed_value string, password_lenght int, whitelist []rune) string{
  if password_lenght == 0{
    password_lenght = 4
  }
  
  if whitelist == nil{
    whitelist = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$%^&*()_+{}|:<>?~")
  }

  totalCombinations := powInt(len(whitelist), password_lenght)
  characters := make([]rune, password_lenght)

  // Counting from one cause math easier this way :)
  for i := 0; i < totalCombinations+1; i++{
    for j := password_lenght-1; j >= 0 ; j--{
      characters[j] = whitelist[rune((i/powInt(len(whitelist),j))%int(len(whitelist)))]
    }

    plain_text := string(characters)
    hashed_pass := getMD5Hash(plain_text)
 
    fmt.Printf("Plain: %s | Hash: %s | same? %t \n", plain_text, hashed_pass, hashed_pass==hashed_value)
    
    if hashed_pass == hashed_value { 
      return string(characters)
    }
  }

  return ""
}


func WordlistBruteForce(hashed_value string) string {
  // For now I will hardcode wordlistpath, but in the future it will be read from the arguments
  wordlistFile, err := os.Open("rockyou.txt")
  if err != nil {
      panic(err)
  }
  defer wordlistFile.Close() 

  scanner := bufio.NewScanner(wordlistFile)
  for scanner.Scan(){
    line := scanner.Text()
    hashed_pass := getMD5Hash(line)
    
    fmt.Printf("Plain: %s | Hash: %s | same? %t \n", line, hashed_pass, hashed_pass==hashed_value)


    if hashed_pass == hashed_value{
      return line
    }
  }

  return ""
}
