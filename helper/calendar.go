package helper

import (
	"Gurumu/config"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func GetClient(config *oauth2.Config) *http.Client {
	tokFile := "./helper/temporary/token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		SaveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func SaveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func CalendarEvent(emailGuru, emailSiswa, pelajaran, tanggal, jam string) (string, error) {

	layout := "2006-01-02 15:04:05"
	value := tanggal + " " + jam + ":00"
	dateTime, err := time.Parse(layout, value)
	if err != nil {
		return "", errors.New("failed convert datetime")
	}
	setEvent := &calendar.Event{
		Summary:     "Gurumu - Kelas " + pelajaran + "anda",
		Description: "Kelas akan berlangsung pada " + tanggal + " pada " + jam + ". Harap datang tepat waktu dan pastikan untuk bergabung dengan panggilan video tepat waktu.",
		ConferenceData: &calendar.ConferenceData{
			CreateRequest: &calendar.CreateConferenceRequest{
				RequestId: "sfsfs",
				ConferenceSolutionKey: &calendar.ConferenceSolutionKey{
					Type: "hangoutsMeet"},
				Status: &calendar.ConferenceRequestStatus{
					StatusCode: "success"},
			}},
		Start: &calendar.EventDateTime{
			DateTime: dateTime.Format(time.RFC3339),
			TimeZone: "Asia/Jakarta",
		},
		End: &calendar.EventDateTime{
			DateTime: dateTime.Add(time.Hour * 1).Format(time.RFC3339),
			TimeZone: "Asia/Jakarta",
		},
		Attendees: []*calendar.EventAttendee{
			{Email: emailGuru},
			{Email: emailSiswa},
		},
		Reminders: &calendar.EventReminders{
			UseDefault: true,
		},
	}

	ctx := context.Background()

	client_id := config.GOOGLE_OAUTH_CLIENT_ID1
	project := config.GOOGLE_PROJECT_ID1
	secret := config.GOOGLE_OAUTH_CLIENT_SECRET1
	b := `{"web":{"client_id":"` + client_id + `","project_id":"` + project + `","auth_uri":"https://accounts.google.com/o/oauth2/auth","token_uri":"https://oauth2.googleapis.com/token","auth_provider_x509_cert_url":"https://www.googleapis.com/oauth2/v1/certs","client_secret":"` + secret + `","redirect_uris":["http://localhost:8000/callback"]}}`
	fmt.Println(b)
	bt := []byte(b)
	fmt.Println(bt)

	config, err := google.ConfigFromJSON(bt, calendar.CalendarEventsScope, calendar.CalendarScope, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := GetClient(config)
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	calendarId := "primary"
	resEvent, err := srv.Events.Insert(calendarId, setEvent).SendUpdates("all").ConferenceDataVersion(1).Do()

	if err != nil {
		log.Fatalf("Unable to create event. %v\n", err)
	}
	fmt.Printf("Event created: %s\n", resEvent.HtmlLink)
	fmt.Println(resEvent.ConferenceData.ConferenceId)

	return resEvent.HtmlLink, nil
}
