package cracker

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

type PassHashPair struct {
  plain, hash string
}

func (r PassHashPair) String() string{
  return "plain: "+r.plain+"\t hash:"+r.hash+"\n"
}


var whitelist []rune = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$%^&*()_+{}|:<>?~")


func GeneratePassHashPairsOfLengthN(lenght int) []PassHashPair{
  totalCombinations := powInt(len(whitelist), lenght)
  characters := make([]rune, lenght)


  pairs := make([]PassHashPair, totalCombinations)  

  for i := 0; i < totalCombinations; i++{
    for j := 0; j < lenght ; j++{
      characters[j] = whitelist[rune((i/powInt(len(whitelist),j))%int(len(whitelist)))]
    }

    pairs[i] = PassHashPair{string(characters), getMD5Hash(string(characters))}

  }

  return pairs
}


func SaveRainbowTableUpToNLenghtPassword(maxLength, minLength int){
  if minLength == 0{
    minLength=1 
  }
  
  var pairs []PassHashPair

  for n:=minLength;n<=maxLength;n++{
    pairs = append(pairs, GeneratePassHashPairsOfLengthN(n)...)
  }
  
  currentDir, err := os.Getwd()
  if err != nil {
    fmt.Println("Error getting current directory:", err)
    return
  }
  rainbowTablesDir := filepath.Join(currentDir, "rainbow_tables")

  if err := os.MkdirAll(rainbowTablesDir, 0755); err != nil {
    fmt.Println("Error creating rainbow_tables directory:", err)
    return
  }

  filePath := filepath.Join(rainbowTablesDir, "numerations.csv") 
  file, err := os.Create(filePath)
  if err != nil {
		panic(err)
	}
	defer file.Close()

  csvWriter := csv.NewWriter(file)
  
  headers := []string{"plain", "hash"}
  csvWriter.Write(headers)

  for _, pair := range pairs{
    record := []string{pair.plain, pair.hash}
    csvWriter.Write(record)
  }
  
  csvWriter.Flush()
  if err := csvWriter.Error(); err != nil{
    panic(err)
  }
 
} 
