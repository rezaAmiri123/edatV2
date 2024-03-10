package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	// "github.com/jackc/pgx/pgtype"
	"github.com/rezaAmiri123/edatV2/am"
	"github.com/rezaAmiri123/edatV2/tm"
	"github.com/stackus/errors"
)

type OutboxStore struct {
	tableName string
	db        DB
}

var _ tm.OutboxStore = (*OutboxStore)(nil)

func NewOutboxStore(tableName string, db DB) OutboxStore {
	return OutboxStore{
		tableName: tableName,
		db:        db,
	}
}

func (s OutboxStore) Save(ctx context.Context, msg am.Message) error {
	const query = "INSERT INTO %s (id, NAME, subject, DATA, metadata, sent_at) VALUES ($1, $2, $3, $4, $5, $6)"

	metadata, err := json.Marshal(msg.Metadata())
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(ctx, s.table(query),
		msg.ID(),
		msg.MessageName(),
		msg.Subject(),
		msg.Data(),
		metadata,
		msg.SentAt(),
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return tm.ErrDuplicteMessage(msg.ID())
			}
		}
	}
	return err
}

func (s OutboxStore) FindUnpublished(ctx context.Context, limit int) ([]am.Message, error) {
	const query = `SELECT id, name, subject, data, metadata, sent_at 
		FROM %s 
		WHERE published_at IS NULL 
		LIMIT %d`

	rows, err := s.db.QueryContext(ctx, s.table(query, limit))
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			err = errors.Wrapf(err, "closing event rows")
		}
	}(rows)

	var msgs []am.Message

	for rows.Next() {
		var metadata []byte
		msg := outboxMessage{}
		err = rows.Scan(
			&msg.id,
			&msg.name,
			&msg.subject,
			&msg.data,
			&metadata,
			&msg.sentAt,
		)
		if err != nil {
			return msgs, err
		}

		err = json.Unmarshal(metadata, &msg.metadata)
		if err != nil {
			return msgs, err
		}

		msgs = append(msgs, msg)
	}

	return msgs, nil
}

func (s OutboxStore) MarkPublished(ctx context.Context, ids ...string) error {
	const query = `UPDATE %s 
		SET published_at = CURRENT_TIMESTAMP 
		WHERE id = ANY ($1)`
	
	msgIDs := &pgtype.TextA
	err := msgIDs.Dims
	
}

func (s OutboxStore) table(query string, args ...any) string {
	params := []any{s.tableName}
	params = append(params, args...)
	return fmt.Sprintf(query, args...)
}
