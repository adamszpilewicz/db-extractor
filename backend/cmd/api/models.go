package main

import (
	"database/sql"
	"time"
)

// Credentials represents the database connection credentials received from frontend.
type Credentials struct {
	Email    string `json:"email"`
	Host     string `json:"host"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type SchemaInfo struct {
	SchemaName string   `json:"schema_name"`
	TableNames []string `json:"table_names"`
}

type ColumnInfo struct {
	ColumnName string `json:"column_name"`
	ColumnType string `json:"column_type"`
	PrimaryKey string `json:"primary_key"`
}

type TableRow map[string]interface{}

type AppView struct {
	ViewName   string `json:"view_name"`
	Owner      string `json:"owner"`
	Definition string `json:"definition"`
}

type AppInfo struct {
	PostgresVersion sql.NullString `json:"postgres_version"`
	CurrentDatabase sql.NullString `json:"current_database"`
	CurrentUser     sql.NullString `json:"current_user"`
	ServerIP        sql.NullString `json:"server_ip"`
	ServerPort      int            `json:"server_port"`
	StartTime       time.Time      `json:"start_time"`
	LastConfigLoad  time.Time      `json:"last_config_load"`
	InRecovery      bool           `json:"in_recovery"`
	//LastWalReceived sql.NullString // using NullString to handle potential NULL values
	//LastWalReplayed sql.NullString // using NullString to handle potential NULL values
	//CurrentWal      string
	//CurrentDbSize   int64 // assuming this is big enough for your DB sizes
}

// DBStatistics define a struct to hold the fetched statistics
type DBStatistics struct {
	DatabaseName                 string
	ActiveConnections            int
	TotalTransactionsCommitted   int
	TotalTransactionsRolledBack  int
	BlocksRead                   int
	CacheHits                    int
	RowsReturnedByQueries        int
	RowsFetchedByQueries         int
	RowsInserted                 int
	RowsUpdated                  int
	RowsDeleted                  int
	ConflictsDetected            int
	ActiveSessions               int
	IdleSessions                 int
	IdleInTransactionSessions    int
	FastpathFunctionCallSessions int
}
