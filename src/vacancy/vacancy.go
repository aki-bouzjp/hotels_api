package vacancy

import (
	"app/src/db"
	"app/src/logger"
	"app/src/model"
	"strings"

	"net/http"
	"time"
)

type VacancyStatus struct {
	HotelID   string    `json:"hotel_id"`
	Empty     bool      `json:"empty"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func SearchVacancies(ids []string) {
	url := "http://127.0.0.1:8081?" + strings.Join(ids, "&ids[]=")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Error("Failed to request search_vacancies. err: %v", err)
		return
	}
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		logger.Error("Failed to request search_vacancies. err: %v", err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		logger.Error("Failed something to request. statud code: %d", res.StatusCode)
		return
	}
}

func GetVacancies(ids []string, db *db.DB) ([]*VacancyStatus, error) {
	vs := []*VacancyStatus{}
	searchIds := []string{}
	now := time.Now()

	for _, id := range ids {
		var v model.Vacancy
		db.Conn.First(&v, id)
		if v.ID == "" {
			searchIds = append(searchIds, id)
			vs = append(vs, &VacancyStatus{
				HotelID:   id,
				Empty:     v.Empty,
				Status:    "getting",
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			})
			continue
		}

		diff := now.Sub(v.UpdatedAt)
		if diff.Seconds() < 10*60 { // 10分以内に更新されたレコード
			vs = append(vs, &VacancyStatus{
				HotelID:   v.HotelID,
				Empty:     v.Empty,
				Status:    "updated",
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			})
		} else {
			searchIds = append(searchIds, id)
			vs = append(vs, &VacancyStatus{
				HotelID:   v.HotelID,
				Empty:     v.Empty,
				Status:    "getting",
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			})
		}
	}
	SearchVacancies(searchIds)
	return vs, nil
}
