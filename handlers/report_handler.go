package handlers

import (
	"encoding/json"
	"kasir-api/services"
	"net/http"
	"time"
)

type ReportHandler struct {
	service *services.ReportService
}

func NewReportHandler(service *services.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (h *ReportHandler) HandleReport(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GenerateReport(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ReportHandler) GenerateReport(w http.ResponseWriter, r *http.Request) {
	dateFromString := r.URL.Query().Get("date_from")
	dateToString := r.URL.Query().Get("date_to")

	var dateFrom, dateTo time.Time
	var err error
	if dateFromString == "" || dateToString == "" {
		now := time.Now()
		dateFrom = time.Date(now.Year(), now.Month(), 0, 0, 0, 0, 0, time.Local)
		dateTo = time.Date(now.Year(), now.Month(), 23, 59, 59, 0, 0, time.Local)
	} else {
		dateFrom, err = time.Parse("2006-01-02", dateFromString)
		if err != nil {
			http.Error(w, "invalid date format, use yyyy-mm-dd", http.StatusBadRequest)
		}
		dateTo, err = time.Parse("2006-01-02", dateToString)
		if err != nil {
			http.Error(w, "invalid date format, use yyyy-mm-dd", http.StatusBadRequest)
		}
		dateTo = time.Date(dateTo.Year(), dateTo.Month(), dateTo.Day(), 23, 59, 59, 0, time.Local)
	}

	report, err := h.service.GenerateReport(dateFrom, dateTo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)

}
