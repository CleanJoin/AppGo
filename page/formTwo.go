package page

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type MessagesData struct {
	Messages []Message `json:"messages"`
}
type Message struct {
	ID         int       `json:"id"`
	MemberName string    `json:"member_name"`
	Text       string    `json:"text"`
	Time       time.Time `json:"time"`
}

type MembersData struct {
	Members []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"members"`
}

var MembersNew = new(MembersData)
var MessagesNew = new(MessagesData)

func Form2(chatWindow fyne.Window, myApp fyne.App) {

	membersForm := widget.NewList(
		func() int {

			return len(MembersNew.Members)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(MembersNew.Members[i].Name)
		})

	messageForm := widget.NewList(
		func() int {

			return len(MessagesNew.Messages)
		},
		func() fyne.CanvasObject {

			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(fmt.Sprintf("%s (%s): %s", MessagesNew.Messages[i].MemberName, MessagesNew.Messages[i].Time.Format("15:04:05"), MessagesNew.Messages[i].Text))
			o.(*widget.Label).Resize(fyne.NewSize(1530, 100))

		})

	inputMessage := widget.NewEntry()
	inputMessage.SetPlaceHolder("Enter Text...")
	//fmt.Println(i)
	textForm := container.NewVBox(inputMessage, widget.NewButton("Отправить", func() {
		SendMessage(Player, inputMessage.Text)
		inputMessage.Text = ""
	}))

	members := container.New(layout.NewGridWrapLayout(fyne.NewSize(250, 800)), membersForm)
	centeredMessage := container.New(layout.NewGridWrapLayout(fyne.NewSize(1530, 800)), messageForm)
	membersAndMessages := container.New(layout.NewHBoxLayout(), members, layout.NewSpacer(), centeredMessage)
	messageBox := container.New(layout.NewGridWrapLayout(fyne.NewSize(1700, 100)), textForm)
	centeredFormMessage := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), messageBox, layout.NewSpacer())
	fullForm := (container.New(layout.NewVBoxLayout(), membersAndMessages, centeredFormMessage))

	chatWindow.SetContent(fullForm)
	chatWindow.Show()

	go func() {
		for {
			GetMembers(Player)
			code, err, _ := GetMessages(Player)
			if code != 200 {
				dialog.ShowConfirm("Ошибка", err, func(b bool) {
					myWindow := myApp.NewWindow("Chat")
					Form1(myWindow, myApp)
					chatWindow.Close()
				}, chatWindow)
			}
			chatWindow.Content().Refresh()
			time.Sleep(time.Second + 25)
		}
	}()

}

func GetMembers(player *Member) (int, string, *MembersData) {
	req, _ := http.Post("http://localhost:10000/api/members", "application/json", strings.NewReader(fmt.Sprintf(`{"api_token":"%s"}`, player.APIToken)))
	body, _ := ioutil.ReadAll(req.Body)
	json.NewDecoder(req.Body).Decode(&Err)
	json.Unmarshal(body, Err)
	json.Unmarshal(body, MembersNew)
	// fmt.Println(MembersNew)

	return req.StatusCode, Err.Erorrs, MembersNew
}

func GetMessages(player *Member) (int, string, *MessagesData) {
	req, _ := http.Post("http://localhost:10000/api/messages", "application/json", strings.NewReader(fmt.Sprintf(`{"api_token":"%s"}`, player.APIToken)))
	body, _ := ioutil.ReadAll(req.Body)
	json.NewDecoder(req.Body).Decode(&Err)
	json.Unmarshal(body, Err)
	json.Unmarshal(body, MessagesNew)
	// fmt.Println(MessagesNew)
	return req.StatusCode, Err.Erorrs, MessagesNew
}

func SendMessage(player *Member, text string) (int, string) {
	req, _ := http.Post("http://localhost:10000/api/message", "application/json", strings.NewReader(fmt.Sprintf(`{"api_token":"%s","text":"%s"}`, player.APIToken, text)))
	body, _ := ioutil.ReadAll(req.Body)
	json.NewDecoder(req.Body).Decode(&Err)
	json.Unmarshal(body, Err)
	json.Unmarshal(body, Player)

	return req.StatusCode, Err.Erorrs
}
