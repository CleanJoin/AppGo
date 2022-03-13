package page

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Member struct {
	APIToken string `json:"api_token"`
	Member   struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"member"`
}
type Server struct {
}
type Erorr struct {
	Erorrs string `json:"error"`
}

var Player = new(Member)

var Err = new(Erorr)
var Logon = new(User)

func LogingUser(login string, password string) (int, string) {
	Logon.Username = login
	Logon.Password = password
	jsonValue, _ := json.Marshal(Logon)
	req, _ := http.Post("http://localhost:10000/api/login", "application/json", bytes.NewBuffer(jsonValue))

	body, _ := ioutil.ReadAll(req.Body)
	json.NewDecoder(req.Body).Decode(&Err)
	json.Unmarshal(body, Err)
	json.Unmarshal(body, Player)

	return req.StatusCode, Err.Erorrs
}

func RegisterUser(login string, password string) (int, string) {
	Logon.Username = login
	Logon.Password = password
	jsonValue, _ := json.Marshal(Logon)
	req, _ := http.Post("http://localhost:10000/api/user", "application/json", bytes.NewBuffer(jsonValue))

	body, _ := ioutil.ReadAll(req.Body)
	json.NewDecoder(req.Body).Decode(&Err)
	json.Unmarshal(body, Err)
	return req.StatusCode, Err.Erorrs
}

func Form1(myWindow fyne.Window, myApp fyne.App) {
	login := widget.NewEntry()
	login.SetPlaceHolder("Enter Login...")
	var grid *fyne.Container
	password := widget.NewEntry()
	password.SetPlaceHolder("Enter Password...")
	password.Password = true

	content := container.NewVBox(login, password, widget.NewButton("Login", func() {
		code, err := LogingUser(login.Text, password.Text)
		if code == 200 {
			chatWindow := myApp.NewWindow("Chat")
			Form2(chatWindow, myApp)
			myWindow.Close()

		} else {
			dialog.ShowInformation("Ошибка", err, myWindow)
		}

	}), widget.NewButton("Register", func() {
		code, err := RegisterUser(login.Text, password.Text)
		if code == 201 {
			dialog.ShowInformation("Successful", fmt.Sprintf("Пользователя %s зарегистрировали", login.Text), myWindow)

		} else {
			dialog.ShowInformation("Error", err, myWindow)

		}

	}))

	grid = container.New(layout.NewGridWrapLayout(fyne.NewSize(500, 170)),
		content)
	myWindow.SetContent(grid)
	myWindow.Show()
}
