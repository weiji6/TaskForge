package models

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

type Room struct {
    ID      int    `json:"id"`
    Name    string `json:"name"`
    Users   []User `json:"users"`
}

type Message struct {
    ID      int    `json:"id"`
    Content string `json:"content"`
    Sender  User   `json:"sender"`
    Room    Room   `json:"room"`
    Time    string `json:"time"` // in UTC format
}

type Notification struct {
    ID      int    `json:"id"`
    Message  Message `json:"message"`
    User     User    `json:"user"`
    Read     bool    `json:"read"`
    TimeSent string  `json:"time_sent"` // in UTC format
}