package report

import (
	"net/http"
)

func TriggerReportHandler(reportService ReportService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		err := reportService.GenerateOrderReport()
		if err != nil {
			http.Error(w, "Gagal membuat laporan", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Laporan berhasil dibuat"))
	}
}
