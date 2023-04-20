package storage

import(
	"testing"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestCreateUser(t *testing.T)() {
	InitStorage("tests.db")
	GetUser(User{ID: "1", Username: "Barbie", Email: "foo@gmail.com"})
	t.Logf("CurrentUser is %+v", CurrentUser)
	GetUser(User{ID: "2", Username: "Jane", Email: "bar@gmail.com"})
	t.Logf("CurrentUser is %+v", CurrentUser)
	GetUser(User{ID: "1"})
	if !cmp.Equal(CurrentUser, User{ID: "1", Username: "Barbie", Email: "foo@gmail.com"}) {
		t.Errorf("Expected CurrentUser %+v, got CurrentUser %+v", User{ID: "1", Username: "Barbie", Email: "foo@gmail.com"}, CurrentUser)
	}
}

func TestSubmitFortune(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	InitStorage("tests.db")
	GetUser(User{ID: fmt.Sprintf("%d", rand.Int()+100), Username: "John", Email: "example@gmail.com"})
	if err := SubmitFortune(fmt.Sprintf("%s%s wishes you good health and prosperity!", CurrentUser.Username, CurrentUser.ID)); err != nil {
		t.Errorf("Expected nil error but got %q", err.Error())
	}
	if err := SubmitFortune(fmt.Sprintf("%s%s wishes you good health and prosperity!", CurrentUser.Username, CurrentUser.ID)); err == nil {
		t.Errorf("Expected submit to fail but got nil error")
	}
}

func TestReceiveFortune(t *testing.T) {
	InitStorage("tests.db")
	GetUser(User{ID: "1"})
	t.Logf("CurrentUser is %s", CurrentUser.Username)
	for i := 0; i < 3; i++ {
		fortune, err := ReceiveFortune()
		if err != nil {
			t.Errorf("Expected to receive fortune but received error %q", err.Error())
		}
		t.Logf("Received fortune %q", fortune.Content)
	}
	received, err := GetReceivedFortunes()
	if err != nil {
		t.Errorf("Expected no error but instead received error %q", err.Error())
	}
	t.Logf("User fortune history is %+v", received)
}

func TestGetSubmit(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	InitStorage("tests.db")
	GetUser(User{ID: fmt.Sprintf("%d", rand.Int()+100), Username: "Catherine", Email: "getsubmit@gmail.com"})
	t.Logf("CurrentUser is %s", CurrentUser.Username)
	t.Logf("CurrentUser submitted is %t", CurrentUser.Submitted)
	if GetSubmit() == true {
		t.Error("Expected Submitted to be false but it was true")
	}
	if err := SubmitFortune(fmt.Sprintf("%s%s wishes you good health and prosperity!", CurrentUser.Username, CurrentUser.ID)); err != nil {
		t.Errorf("Expected nil error but got %q", err.Error())
	}
	if GetSubmit() == false {
		t.Error("Expected CanSubmit to be true but it was false")
	}
}

func TestGetLastTime(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	InitStorage("tests.db")
	GetUser(User{ID: fmt.Sprintf("%d", rand.Int()+100), Username: "Mac", Email: "getlasttime@gmail.com"})
	t.Logf("CurrentUser is %s", CurrentUser.Username)
	if err := SubmitFortune(fmt.Sprintf("%s%s wishes you good health and prosperity!", CurrentUser.Username, CurrentUser.ID)); err != nil {
		t.Errorf("Expected nil error but got %q", err.Error())
	}
	t.Logf("LastTime is %s", CurrentUser.LastTime)
	if CurrentUser.LastTime != GetLastTime() {
		t.Errorf("Expected LastTime and GetLastTime to return the same time")
	}
}

