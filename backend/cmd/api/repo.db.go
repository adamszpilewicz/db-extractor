package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DatabaseRepo struct {
	DB  *pgxpool.Pool
	app *application
}

func NewDatabaseRepo(app *application, creds Credentials) *DatabaseRepo {
	// Set app configuration fields
	app.config.username = creds.User
	app.config.password = creds.Password
	app.config.database = creds.Database
	app.config.host = creds.Host

	return &DatabaseRepo{app: app}
}

func (dr *DatabaseRepo) Connect() error {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		dr.app.config.host, dr.app.config.username, dr.app.config.password, dr.app.config.database)

	var err error
	dr.DB, err = pgxpool.Connect(context.Background(), connectionString)

	if err != nil {
		return err
	}

	return nil
}

func (dr *DatabaseRepo) CheckConnection() error {
	conn, err := dr.DB.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	return nil
}

func (repo *DatabaseRepo) DiscoverSchemas() ([]SchemaInfo, error) {
	rows, err := repo.DB.Query(context.Background(), querySchemas)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Use a map to collect tables by schema
	schemaMap := make(map[string][]string)
	var schemas []SchemaInfo

	for rows.Next() {
		var schemaName string
		var tableName string

		err := rows.Scan(&schemaName, &tableName)
		if err != nil {
			return nil, err
		}

		schemaMap[schemaName] = append(schemaMap[schemaName], tableName)
	}

	// Convert the map into a slice
	for schema, tables := range schemaMap {
		schemas = append(schemas, SchemaInfo{
			SchemaName: schema,
			TableNames: tables,
		})
	}

	return schemas, nil
}

func (repo *DatabaseRepo) DiscoverTableColumns(schemaName, tableName string) ([]ColumnInfo, error) {
	rows, err := repo.DB.Query(context.Background(), queryTableDetails, tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var columns []ColumnInfo
	for rows.Next() {
		var columnInfo ColumnInfo
		if err := rows.Scan(&columnInfo.ColumnName, &columnInfo.ColumnType, &columnInfo.PrimaryKey); err != nil {
			return nil, err
		}
		columns = append(columns, columnInfo)
	}
	return columns, nil
}

func (repo *DatabaseRepo) FetchTableData(schemaName, tableName string) ([]TableRow, error) {
	// Construct the query to fetch all rows from the given table
	query := fmt.Sprintf(queryAll, schemaName, tableName)

	rows, err := repo.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ProcessRows(rows)
}

// ExecuteQuery executes the given SQL query and returns the results
func (repo *DatabaseRepo) ExecuteQuery(query string) ([]TableRow, error) {
	// Execute the provided query
	rows, err := repo.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ProcessRows(rows)
}

func (repo *DatabaseRepo) FetchReplicationSlotsStats() ([]TableRow, error) {
	rows, err := repo.DB.Query(context.Background(), queryReplicationSlots)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ProcessRows(rows)
}

func (repo *DatabaseRepo) FetchProcedures() ([]TableRow, error) {
	rows, err := repo.DB.Query(context.Background(), queryStoredProcedures)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ProcessRows(rows)
}

func (repo *DatabaseRepo) FetchAppViews() ([]AppView, error) {
	rows, err := repo.DB.Query(context.Background(), queryViews)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var views []AppView

	for rows.Next() {
		var v AppView
		if err := rows.Scan(&v.ViewName, &v.Owner, &v.Definition); err != nil {
			return nil, err
		}
		views = append(views, v)
	}

	return views, nil
}

func (dr *DatabaseRepo) GetAppInfo() (AppInfo, error) {
	var info AppInfo

	ctx := context.Background()
	row := dr.DB.QueryRow(ctx, queryInfo)
	if err := row.Scan(&info.PostgresVersion, &info.CurrentDatabase, &info.CurrentUser, &info.ServerIP, &info.ServerPort, &info.StartTime, &info.LastConfigLoad, &info.InRecovery); //&info.LastWalReceived, &info.LastWalReplayed, &info.CurrentWal, &info.CurrentDbSize
	err != nil {
		return AppInfo{}, err
	}

	return info, nil
}

func (dr *DatabaseRepo) FetchDBStatistics() (DBStatistics, error) {
	row := dr.DB.QueryRow(context.Background(), queryStatistics)

	var stats DBStatistics

	err := row.Scan(
		&stats.DatabaseName,
		&stats.ActiveConnections,
		&stats.TotalTransactionsCommitted,
		&stats.TotalTransactionsRolledBack,
		&stats.BlocksRead,
		&stats.CacheHits,
		&stats.RowsReturnedByQueries,
		&stats.RowsFetchedByQueries,
		&stats.RowsInserted,
		&stats.RowsUpdated,
		&stats.RowsDeleted,
		&stats.ConflictsDetected,
		&stats.ActiveSessions,
		&stats.IdleSessions,
		&stats.IdleInTransactionSessions,
		&stats.FastpathFunctionCallSessions,
	)
	if err != nil {
		return DBStatistics{}, err
	}

	return stats, nil
}

func (repo *DatabaseRepo) FetchPgSettings() ([]TableRow, error) {
	rows, err := repo.DB.Query(context.Background(), queryPgSettings)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ProcessRows(rows)
}
