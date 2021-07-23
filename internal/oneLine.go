package internal

import (
  "math/rand"
  "strconv"
  "time"
)

const e = "\033["
const colorRed = e + "31m"
const cr = "\r"
const clearline = cr + e + "2K"
const cursorColor = e + "1b[1 q"
const a = 65 // Uppercase a

func OneLine() (word string, captcha string) {
  r := rand.New(rand.NewSource(time.Now().UnixNano()))

  linelength := 32
  captcha += cursorColor
  captcha += clearline + "| " + cursorToX(linelength+3) + "|"
  pos := make([]int, linelength)
  curr := make([]rune, linelength)
  for i := range pos {
    pos[i] = i
  }

  var wordx int
  word = Words[r.Intn(len(Words))]
  // showWord true when we start showing the word
  showWord := false
  for i := 0; i < 50; i++ {
    if i % 10 == 0 {
      wordx = r.Intn(20 - len(word))
      showWord = i > 1
    }
    order := r.Perm(len(pos))
    var line string
    for _, p := range order {
      var character rune
      if showWord {
        // we are showing the word
        if p >= wordx && p < wordx+len(word) {
          // in here write word character
          character = rune(word[p-wordx])
        } else {
          // in here write random char - some chance it's same as already there
          if r.Float32() > 0.5 {
            // chance to keep same char
            character = curr[p]
          } else {
            character = rune(a + r.Intn(24))
            curr[p] = character
          }
        }
      } else {
        // in here we are not showing the word yet
        character = rune(a + r.Intn(24))
        curr[p] = character
      }
      command := cursorToX(p+2) + string(character)
      line += command
    }
    captcha += line
  }

  captcha += clearline

  return word, captcha
}

func cursorToX(x int) string {
  return "\r" + e + strconv.Itoa(x) + "C"
}

