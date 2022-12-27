package action

import (
	"TaskFinal/model"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetTask(taskName string) (task model.Task, tasks []json.RawMessage, err error) {
	response, err := http.Get("https://kuvaev-ituniversity.vps.elewise.com/tasks/" + taskName)
	if err != nil {
		return model.Task{}, nil, err
	}
	var body []byte
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return model.Task{}, nil, err
	}
	// делим на задачи
	return parseTasks(body)
}

func SendTask(task model.Check) (answer model.Answer, err error) {
	taskJson, err := json.Marshal(task)
	if err != nil {
		return model.Answer{}, err
	}
	response, err := http.Post("https://kuvaev-ituniversity.vps.elewise.com/tasks/solution", "application/json", bytes.NewBuffer(taskJson))
	if err != nil {
		return model.Answer{}, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return model.Answer{}, err
	}
	answerModel := model.Answer{}
	err = json.Unmarshal(body, &answerModel)
	if err != nil {
		return model.Answer{}, err
	}
	return answerModel, nil
}

func parseTasks(body []byte) (task model.Task, tasks []json.RawMessage, err error) {
	err = json.Unmarshal(body, &tasks)
	if err != nil {
		return model.Task{}, nil, err
	}
	length := len(tasks)
	taskValues := make([]model.TaskInfo, length)
	for i := 0; i < length; i++ {
		var taskValue []json.RawMessage
		err = json.Unmarshal(tasks[i], &taskValue)
		if err != nil {
			return model.Task{}, nil, err
		}
		var array []int
		var count int
		if len(taskValue) == 2 {
			// циклическая ротация
			err = json.Unmarshal(taskValue[1], &count)
			if err != nil {
				return model.Task{}, nil, err
			}
		} else {
			count = -1
		}
		err = json.Unmarshal(taskValue[0], &array)
		if err != nil {
			return model.Task{}, nil, err
		}
		taskValues[i] = model.TaskInfo{
			A: array,
			K: count,
		}
	}
	return model.Task{Tasks: taskValues}, tasks, nil
}
