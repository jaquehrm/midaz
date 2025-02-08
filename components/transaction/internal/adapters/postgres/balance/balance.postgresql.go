package balance

import (
	"context"
	"database/sql"
	"errors"
	"github.com/LerianStudio/midaz/pkg"
	"github.com/LerianStudio/midaz/pkg/constant"
	"github.com/LerianStudio/midaz/pkg/mmodel"
	"github.com/LerianStudio/midaz/pkg/mopentelemetry"
	"github.com/LerianStudio/midaz/pkg/mpostgres"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// Repository provides an interface for operations related to balance template entities.
//
//go:generate mockgen --destination=balance.mock.go --package=balance . Repository
type Repository interface {
	Create(ctx context.Context, balance *mmodel.Balance) error
	ListByAccountIDs(ctx context.Context, organizationID, ledgerID uuid.UUID, ids []uuid.UUID) ([]*mmodel.Balance, error)
	ListByAliases(ctx context.Context, organizationID, ledgerID uuid.UUID, aliases []string) ([]*mmodel.Balance, error)
	SelectForUpdate(ctx context.Context, organizationID, ledgerID uuid.UUID, balances []*mmodel.Balance) error
}

// BalancePostgreSQLRepository is a Postgresql-specific implementation of the BalanceRepository.
type BalancePostgreSQLRepository struct {
	connection *mpostgres.PostgresConnection
	tableName  string
}

// NewBalancePostgreSQLRepository returns a new instance of BalancePostgreSQLRepository using the given Postgres connection.
func NewBalancePostgreSQLRepository(pc *mpostgres.PostgresConnection) *BalancePostgreSQLRepository {
	c := &BalancePostgreSQLRepository{
		connection: pc,
		tableName:  "balance",
	}

	_, err := c.connection.GetDB()
	if err != nil {
		panic("Failed to connect database")
	}

	return c
}

func (r *BalancePostgreSQLRepository) Create(ctx context.Context, balance *mmodel.Balance) error {
	tracer := pkg.NewTracerFromContext(ctx)

	ctx, span := tracer.Start(ctx, "postgres.create_balances")
	defer span.End()
	db, err := r.connection.GetDB()
	if err != nil {
		mopentelemetry.HandleSpanError(&span, "Failed to get database connection", err)
	}

	record := &BalancePostgreSQLModel{}
	record.FromEntity(balance)

	ctx, spanExec := tracer.Start(ctx, "postgres.create.exec")

	err = mopentelemetry.SetSpanAttributesFromStruct(&spanExec, "balance_repository_input", record)
	if err != nil {
		mopentelemetry.HandleSpanError(&spanExec, "Failed to convert balance record from entity to JSON string", err)

		return err
	}

	result, err := db.ExecContext(ctx, `INSERT INTO balance VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) RETURNING *`,
		record.ID,
		record.OrganizationID,
		record.LedgerID,
		record.AccountID,
		record.Alias,
		record.AssetCode,
		record.Available,
		record.OnHold,
		record.Scale,
		record.Version,
		record.AccountType,
		record.AllowSending,
		record.AllowReceiving,
		record.CreatedAt,
		record.UpdatedAt,
		record.DeletedAt,
	)
	if err != nil {
		mopentelemetry.HandleSpanError(&spanExec, "Failed to execute query", err)

		return err
	}

	spanExec.End()

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		mopentelemetry.HandleSpanError(&span, "Failed to get rows affected", err)

		return err
	}

	if rowsAffected == 0 {
		err := pkg.ValidateBusinessError(constant.ErrEntityNotFound, reflect.TypeOf(mmodel.Balance{}).Name())

		mopentelemetry.HandleSpanError(&span, "Failed to create balance. Rows affected is 0", err)

		return err
	}

	return nil
}

// ListByAccountIDs list Balances entity from the database using the provided accountIDs.
func (r *BalancePostgreSQLRepository) ListByAccountIDs(ctx context.Context, organizationID, ledgerID uuid.UUID, accountIds []uuid.UUID) ([]*mmodel.Balance, error) {
	tracer := pkg.NewTracerFromContext(ctx)

	ctx, span := tracer.Start(ctx, "postgres.list_balances_by_ids")
	defer span.End()

	db, err := r.connection.GetDB()
	if err != nil {
		mopentelemetry.HandleSpanError(&span, "Failed to get database connection", err)

		return nil, err
	}

	var balances []*mmodel.Balance

	ctx, spanQuery := tracer.Start(ctx, "postgres.list_by_ids.query")

	rows, err := db.QueryContext(
		ctx,
		"SELECT * FROM balance WHERE organization_id = $1 AND ledger_id = $2 AND account_id = ANY($3) AND deleted_at IS NULL ORDER BY created_at DESC",
		organizationID,
		ledgerID,
		pq.Array(accountIds),
	)
	if err != nil {
		mopentelemetry.HandleSpanError(&spanQuery, "Failed to execute query", err)

		return nil, err
	}
	defer rows.Close()

	spanQuery.End()

	for rows.Next() {
		var balance BalancePostgreSQLModel
		if err := rows.Scan(
			&balance.ID,
			&balance.OrganizationID,
			&balance.LedgerID,
			&balance.AccountID,
			&balance.Alias,
			&balance.AssetCode,
			&balance.Available,
			&balance.OnHold,
			&balance.Scale,
			&balance.Version,
			&balance.AccountType,
			&balance.AllowSending,
			&balance.AllowReceiving,
			&balance.CreatedAt,
			&balance.UpdatedAt,
			&balance.DeletedAt,
		); err != nil {
			mopentelemetry.HandleSpanError(&span, "Failed to scan row", err)

			return nil, err
		}

		balances = append(balances, balance.ToEntity())
	}

	if err := rows.Err(); err != nil {
		mopentelemetry.HandleSpanError(&span, "Failed to iterate rows", err)

		return nil, err
	}

	return balances, nil
}

// ListByAliases list Balances entity from the database using the provided aliases.
func (r *BalancePostgreSQLRepository) ListByAliases(ctx context.Context, organizationID, ledgerID uuid.UUID, aliases []string) ([]*mmodel.Balance, error) {
	tracer := pkg.NewTracerFromContext(ctx)

	ctx, span := tracer.Start(ctx, "postgres.list_balances_by_aliases")
	defer span.End()

	db, err := r.connection.GetDB()
	if err != nil {
		mopentelemetry.HandleSpanError(&span, "Failed to get database connection", err)

		return nil, err
	}

	var balances []*mmodel.Balance

	ctx, spanQuery := tracer.Start(ctx, "postgres.list_by_aliases.query")

	rows, err := db.QueryContext(
		ctx,
		"SELECT * FROM balance WHERE organization_id = $1 AND ledger_id = $2 AND alias = ANY($3) AND deleted_at IS NULL ORDER BY created_at DESC",
		organizationID,
		ledgerID,
		pq.Array(aliases),
	)
	if err != nil {
		mopentelemetry.HandleSpanError(&spanQuery, "Failed to execute query", err)

		return nil, err
	}
	defer rows.Close()

	spanQuery.End()

	for rows.Next() {
		var balance BalancePostgreSQLModel
		if err := rows.Scan(
			&balance.ID,
			&balance.OrganizationID,
			&balance.LedgerID,
			&balance.AccountID,
			&balance.Alias,
			&balance.AssetCode,
			&balance.Available,
			&balance.OnHold,
			&balance.Scale,
			&balance.Version,
			&balance.AccountType,
			&balance.AllowSending,
			&balance.AllowReceiving,
			&balance.CreatedAt,
			&balance.UpdatedAt,
			&balance.DeletedAt,
		); err != nil {
			mopentelemetry.HandleSpanError(&span, "Failed to scan row", err)

			return nil, err
		}

		balances = append(balances, balance.ToEntity())
	}

	if err := rows.Err(); err != nil {
		mopentelemetry.HandleSpanError(&span, "Failed to iterate rows", err)

		return nil, err
	}

	return balances, nil
}

// SelectForUpdate a Balance entity into Postgresql.
func (r *BalancePostgreSQLRepository) SelectForUpdate(ctx context.Context, organizationID, ledgerID uuid.UUID, balances []*mmodel.Balance) error {
	tracer := pkg.NewTracerFromContext(ctx)
	logger := pkg.NewLoggerFromContext(ctx)

	ctx, span := tracer.Start(ctx, "postgres.update_balances")
	defer span.End()

	db, err := r.connection.GetDB()
	if err != nil {
		mopentelemetry.HandleSpanError(&span, "Failed to get database connection", err)

		return err
	}

	tx, err := db.Begin()
	if err != nil {
		mopentelemetry.HandleSpanError(&span, "Failed to init balances", err)

		return err
	}

	for _, blc := range balances {

		query := "SELECT * FROM balance WHERE organization_id = $1 AND ledger_id = $2 AND id = $3 AND deleted_at IS NULL FOR UPDATE"

		row := tx.QueryRowContext(ctx, query, organizationID, ledgerID, blc.ID)

		var balance BalancePostgreSQLModel
		err = row.Scan(
			&balance.ID,
			&balance.OrganizationID,
			&balance.LedgerID,
			&balance.AccountID,
			&balance.Alias,
			&balance.AssetCode,
			&balance.Available,
			&balance.OnHold,
			&balance.Scale,
			&balance.Version,
			&balance.AccountType,
			&balance.AllowSending,
			&balance.AllowReceiving,
			&balance.CreatedAt,
			&balance.UpdatedAt,
			&balance.DeletedAt,
		)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				logger.Errorf("registro não encontrado para ID %s", blc.ID)

				return err
			}

			logger.Errorf("erro no select for update: %v", err)

			return err
		}

		var updates []string

		var args []any

		updates = append(updates, "available = $"+strconv.Itoa(len(args)+1))
		args = append(args, blc.Available)

		updates = append(updates, "on_hold = $"+strconv.Itoa(len(args)+1))
		args = append(args, blc.OnHold)

		updates = append(updates, "scale = $"+strconv.Itoa(len(args)+1))
		args = append(args, blc.Scale)

		updates = append(updates, "version = $"+strconv.Itoa(len(args)+1))
		version := blc.Version + 1
		args = append(args, version)

		updates = append(updates, "updated_at = $"+strconv.Itoa(len(args)+1))
		args = append(args, time.Now(), organizationID, ledgerID, blc.ID)

		queryUpdate := `UPDATE balance SET ` + strings.Join(updates, ", ") +
			` WHERE organization_id = $` + strconv.Itoa(len(args)-2) +
			` AND ledger_id = $` + strconv.Itoa(len(args)-1) +
			` AND id = $` + strconv.Itoa(len(args)) +
			` AND deleted_at IS NULL`

		result, err := tx.ExecContext(ctx, queryUpdate, args...)
		if err != nil {
			logger.Errorf("Err on result exec content: %v", err)

			return err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil || rowsAffected == 0 {
			if err == nil {
				err = sql.ErrNoRows
			}
			logger.Errorf("Err on rows affected: %v", err)

			return err
		}
	}

	if commitErr := tx.Commit(); commitErr != nil {
		err := pkg.ValidateBusinessError(constant.ErrEntityNotFound, reflect.TypeOf(mmodel.Account{}).Name())

		mopentelemetry.HandleSpanError(&span, "Failed to commit balances", err)

		return commitErr
	}

	return nil
}
