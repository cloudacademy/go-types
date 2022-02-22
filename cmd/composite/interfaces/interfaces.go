package main

import (
	"fmt"
	"time"
)

type activity interface {
	walk(int) error
	talk(string) error
	sleep()
}

type person struct {
	name   string
	age    int
	asleep bool
	activity
}

type experimentFunc func() string

type science interface {
	doScience() string
}

type scientist struct {
	person
	experiment experimentFunc
	science
}

func (p person) walk(distance int) error {
	if !p.asleep {
		fmt.Printf("walking %d...\n", distance)
		return nil
	} else {
		return fmt.Errorf("%s is asleep", p.name)
	}
}

func (p person) talk(speech string) error {
	if !p.asleep {
		fmt.Printf("%s...\n", speech)
		return nil
	} else {
		return fmt.Errorf("%s is asleep", p.name)
	}
}

func (p *person) sleep() {
	if !p.asleep {
		p.asleep = true
		fmt.Println("going to sleep...")
		time.Sleep(1 * time.Second)
		p.asleep = false
		fmt.Println("awake...")
	}
}

func (s *scientist) doScience() string {
	return s.experiment()
}

var measureSpeedOfLight experimentFunc = func() string { return "measuring the speed of light..." }
var splitAtom experimentFunc = func() string { return "splitting the atom..." }

func main() {
	s1 := scientist{}
	s1.name = "Albert Einstein"
	s1.age = 30
	s1.experiment = measureSpeedOfLight

	s2 := scientist{}
	s2.name = "Ernst Rutherford"
	s2.age = 30
	s2.experiment = splitAtom

	scientists := []*scientist{&s1, &s2}

	for _, s := range scientists {
		fmt.Println()
		fmt.Printf("%s\n", s.name)
		s.walk(100)
		s.walk(100)
		s.talk("finished walking")
		s.sleep()
		fmt.Printf("%s\n", s.doScience())
	}

	fmt.Println()
}
