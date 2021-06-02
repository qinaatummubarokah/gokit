package reportcsv

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/go-kit/kit/log"
)

var RepoErr = errors.New("Unable to handle Repo Request")

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *repo) GetData(ctx context.Context, id string) ([]Csvdata, error) {
	var data []Csvdata

	err := repo.db.QueryRow(`
	SELECT DISTINCT
	payment_transactions.id,
	users.full_name AS user_name
	FROM "payment_transactions"
	LEFT JOIN voucher_profiles vp ON vp.id = payment_transactions.cc_identifier::integer
	LEFT JOIN users ON users.id = vp.user_id
	LEFT JOIN deprtments ON departments.id = vp.department_id
	LEFT JOIN accounts ON accounts.id = payment_transactions.account_id
	LEFT JOIN payment_transaction_infos info ON info.payment_transaction_id = payment_transactions.id
	WHERE
	"payment_transactions"."state" IN (2, 3) AND
	"payment_transactions"."payment_type" IN ('ecv', 'edc') AND
	("payment_transactions"."completed_at" BETWEEN '2017-07-01 00:00:00.000000' AND '2021-05-28 00:00:00.000000')`).Scan(&data)

	if err != nil {
		return data, RepoErr
	}

	return data, nil
}
