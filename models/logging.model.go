package models

import (
	"fmt"
)

const (
	operatinglogpath = "samurainegi/log/operatelog"
)

type OperatingLog struct {
	Who  string `json:"who,omitempty" firestore:"who,omitempty"`
	Did  string `json:"did,omitempty" firestore:"did,omitempty"`
	What string `json:"what,omitempty" firestore:"what,omitempty"`
	When string `json:"when,omitempty" firestore:"when,omitempty"`
}

func (OL OperatingLog) LoggingOperating() {
	ctx, client := Setup()
	defer client.Close()
	logCol := client.Collection(operatinglogpath)
	doc, wr, err := logCol.Add(ctx, OL)
	if err != nil {
		// TODO: Handle error.
		fmt.Println(err)
	}
	fmt.Println("add operating log:", doc, wr)
}
