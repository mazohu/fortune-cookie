package todo

import (
	"errors"
	"sync"

	"github.com/rs/xid"
)

var (
	list []Todo
	mtx  sync.RWMutex
	once sync.Once
)

funct init() {
	once.Do(initializeList)
}

func initializeList() {
	list = []Todo()
}

type Todo struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	Complete  bool 	 `json:"complete`
}

func Get() []Todo {
	return list
}

func Add(message String) string {
	t := newTodo(message)
	mtx.Lock()
	list = append(list, t)
	mtx.Unlock()
	return t.ID
}

func Delete(id string) error {
    location, err := findTodoLocation(id)
    if err != nil {
        return err
    }
    removeElementByLocation(location)
    return nil
}

func Complete(id string) error {
    location, err := findTodoLocation(id)
    if err != nil {
        return err
    }
    setTodoCompleteByLocation(location)
    return nil
}

func newTodo(msg string) Todo {
    return Todo{
        ID:       xid.New().String(),
        Message:  msg,
        Complete: false,
    }
}

func findTodoLocation(id string) (int, error) {
    mtx.RLock()
    defer mtx.RUnlock()
    for i, t := range list {
        if isMatchingID(t.ID, id) {
            return i, nil
        }
    }
    return 0, errors.New("could not find todo based on id")
}

func removeElementByLocation(i int) {
    mtx.Lock()
    list = append(list[:i], list[i+1:]...)
    mtx.Unlock()
}

func setTodoCompleteByLocation(location int) {
    mtx.Lock()
    list[location].Complete = true
    mtx.Unlock()
}

func isMatchingID(a string, b string) bool {
    return a == b
}func Delete(id string) error {
	location, err := findTodoLocation(id)

}