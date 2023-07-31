package day_2023_07

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "speak" {
		talk = "speak"
	} else {
		talk = "Hi"
	}
	return
}

func Day_2023_07_28() {
	// var peo People = Student{}
	var peo People = &Student{}
	fmt.Println(peo.Speak("speak"))

}
