package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"example.com/shared/api/core"
)

// UploadImage godoc
// @Summary      Upload product image
// @Description  Upload an image for a product
// @Tags         products
// @Accept       multipart/form-data
// @Produce      json
// @Param        productId  path      string  true  "Product ID"
// @Param        image      formData file    true  "Image file"
// @Success      204  "No Content"
// @Failure      400  {object}  core.ErrorMsg
// @Failure      404  {object}  core.ErrorMsg
// @Failure      500  {object}  core.ErrorMsg
// @Router       /products/{productId}/img [post]
func UploadImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	productID := r.PathValue("productId")

	if productID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(core.ErrorMsg{Msg: "missing product id"})
		return
	}

	err := r.ParseMultipartForm(20 << 20) // 20MB
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(core.ErrorMsg{Msg: "invalid form"})
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(core.ErrorMsg{Msg: "error retrieving file"})
		return
	}
	defer file.Close()

	// Proper Content-Type detection
	fileBuffer := make([]byte, 512)
	_, err = file.Read(fileBuffer)
	if err != nil {
		msg := "failed to read uploaded file for content-type detection"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(core.ErrorMsg{Msg: msg})
		return

	}
	fileType := http.DetectContentType(fileBuffer)

	// Reset file reader position
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		msg := "failed to reset file pointer after reading"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(core.ErrorMsg{Msg: msg})
		return
	}

	allowedTypes := map[string]bool{
		"image/jpeg":    true,
		"image/png":     true,
		"image/gif":     true,
		"image/webp":    true,
		"image/avif":    true,
		"image/svg+xml": true,
	}

	if !allowedTypes[fileType] {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(core.ErrorMsg{Msg: "invalid image format"})
		return
	}

	fmt.Printf("FileName: %s\n", header.Filename)
	fmt.Printf("ProductID: %s\n", productID)
	fmt.Printf("FileType: %s\n", fileType)

	w.WriteHeader(http.StatusNoContent)
}
