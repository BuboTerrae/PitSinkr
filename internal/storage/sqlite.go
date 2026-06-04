package storage

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type DNSLog struct {
	ID        int64     `json:"id"`
	ClientIP  string    `json:"client_ip"`
	Domain    string    `json:"domain"`
	Type      string    `json:"query_type"`
	Timestamp time.Time `json:"timestamp"`
}

func OpenDB() (*sql.DB, error) {
	return sql.Open("sqlite3", "./data.db")
}

func InitDB(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS dns_logs(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	client_ip TEXT NOT NULL,
	domain TEXT NOT NULL,
	query_type TEXT NOT NULL,
	timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
	)`
	// query:=`DROP TABLE dns_logs`

	_, err := db.Exec(query)

	return err
}

func InsertLog(db *sql.DB, log DNSLog) error {
	_, err := db.Exec(`INSERT INTO dns_logs(client_ip, domain, query_type) VALUES (?, ?, ?)`, log.ClientIP, log.Domain, log.Type)

	return err
}

func FetchLogs(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query(`SELECT id, client_ip, domain, query_type, timestamp FROM dns_logs`)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		}

		defer rows.Close()

		logs := make([]DNSLog, 0)

		for rows.Next() {

			var logEntry DNSLog

			err := rows.Scan(
				&logEntry.ID,
				&logEntry.ClientIP,
				&logEntry.Domain,
				&logEntry.Type,
				&logEntry.Timestamp,
			)

			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

			logs = append(logs, logEntry)
		}

		c.JSON(200, logs)
	}
}
