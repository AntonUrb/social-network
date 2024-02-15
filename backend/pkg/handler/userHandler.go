package handler

import (
	"backend/pkg/model"
	"backend/pkg/repository"
	"backend/util"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	userRepo *repository.UserRepository
	sessionRepo *repository.SessionRepository
}

func NewUserHandler(uRepo *repository.UserRepository, sRepo *repository.SessionRepository) *UserHandler {
	return &UserHandler{userRepo: uRepo, sessionRepo: sRepo}
}

func (h *UserHandler) UserRegisterHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Parse the multipart form data from the request
	err1 := r.ParseMultipartForm(10 << 20) // Maximum memory 10MB, change this based on your requirements
	if err1 != nil {
		http.Error(w, "Error parsing form data: "+err1.Error(), http.StatusBadRequest)
		return
	}
	var regData model.RegistrationData // variable based on struct fields for storing registration data

	regData.Username = r.FormValue("username")
	regData.Email = r.FormValue("email")
	regData.Password = r.FormValue("password")
	regData.FirstName = r.FormValue("first_name")
	regData.LastName = r.FormValue("last_name")
	regData.DOB = r.FormValue("dob")
	regData.About = r.FormValue("about")
	

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(regData.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	// Change input password data to hashed variant
	regData.Password = string(hashedPassword)

	util.ImageSave(w, r, &regData) // parses image data from request to the variable
	// Store user in database
	userID, err := h.userRepo.RegisterUser(regData)
	if err != nil {
		http.Error(w, "Error registering user: "+ err.Error(), http.StatusInternalServerError)
		return
	}

	// Generate a session token and store it in database with expiration time
	sessionToken := util.GenerateSessionToken()
	h.sessionRepo.StoreSessionInDB(sessionToken, int(userID))

	// Set a cookie with the session token
	http.SetCookie(w, &http.Cookie{
		Name: "session_token",
		Value: sessionToken,
		MaxAge: 60*15, // 15 minutes
		Path:     "/", // Make cookie available for all paths
	})

	// Send a success response
	response := map[string]interface{}{
		"message": "User registration successful",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) GetUserProfileByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the URL
	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Invalid user ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Get the user profile from the database
	profile, err := h.userRepo.GetUserProfileByID(userID)
	if err != nil {
		http.Error(w, "Error getting user profile: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}