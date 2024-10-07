package database

import (
	"encoding/json"
	"fmt"
	studentModels "gosql/models"
	util "gosql/utility"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var ResponMsg = studentModels.Responses{}

func GetJSON(sqlString string, taskID string) (string, error) {
	db, err := connect()
	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	defer db.Close()

	rows, err := db.Query(sqlString, taskID)
	if err != nil {
		return "", err
	}

	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}

	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

func Insert(c *gin.Context) {

	var person studentModels.Student
	db, err := connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	if c.ShouldBind(&person) == nil {
		_, err = db.Exec("insert into tb_student (name,age,grade) values (?, ?, ?)", &person.Name, &person.Age, &person.Grade)
		if err != nil {
			ResponMsg.Message = "Insert Failed"
			ResponMsg.Error = err
		} else {
			ResponMsg.Message = "Insert Completed"
		}
	}

	c.JSON(200, gin.H{
		"response": ResponMsg,
	})
}

func Get(c *gin.Context) {

	var student []studentModels.Student
	db, err := connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	if c.ShouldBind(&student) == nil {

		rows, err := db.Query("select id,name,age,grade from tb_student")
		if err != nil {
			fmt.Println(err)
			return
		}

		for rows.Next() {
			var each = studentModels.Student{}
			var err = rows.Scan(&each.ID, &each.Name, &each.Age, &each.Grade)
			if err != nil {
				ResponMsg.Message = "Get Failed"
				ResponMsg.Error = err
			} else {
				ResponMsg.Message = "Get Completed"
				student = append(student, each)
			}
		}

		c.JSON(200, gin.H{
			"response": ResponMsg,
			"data":     student,
		})

	}
}

func GetDetail(id string, c *gin.Context) {

	var student = studentModels.Student{}
	db, err := connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	if c.ShouldBind(&student) == nil {

		err := db.QueryRow("select id,name,age,grade from tb_student where id = ?", id).Scan(&student.ID, &student.Name, &student.Age, &student.Grade)
		if err != nil {
			ResponMsg.Message = err.Error()
		}
	}

	c.JSON(200, gin.H{
		"response": ResponMsg,
		"data":     student,
	})
}

func MassInsert(c *gin.Context) {

	var person studentModels.Student
	db, err := connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	if c.ShouldBind(&person) == nil {

		t := person.Total
		for i := 0; i < t; i++ {
			_, err = db.Exec("insert into tb_student (name,age,grade) values (?, ?, ?)", util.RandomFullname(), util.RandomID(20), util.RandomID(20))
			if err != nil {
				ResponMsg.Message = "Insert Failed"
				ResponMsg.Error = err
			} else {
				ResponMsg.Message = "Insert Completed"
			}
		}

	}

	c.JSON(200, gin.H{
		"response": ResponMsg,
	})
}
