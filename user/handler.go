package user

import (
	"log"
	"net/http"
)

type Handler struct{}

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("listing all users")
	w.Write([]byte("list of users"))
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	log.Println("creating a user")
	w.Write([]byte("user created!"))
}

func (h *Handler) FindUserByID(w http.ResponseWriter, r *http.Request) {
	log.Println("Searching for a user")
	userId := r.PathValue("id")
	w.Write([]byte("Searching for user with ID:, " + userId + "!"))
}

func (h *Handler) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	log.Println("Updating a user")
	userId := r.PathValue("id")
	w.Write([]byte("Updating user with ID:, " + userId + "!"))
}

func (h *Handler) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	log.Println("Deleting a user")
	userId := r.PathValue("id")
	w.Write([]byte("Deleting user with ID:, " + userId + "!"))
}
