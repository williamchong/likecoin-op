
-- +migrate Up
CREATE TABLE likecoin_evm_migrate_book_submission (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    like_class_id VARCHAR(255) NOT NULL,
    evm_class_id VARCHAR(42) NOT NULL,
    status VARCHAR(255) NOT NULL,
    failed_reason TEXT NULL
);

-- +migrate Down
DROP TABLE likecoin_evm_migrate_book_submission;
