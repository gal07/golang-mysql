package database

import (
	"encoding/json"
	"fmt"
	studentModels "gosql/models"
	util "gosql/utility"
	"strconv"

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

	// bind
	if err := c.Bind(&person); err != nil {

	}

	// validate
	bol := util.Validates(c, person)
	if bol {
		return
	}

	// Transaction begin
	tx, err := db.BeginTx(c, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tx.Rollback()

	// Execute
	_, err = db.ExecContext(c, "insert into tb_student (name,age,grade) values (?, ?, ?)", &person.Name, &person.Age, &person.Grade)
	if err != nil {
		ResponMsg.Message = "Insert Failed"
		ResponMsg.Error = err
	} else {
		ResponMsg.Message = "Insert Completed"
		tx.Commit()
	}

	c.JSON(200, gin.H{
		"response": ResponMsg,
	})
}

func Update(id int, c *gin.Context) {

	var person studentModels.Student
	db, err := connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// bind
	if err := c.Bind(&person); err != nil {

	}

	// validate
	bol := util.Validates(c, person)
	if bol {
		return
	}

	// Transaction begin
	tx, err := db.BeginTx(c, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tx.Rollback()

	// Execute
	_, err = db.ExecContext(c, "UPDATE tb_student SET name = ?,age = ?,grade = ? WHERE id = ?", &person.Name, &person.Age, &person.Grade, id)
	if err != nil {
		ResponMsg.Message = "Update Failed"
		ResponMsg.Error = err
	} else {
		ResponMsg.Message = "Update Completed"
		tx.Commit()
	}

	c.JSON(200, gin.H{
		"response": ResponMsg,
	})
}

func Delete(id int, c *gin.Context) {

	var delete studentModels.Delete
	db, err := connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// bind
	if err := c.Bind(&delete); err != nil {

	}

	// validate
	bol := util.Validates(c, delete)
	if bol {
		return
	}

	// Transaction begin
	tx, err := db.BeginTx(c, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tx.Rollback()

	//match id from uri and bod
	if id != delete.ID {
		ResponMsg.Message = "ID not match"
	} else {

		// Execute
		_, err = db.ExecContext(c, "UPDATE tb_student SET isdelete = ? WHERE id = ?", 1, id)
		if err != nil {
			ResponMsg.Message = "Deleted Failed"
			ResponMsg.Error = err
		} else {
			ResponMsg.Message = "Deleted Completed"
			tx.Commit()
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

		rows, err := db.Query("select id,name,age,grade from tb_student LIMIT ? OFFSET ?", 100, 1)
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

		err := db.QueryRow("select id,name,age,grade from tb_student where id = ? and isdelete = ?", id, 0).Scan(&student.ID, &student.Name, &student.Age, &student.Grade)
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

	var total studentModels.MassInsert
	db, err := connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// bind
	if err := c.Bind(&total); err != nil {

	}

	// validate
	bol := util.Validates(c, total)
	if bol {
		return
	}

	// Transaction begin
	tx, err := db.BeginTx(c, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	t := total.Total
	fmt.Println(t)
	for i := 0; i < t; i++ {
		result, err := tx.ExecContext(c, "insert into tb_student (name,age,grade) values (?, ?, ?)", util.RandomFullname(), util.RandomID(20), util.RandomID(20))
		if err != nil {
			fmt.Println("failed insert ke - ", i)
			ResponMsg.Message = "Insert Failed"
			ResponMsg.Error = err
		} else {
			fmt.Println("success insert ke - ", i)
		}
		fmt.Println(result.LastInsertId())
		ResponMsg.Message = "Insert Completed " + strconv.Itoa(i)

	}

	tx.Commit()

	c.JSON(200, gin.H{
		"response": ResponMsg,
	})
}

func Search(c *gin.Context) {

	var student []studentModels.Student
	var search studentModels.Search

	// connection
	db, err := connect()
	if err != nil {
		fmt.Println(err)
		return
	}

	// close connection
	defer db.Close()

	// bind
	if err := c.Bind(&search); err != nil {

	}

	// validate
	bol := util.Validates(c, search)
	if bol {
		return
	}

	findstudent := "%" + search.Search + "%"

	// Querying to find spesific student
	rows, err := db.Query("SELECT id,name,grade,age from tb_student where name LIKE ? and isdelete = ?", findstudent, 0)
	if err != nil {
		fmt.Println(err)
		ResponMsg.Message = err.Error()
		return
	} else {
		ResponMsg.Message = "Get Student"
	}

	if !rows.Next() {
		ResponMsg.Message = "Not Found"
	}
	// re-arrange and append student data
	for rows.Next() {
		var each = studentModels.Student{}
		var err = rows.Scan(&each.ID, &each.Name, &each.Grade, &each.Age)
		if err != nil {
			fmt.Println(err)
			ResponMsg.Message = "Get Failed"
			ResponMsg.Error = err
		} else {
			ResponMsg.Message = "Get Completed"
			student = append(student, each)
		}
	}

	// response
	c.JSON(200, gin.H{
		"response": ResponMsg,
		"data":     student,
	})
}
