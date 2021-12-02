package responses

import (
	"net/http"
)

// Fungsi untuk memberikan respon ketika controller gagal dijalankan
func StatusFailed(message string) map[string]interface{} {
	var result = map[string]interface{}{
		"status":  "failed",
		"message": message,
	}
	return result
}

// Fungsi untuk memberikan respon ketika Database gagal dijalankan
func StatusInternalServerError() map[string]interface{} {
	var result = map[string]interface{}{
		"status":  "failed",
		"message": "internal server error",
	}
	return result
}

// Fungsi untuk memberikan respon ketika controller service error dijalankan
func StatusFailedInternal(message string, data interface{}) map[string]interface{} {
	var result = map[string]interface{}{
		"status":  "Unauthorized failed",
		"message": message,
		"data":    data,
	}
	return result
}

func StatusUnauthorized() map[string]interface{} {
	var result = map[string]interface{}{
		"status":  "Unauthorized",
		"message": "Unauthorized Access",
	}
	return result
}

// Fungsi untuk memberikan respon ketika controller berhasil dijalankan
func StatusSuccess(message string) map[string]interface{} {
	var result = map[string]interface{}{
		"status":  "success",
		"message": message,
	}
	return result
}

// Fungsi untuk memberikan respon ketika controller berhasil dijalankan dan menerima masukan data
func StatusSuccessData(message string, data interface{}) map[string]interface{} {
	var result = map[string]interface{}{
		"status":  "success",
		"message": message,
		"data":    data,
	}
	return result
}

// function response false param
func FalseParamResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "False Param",
	}
	return result
}

func StatusSuccessLogin(message string, id, token, name interface{}) map[string]interface{} {
	var result = map[string]interface{}{
		"status":  "success",
		"message": message,
		"id":      id,
		"token":   token,
		"name":    name,
	}
	return result
}
