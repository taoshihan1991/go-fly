package tools

import "fmt"

func MyTest() {
	type MConn struct {
		Name string
	}
	var conn *MConn
	var conn2 MConn
	conn3 := new(MConn)
	conn4 := &MConn{}
	fmt.Printf("%v,%v,%v,%v \r\n", conn, conn2, conn3, conn4)

	var mMap map[string][]*MConn
	m1, _ := mMap["name"]
	//if ok {
	//	m1.Name = "qqq"
	//}
	fmt.Printf("ssss%T", m1)
}
func MyStruct() {
	type s2 struct {
		name string
	}
	aa := s2{
		name: "aa",
	}
	bb := s2{
		name: "aa",
	}
	fmt.Printf("%v\n", aa == bb)

	type s1 struct {
		one   map[string]string
		two   []string
		three string
	}

	a := &s1{
		one:   map[string]string{"aaa": "bbb"},
		two:   []string{"aaa", "bbb"},
		three: "aaaa",
	}
	b := &s1{
		one:   map[string]string{"aaa": "bbb"},
		two:   []string{"aaa", "bbb"},
		three: "aaaa",
	}
	c := a
	fmt.Printf("%v;%v", a == b, a == c)
}
