package main

import (
  "log"
  "os"

  "github.com/faiface/beep"
)

func main() {
  f, err := os.Open("hakucyou.mp3")
  if err != nil {
    log.Fatal(err)
  }
  
  streamer, format, err := mp3.Decode(f)
  if err != nil {
    log.Fatal(err)
  }
  defer streamer.Close()

  speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

  speaker.Play(streamer)
  select {}
}
