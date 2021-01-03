package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

type HasNameAge interface{
	Name() string
	Age() int
}

type Greeter interface{
	HasNameAge
	Greet() string
}

//****************
//**** Member ****
//****************
type Member struct{
	name string
	birthdate time.Time
}

// factory pattern for constructor
func MemberCtor(Name string, Birthdate time.Time) *Member{
	return &Member{name: Name, birthdate: Birthdate}
}

func (m Member) Name() string{
	return m.name
}

func (m Member) Age() int{
	return time.Now().Year() - m.birthdate.Year()
}

func (m Member) Greet() string{
	return "Name: " + m.Name() + ", Age: " + strconv.Itoa(m.Age())
}

//****************
//**** Teacher ***
//****************
type Teacher struct{
	Member
	lecture string
}

func TeacherCtor(Name string, Birthdate time.Time, Lecture string) *Teacher{
	MemberCtor(Name, Birthdate)
	return &Teacher{lecture: Lecture}
}

func (t Teacher) Teaches() string{
	return t.lecture
}
func (t Teacher) Greet() string{
	return  t.Member.Greet()  + ", Teaches: " + t.Teaches()
}

//****************
//**** Student ***
//****************
type Student struct{
	Member
	id_no string
}

func StudentCtor(Name string, Birthdate time.Time, Id_no string) *Student{
	MemberCtor(Name, Birthdate)
	return &Student{id_no: Id_no}
}

func (s Student) ID_no() string{
	return s.id_no
}
func (s Student) Greet() string{
	return  s.Member.Greet()  + ", ID_no: " + s.ID_no()
}

//****************
//*** Visiting ***
//****************
type Visiting struct{
	Student
	valid_thru time.Time
}

func VisitingCtor(Name string, Birthdate time.Time, Id_no string, Valid_thru time.Time) *Visiting{
	StudentCtor(Name, Birthdate, Id_no)
	return &Visiting{valid_thru: Valid_thru}
}

func (v Visiting) expired() bool{
	return time.Now().After(v.valid_thru)
}

func (v Visiting) Greet() string{
	return  v.Student.Greet()  + ", Valid thru: " + strconv.Itoa(v.valid_thru.Year()) + "-" + strconv.Itoa(int(v.valid_thru.Month())) + "-" + strconv.Itoa(v.valid_thru.Day())
}


func main(){
	// time is set as 0
	greeters := []Greeter{Teacher{Member: Member{"Harry", time.Date(1971, 12, 7, 0, 0, 0, 0, time.UTC)}, lecture: "Programming Languages"},
							Teacher{Member: Member{"Natasha", time.Date(1975, 9, 21, 0, 0, 0, 0, time.UTC)}, lecture: "Forbidden Archeology"},
							Student{Member: Member{"YY", time.Date(1999, 3, 16, 0, 0, 0, 0, time.UTC)}, id_no: "2051"},
							Student{Member: Member{"SH", time.Date(2000, 10, 5, 0, 0, 0, 0, time.UTC)}, id_no: "4968"},
							Visiting{Student: Student{Member: Member{"Alice", time.Date(1995, 7, 12, 0, 0, 0, 0, time.UTC)}, id_no: "9595"}, valid_thru: time.Date(2019, 12, 25, 0, 0, 0, 0, time.UTC)},
							Visiting{Student: Student{Member: Member{"Vanessa", time.Date(1998, 3, 27, 0, 0, 0, 0, time.UTC)}, id_no: "9598"}, valid_thru: time.Date(2019, 2, 28, 0, 0, 0, 0, time.UTC)},
	}

	fmt.Println("A few CAU greeters...")
	for _, value := range greeters{
		fmt.Println(value.Greet())
	}

	fmt.Println("\nMembers sorted on age...")
	sort.Slice(greeters, func(i, j int) bool{
		return greeters[i].Age() < greeters[j].Age()
	})
	for _, value := range greeters{
		fmt.Println(value.Greet())
	}

	fmt.Println("\nMembers reverse sorted on name...")
	sort.Slice(greeters, func(i, j int) bool{
		return greeters[i].Name() > greeters[j].Name()
	})
	for _, value := range greeters{
		fmt.Println(value.Greet())
	}
}