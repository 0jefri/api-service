package report

import (
	"fmt"
	"log"
	"time"

	ordershiping "github.com/api-service/api/order_shiping"
	"github.com/xuri/excelize/v2"
)

type ReportService interface {
	GenerateOrderReport() error
}

type reportService struct {
	orderRepo ordershiping.OrderShippingRepository
}

func NewReportService(orderRepo ordershiping.OrderShippingRepository) ReportService {
	return &reportService{orderRepo: orderRepo}
}

func (s *reportService) GenerateOrderReport() error {
	// Fetch data dari repository
	orders, err := s.orderRepo.FetchOrders()
	log.Println("data:", orders)
	if err != nil {
		return fmt.Errorf("gagal mengambil data: %w", err)
	}

	// Buat file Excel
	file := excelize.NewFile()
	sheetName := "Orders"
	index, _ := file.NewSheet(sheetName)

	// Header
	file.SetCellValue(sheetName, "A1", "OrderID")
	file.SetCellValue(sheetName, "B1", "EcommerceID")
	file.SetCellValue(sheetName, "C1", "ShippingID")
	file.SetCellValue(sheetName, "D1", "OriginLongitude")
	file.SetCellValue(sheetName, "E1", "DestinationLatitude")
	file.SetCellValue(sheetName, "F1", "DestinationLongitude")
	file.SetCellValue(sheetName, "G1", "TotalPaymentShipping")

	// Isi data
	for i, order := range orders {
		row := i + 2
		file.SetCellValue(sheetName, fmt.Sprintf("A%d", row), order.OrderID)
		file.SetCellValue(sheetName, fmt.Sprintf("B%d", row), order.EcommerceID)
		file.SetCellValue(sheetName, fmt.Sprintf("C%d", row), order.ShippingID)
		file.SetCellValue(sheetName, fmt.Sprintf("D%d", row), order.OriginLongitude)
		file.SetCellValue(sheetName, fmt.Sprintf("E%d", row), order.DestinationLatitude)
		file.SetCellValue(sheetName, fmt.Sprintf("F%d", row), order.DestinationLongitude)
		file.SetCellValue(sheetName, fmt.Sprintf("G%d", row), order.TotalPaymentShipping)
	}

	file.SetActiveSheet(index)
	// Simpan file
	fileName := fmt.Sprintf("reports/order_report_%s.xlsx", time.Now().Format("20060102_150405"))
	if err := file.SaveAs(fileName); err != nil {
		return fmt.Errorf("gagal menyimpan laporan: %w", err)
	}

	fmt.Printf("Laporan berhasil dibuat: %s\n", fileName)
	return nil
}
