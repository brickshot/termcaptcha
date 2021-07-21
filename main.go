package main

import (
  "fmt"
  "math/rand"
  "strconv"
  "time"
)

const e = "\033["

func main() {
  colorRed := e + "31m"
  cr := "\r"
  clearline := cr + e + "2K"

  var words = []string{
    "HELLO", "THERE", "ABLAZE", "ABOARD", "BOOHOO", "CHROMA", "ENTREE", "GUFFAW",
  }

  A := 65 // Uppercase A

  // Create and seed the generator.
  // Typically a non-fixed seed should be used, such as time.Now().UnixNano().
  // Using a fixed seed will produce the same output on every run.
  // r := rand.New(rand.NewSource(99))
  r := rand.New(rand.NewSource(time.Now().UnixNano()))

  linelength := 30
  fmt.Printf(clearline + colorRed + "| " + cursorToX(linelength+3) + "|")
  pos := make([]int, linelength)
  for i := range pos {
    pos[i] = i
  }

  var wordx int
  word := words[r.Intn(len(words))]
  for i := 0; i < 1000; i++ {
    if i%100 == 0 {
      wordx = r.Intn(20 - len(word))
    }
    order := r.Perm(len(pos))
    var line string
    for _, p := range order {
      var character rune
      if p >= wordx && p < wordx+len(word) {
        character = rune(word[p-wordx])
      } else {
        character = rune(A + r.Intn(24))
      }
      command := cursorToX(p+2) + string(character)
      line += command
    }
    fmt.Printf("%v", line)
    time.Sleep(10 * time.Millisecond)
  }
  fmt.Printf(clearline + "What word did you see? ")

}

func cursorToX(x int) string {
  return "\r" + e + strconv.Itoa(x) + "C"
}
