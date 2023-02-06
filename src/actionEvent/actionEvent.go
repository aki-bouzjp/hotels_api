package actionEvent

import (
	"app/src/logger"
	"app/src/util"
	"regexp"
	"strconv"

	"fmt"
	"os"
	"time"

	"github.com/gocarina/gocsv"
)

type EventLog struct {
	Cps       string `json:"cps"`
	Type      string `json:"type"`
	Action    string `json:"action"`
	CreatedAt string `json:"created_at"`
}
type ActionEvent struct {
	sessions  []*Session
	selects   []*Select
	feedbacks []*Feedback
}

const (
	watchPeriod    = 3000 * time.Millisecond
	writeWait      = 30 * time.Second
	pongWait       = 30 * time.Second
	timeoutWait    = 60 * time.Second * 5
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
	// timeoutWait    = 60 * time.Second * 15
)

func New() *ActionEvent {
	return &ActionEvent{
		sessions:  []*Session{},
		selects:   []*Select{},
		feedbacks: []*Feedback{},
	}
}

func (a *ActionEvent) Run() {
	ticker := time.NewTicker(watchPeriod)
	done := make(chan bool)
	defer func() {
		done <- true
		ticker.Stop()
	}()

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				time := time.Now()
				filePath := fmt.Sprintf(
					"tmp/%d%02d%02d%02d%02d%02d",
					time.Year(),
					time.Month(),
					time.Day(),
					time.Hour(),
					time.Minute(),
					time.Second(),
				)
				a.WriteSessionCsv(filePath)
				a.WriteSelectCsv(filePath)
				a.WriteFeedbackCsv(filePath)
			}
		}
	}()

	// done <- true

	// ticker := time.NewTicker(500 * time.Millisecond)
	// done := make(chan bool)
	// go func() {
	// 	for {
	// 		select {
	// 		case <-done:
	// 			return
	// 		case t := <-ticker.C:
	// 			fmt.Println("Tick at", t)
	// 		}
	// 	}
	// }()
	// time.Sleep(1600 * time.Millisecond)
	// ticker.Stop()
	// done <- true
	// fmt.Println("Ticker stopped")
}

func (a *ActionEvent) WriteSessionCsv(filePath string) {
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		logger.Error("Failed to create OpenFile to write session csv. filePath: %v, err: %v", filePath, err)
	}
	defer f.Close()
	if _, err = gocsv.MarshalString(a.sessions); err != nil {
		logger.Error("Failed to write session csv. filePath: %v, err: %v", filePath, err)
	}
	a.sessions = []*Session{}
}

func (a *ActionEvent) WriteSelectCsv(filePath string) {
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		logger.Error("Failed to create OpenFile to write select csv. filePath: %v, err: %v", filePath, err)
	}
	defer f.Close()
	if _, err = gocsv.MarshalString(a.selects); err != nil {
		logger.Error("Failed to write select csv. filePath: %v, err: %v", filePath, err)
	}
	a.selects = []*Select{}
}

func (a *ActionEvent) WriteFeedbackCsv(filePath string) {
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		logger.Error("Failed to create OpenFile to write feedback csv. filePath: %v, err: %v", filePath, err)
	}
	defer f.Close()
	if _, err = gocsv.MarshalString(a.selects); err != nil {
		logger.Error("Failed to write feedback csv. filePath: %v, err: %v", filePath, err)
	}
	a.feedbacks = []*Feedback{}
}

func (a *ActionEvent) Append(logs []EventLog) {
	for _, log := range logs {
		cps, err := decryptCps(log.Cps)
		if err != nil {
			logger.Error("Failed to decrypt cps. cps: %v, err: %v", cps, err)
			continue
		}
		createdAt, err := time.Parse("2006-01-02 15:04:05", log.CreatedAt)
		if err != nil {
			logger.Error("Failed to parse created_at.err: %v", err)
			continue
		}
		if log.Type == "session" {
			a.sessions = append(
				a.sessions,
				&Session{
					Uuid:      cps.Uuid,
					Lng:       cps.Lng,
					Lat:       cps.Lat,
					ZoomLevel: cps.ZoomLevel,
					CreatedAt: createdAt,
				},
			)
		} else if log.Type == "select" {
			a.selects = append(
				a.selects,
				&Select{
					Uuid:      cps.Uuid,
					HotelID:   cps.HotelId,
					Lng:       cps.Lng,
					Lat:       cps.Lat,
					ZoomLevel: cps.ZoomLevel,
					CreatedAt: createdAt,
				},
			)
		} else if log.Type == "feedback" {
			a.feedbacks = append(
				a.feedbacks,
				&Feedback{
					Uuid:      cps.Uuid,
					HotelID:   cps.HotelId,
					Action:    log.Action,
					ZoomLevel: cps.ZoomLevel,
					CreatedAt: createdAt,
				},
			)
		}
	}
}

// uuid:hotel_id:lng:lat:zoom_level
func decryptCps(cps string) (Cps, error) {
	decrypedText, err := util.Decrypt(cps)
	if err != nil {
		return Cps{}, err
	}
	rep := regexp.MustCompile(`\s*:\s*`)
	res := rep.Split(decrypedText, -1)
	lng, err := strconv.ParseFloat(res[2], 64)
	if err != nil {
		return Cps{}, err
	}
	lat, err := strconv.ParseFloat(res[3], 64)
	if err != nil {
		return Cps{}, err
	}
	zoomLevel, err := strconv.Atoi(res[4])
	if err != nil {
		return Cps{}, err
	}
	return Cps{
		Uuid:      res[0],
		HotelId:   res[1],
		Lng:       lng,
		Lat:       lat,
		ZoomLevel: zoomLevel,
	}, nil
}
