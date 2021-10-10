package parseTasks

import (
	"encoding/json"
	"io/ioutil"
	"os"
//	"fmt"
	log "github.com/sirupsen/logrus"
)


type Tasks struct {
	Tasks []Task `json:"tasks"`
}
type Task struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Flags []string `json:"flags"`
}

func ParseTasks() Tasks {
	log.SetFormatter(&log.JSONFormatter{})
	jsonfile, err := os.Open("todo.json")

  if err != nil {
    log.Info(err)
  }
//	fmt.Println("We've succesfully opened json file")

	// read opened file as a byte array
	bytevalue, _ := ioutil.ReadAll(jsonfile)

	var tasks Tasks
	json.Unmarshal(bytevalue, &tasks)
	//fmt.Println(json.Unmarshal(bytevalue, &tasks))


//	for i := 0; i < len(tasks.Tasks); i++ {
//		fmt.Println("task id: " + tasks.Tasks[i].ID)
//		fmt.Println("task title: " + tasks.Tasks[i].Title)
//		fmt.Println("task description: " + tasks.Tasks[i].Description)
//		//				fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
//	}
	return tasks
}
