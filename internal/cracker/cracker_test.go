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
    {"CODE", "08054846bbc9933fd0395f8be516a9f9", "CODE"},
    {"PASS", "7a95bf926a0333f57705aeac07a362a2", "PASS"},
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      result := PureBruteforce(tt.input,4,[]rune("ABCDEFGHIJKLMOPQRSTUVWXYZ"))
      if result != tt.expected {
        t.Errorf("got %v, want %v", result, tt.expected)
      }
    })
  }
}

