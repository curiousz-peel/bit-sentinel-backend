package main

import "time"

type Author struct {
	ID          string   `json:"id"`
	FirstName   string   `json:"firstName"`
	LastName    string   `json:"lastName"`
	Profession  string   `json:"profession"`
	Description string   `json:"description"`
	Topics      []string `json:"topics"`
}

type Course struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author Author `json:"author"`
	//need to decide on the Content type..
	Content   string    `json:"content"`
	Tags      []string  `json:"tags"`
	AddedDate time.Time `json:"addedDate"`
	Visible   bool      `json:"visible"`
}

type Subscription struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	ExpiresAt time.Time `json:"expiresAt"`
	Price     int       `json:"price"`
}

type User struct {
	ID       string    `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Birthday time.Time `json:"birthday"`
	IsMod    bool      `json:"isModerator"`
}

type Account struct {
	ID               string       `json:"id"`
	User             User         `json:"user"`
	Subscription     Subscription `json:"subscription"`
	Courses          []Course     `json:"courses"`
	RegistrationDate time.Time    `json:"registrationDate"`
}
