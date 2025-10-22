package router

import (
	"fmt"
	"golang-elk/controller"
	"net/http"

)

// Hàm khởi tạo router
func InitRouter() *http.ServeMux {
	router := http.NewServeMux()

	// Đăng ký route theo pattern mới của Go 1.22

	// Trang mặc định
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hệ thống thu thập log đang chạy!")
	})
	router.HandleFunc("GET /api/base", controller.Base)
	return router
}
