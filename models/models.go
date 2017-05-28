package models

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/BeauceronSecurity/gophish/config"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3" // Blank import needed to import sqlite3
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var db *gorm.DB
var err error

// ErrUsernameTaken is thrown when a user attempts to register a username that is taken.
var ErrUsernameTaken = errors.New("username already taken")

// Logger is a global logger used to show informational, warning, and error messages
var Logger = log.New(os.Stdout, " ", log.Ldate|log.Ltime|log.Lshortfile)

const (
	CAMPAIGN_IN_PROGRESS string = "In progress"
	CAMPAIGN_QUEUED      string = "Queued"
	CAMPAIGN_CREATED     string = "Created"
	CAMPAIGN_EMAILS_SENT string = "Emails Sent"
	CAMPAIGN_COMPLETE    string = "Completed"
	EVENT_SENT           string = "Email Sent"
	EVENT_SENDING_ERROR  string = "Error Sending Email"
	EVENT_OPENED         string = "Email Opened"
	EVENT_CLICKED        string = "Clicked Link"
	EVENT_DATA_SUBMIT    string = "Submitted Data"
	STATUS_SUCCESS       string = "Success"
	STATUS_SENDING       string = "Sending"
	STATUS_UNKNOWN       string = "Unknown"
	ERROR                string = "Error"
)

// Flash is used to hold flash information for use in templates.
type Flash struct {
	Type    string
	Message string
}

// Response contains the attributes found in an API response
type Response struct {
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

// Copy of auth.GenerateSecureKey to prevent cyclic import with auth library
func generateSecureKey() string {
	k := make([]byte, 32)
	io.ReadFull(rand.Reader, k)
	return fmt.Sprintf("%x", k)
}

// Setup initializes the Conn object
// It also populates the Gophish Config object
func Setup() error {
	// Open our database connection
	db, err = gorm.Open(config.Conf.DBName, config.Conf.DBPath)
	db.LogMode(false)
	db.SetLogger(Logger)
	db.DB().SetMaxOpenConns(1)
	if err != nil {
		Logger.Println(err)
		return err
	}
	return nil
}
