package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main()  {
	var data map[string]interface{}
	var message = `{"course_timetable_id":8753,"state":1,"student_ids":[9054,9055,9217,9221]}`
	err := json.Unmarshal([]byte(message), &data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
	course_timetable_id := data["course_timetable_id"]  // interface{}类型
	fmt.Println(reflect.TypeOf(course_timetable_id))
	fmt.Println(course_timetable_id.(float64) + 1)
	student_ids := data["student_ids"].([]interface{})
	fmt.Println(reflect.TypeOf(student_ids))
	fmt.Println(reflect.TypeOf(student_ids[0]))

	fmt.Println(data["state"].(float64) == 1)

}
