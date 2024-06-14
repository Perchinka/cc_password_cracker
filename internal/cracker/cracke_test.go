package cracker

import (
  "testing"
)

func TestFourDigitPass(t *testing.T){
  tests := []struct {
    name  string
    input string
    expected string
  }{
    {"Numeric", "81dc9bdb52d04dc20036dbd8313ed055", "1234"},
    {"Alphabetic", "2378e46cc86f8ea4e157da9f7354d670", "help"},
    {"Chars", "c5df4802c14ccf1cd5b83d5d3ca6b238", "#$!@"},
    {"Mixed", "0aefcbbcb29cc747edb12dbeaa8d9c2b", "h3!p"},
    
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      result := CrackPassword(4,tt.input)
      if result != tt.expected {
        t.Errorf("got %v, want %v", result, tt.expected)
      }
    })
  }
}
