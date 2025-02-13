package main

import (
	"log"
	"regexp"
)

func validatePassword(password string) (isValid bool, msg string) {
	if len([]rune(password)) < 8 {
		return false, "Password minimal 8 karakter ya! 💪"
	}
	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return false, "Harus ada huruf besar nih! 🔠"
	}
	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return false, "Jangan lupa huruf kecil juga ya! 🔡"
	}
	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		return false, "Angka juga wajib ada! 🔢"
	}
	if !regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString(password) {
		return false, "Tambahkan simbol biar lebih kuat! 🔣"
	}
	return true, "Password aman! Mantap! ✅"
}

func main() {
	password := "B@n9L0k3rD0n9B@n9!2025"
	isValid, message := validatePassword(password)
	log.Printf("valid = %t, message = %s", isValid, message)
}
