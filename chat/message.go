package main
import (
  "time"
)

//messageは1つのメッセージを表しマス。
type message struct {
  Name string
  Message string
  When time.Time
}
