package database

import "time"

type CommandsPending struct {
	ChatID  int       `bson:"chat_id"`
	Steps   int       `bson:"steps"`
	Command string    `bson:"command"`
	Process int       `bson:"process"`
	Date    time.Time `bson:"date"`
}
