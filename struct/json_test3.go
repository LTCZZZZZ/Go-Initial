package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type StudentAttendanceState struct {
	StudentId    uint   `json:"student_id"`
	StateBefore  uint32 `json:"state_before"`
	StateCurrent uint32 `json:"state_current"`
}

func main() {
	{
		var data map[string]interface{}
		var message = `{"course_timetable_id":8753,"state":1,"student_ids":[9054,9055,9217,9221]}`
		err := json.Unmarshal([]byte(message), &data)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(data)
		course_timetable_id := data["course_timetable_id"] // interface{}类型
		fmt.Println(reflect.TypeOf(course_timetable_id))
		fmt.Println(course_timetable_id.(float64) + 1)
		student_ids := data["student_ids"].([]interface{})
		fmt.Println(reflect.TypeOf(student_ids))
		fmt.Println(reflect.TypeOf(student_ids[0]))

		fmt.Println(data["state"].(float64) == 1)
	}

	{
		fmt.Println("=====================================")
		message := `{"course_timetable_id":10016,"students":[{"student_id":15254,"state_before":2,"state_current":1},{"student_id":457307,"state_before":2,"state_current":1}]}`
		//message := `[{"student_id":15254,"state_before":2,"state_current":1},{"student_id":457307,"state_before":2,"state_current":1}]`
		fmt.Println(message[0:1] == "{")
		//var data struct {
		//	CourseTimetableId uint32                   `json:"course_timetable_id"`
		//	Students          []StudentAttendanceState `json:"students"`
		//} // 这里由于结构复杂，改用了预定义的结构体
		//var data interface{}
		var data []StudentAttendanceState
		err := json.Unmarshal([]byte(message), &data)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(reflect.TypeOf(data))
		fmt.Println("HandleStudentAttendUpdate ", data)
	}

}
