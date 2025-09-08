
-- +migrate Up

ALTER TABLE transaction_nonce DROP CONSTRAINT transaction_nonce_nonce_check;
ALTER TABLE transaction_nonce ADD CONSTRAINT transaction_nonce_nonce_check CHECK (nonce >= 0);

ALTER TABLE evm_transaction_request DROP CONSTRAINT evm_transaction_request_nonce_check;
ALTER TABLE evm_transaction_request ADD CONSTRAINT evm_transaction_request_nonce_check CHECK (nonce >= 0);

-- +migrate Down

ALTER TABLE transaction_nonce DROP CONSTRAINT transaction_nonce_nonce_check;
ALTER TABLE transaction_nonce ADD CONSTRAINT transaction_nonce_nonce_check CHECK (nonce > 0);

ALTER TABLE evm_transaction_request DROP CONSTRAINT evm_transaction_request_nonce_check;
ALTER TABLE evm_transaction_request ADD CONSTRAINT evm_transaction_request_nonce_check CHECK (nonce > 0);
