package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"repo/internal/application/dtos"
	"repo/internal/application/services"
	"strconv"
)

type ReportHandler struct {
	reportService *services.ReportService
}

func NewReportHandler(reportService *services.ReportService) *ReportHandler {
	return &ReportHandler{reportService: reportService}
}

func (h *ReportHandler) GetAllReports(w http.ResponseWriter, r *http.Request) {
	reports, err := h.reportService.GetAllReports()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(reports)
}

func (h *ReportHandler) GetReport(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid report ID", http.StatusBadRequest)
		return
	}

	report, err := h.reportService.GetReportByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(report)
}

func (h *ReportHandler) CreateReport(w http.ResponseWriter, r *http.Request) {
	var reportDTO dtos.ReportDTO
	if err := json.NewDecoder(r.Body).Decode(&reportDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdReport, err := h.reportService.CreateReport(&reportDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdReport)
}

func (h *ReportHandler) UpdateReport(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid report ID", http.StatusBadRequest)
		return
	}

	var reportDTO dtos.ReportDTO
	if err := json.NewDecoder(r.Body).Decode(&reportDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	reportDTO.ID = id

	updatedReport, err := h.reportService.UpdateReport(&reportDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updatedReport)
}

func (h *ReportHandler) DeleteReport(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid report ID", http.StatusBadRequest)
		return
	}

	err = h.reportService.DeleteReport(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
