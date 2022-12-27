package service

import (
	"TaskFinal/model"
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"

	"TaskFinal/action"
)

const (
	USER_NAME   = "IpatovDmitriy"
	CIRCLE_TASK = "Циклическая ротация"
	COUPLE_TASK = "Чудные вхождения в массив"
	SEQ_TASK    = "Проверка последовательности"
	MISS_TASK   = "Поиск отсутствующего элемента"
)

func BuildRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/task", func(r chi.Router) {
		r.Route("/{taskName}", func(r chi.Router) {
			r.Use(taskCtx)
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				answer, err := resolveTask(r.Context().Value("taskName").(string))
				Make(w, answer, err)
			})
		})
	})
	r.Get("/tasks", func(w http.ResponseWriter, r *http.Request) {
		taskNames := make([]string, 0)
		taskNames = append(taskNames, CIRCLE_TASK, COUPLE_TASK, SEQ_TASK, MISS_TASK)
		length := len(taskNames)
		result := make([]model.Answer, length)
		var err error
		for i := 0; i < length; i++ {
			result[i], err = resolveTask(taskNames[i])
			if err != nil {
				Make(w, result, nil)
				return
			}
		}
		Make(w, result, nil)
	})
	return r
}

func taskCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "taskName", chi.URLParam(r, "taskName"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func resolveTask(taskName string) (answer model.Answer, err error) {
	taskValue, tasks, err := action.GetTask(taskName)
	if err != nil {
		return model.Answer{}, err
	}
	length := len(taskValue.Tasks)
	check := model.Check{
		UserName: USER_NAME,
		Task:     taskName,
		TaskResults: model.Result{
			Payload: tasks,
			Results: make([]any, length),
		},
	}
	switch taskName {
	case CIRCLE_TASK:
		for i := 0; i < length; i++ {
			check.TaskResults.Results[i] = action.SolutionCircle(taskValue.Tasks[i].A, taskValue.Tasks[i].K)
		}
	case COUPLE_TASK:
		for i := 0; i < len(taskValue.Tasks); i++ {
			check.TaskResults.Results[i] = action.SolutionCouple(taskValue.Tasks[i].A)
		}
	case SEQ_TASK:
		for i := 0; i < len(taskValue.Tasks); i++ {
			check.TaskResults.Results[i] = action.SolutionSeq(taskValue.Tasks[i].A)
		}
	case MISS_TASK:
		for i := 0; i < len(taskValue.Tasks); i++ {
			check.TaskResults.Results[i] = action.SolutionMiss(taskValue.Tasks[i].A)
		}
	}
	answer, err = action.SendTask(check)
	if err != nil {
		return answer, err
	}
	answer.TaskName = taskName
	return answer, nil
}
