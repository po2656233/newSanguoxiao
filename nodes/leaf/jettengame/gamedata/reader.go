package gamedata

import (
	"reflect"

	"github.com/po2656233/goleaf/log"
	"github.com/po2656233/goleaf/recordfile"
)

func readRf(st interface{}) *recordfile.RecordFile {
	rf, err := recordfile.New(st)
	if err != nil {
		log.Fatal("%v", err)
	}
	fn := reflect.TypeOf(st).Name() + ".txt"
	err = rf.Read("gamedata/" + fn)
	if err != nil {
		log.Fatal("%v: %v", fn, err)
	}

	return rf
}

//
//func Example() {
//	type Record struct {
//		// index 0
//		IndexInt int32 "index"
//		// index 1
//		IndexStr string "index"
//		_Number  int32
//		Str      string
//		Arr1     [2]int
//		Arr2     [3][2]int
//		Arr3     []int
//		St       struct {
//			Name string "name"
//			Num  int32    "num"
//		}
//		M map[string]int
//	}
//
//	rf, err := recordfile.New(Record{})
//	if err != nil {
//		return
//	}
//
//	err = rf.Read("gamedata/test.txt")
//	if err != nil {
//		return
//	}
//
//	for i := 0; i < rf.NumRecord(); i++ {
//		r := rf.Record(i).(*Record)
//		fmt.Println(r.IndexInt)
//	}
//
//	r := rf.Index(2).(*Record)
//	fmt.Println(r.Str)
//
//	r = rf.Indexes(0)[2].(*Record)
//	fmt.Println(r.Str)
//
//	r = rf.Indexes(1)["three"].(*Record)
//	fmt.Println(r.Str)
//	fmt.Println(r.Arr1[1])
//	fmt.Println(r.Arr2[2][0])
//	fmt.Println(r.Arr3[0])
//	fmt.Println(r.St.Name)
//	fmt.Println(r.M["key6"])
//
//	// Output:
//	// 1
//	// 2
//	// 3
//	// cat
//	// cat
//	// book
//	// 6
//	// 4
//	// 6
//	// name5566
//	// 6
//}
