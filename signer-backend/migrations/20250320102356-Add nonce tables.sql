
-- +migrate Up

CREATE TABLE evm_transaction_request (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    status TEXT NOT NULL,
    signer_address VARCHAR(42) NOT NULL,
    to_address VARCHAR(42) NOT NULL,
    amount TEXT NOT NULL,
    method TEXT NOT NULL,
    params_hex TEXT NOT NULL,
    call_data_hex TEXT NOT NULL,
    gas_limit NUMERIC NULL CHECK (gas_limit > 0),
    gas_price TEXT NULL,
    nonce NUMERIC NULL CHECK (nonce > 0),
    signed_tx_hash VARCHAR(66) NULL,
    submitted_at TIMESTAMP NULL,
    block_hash VARCHAR(66) NULL,
    receipt_status TEXT NULL,
    failed_reason TEXT NULL
);

CREATE TABLE transaction_nonce (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    eth_address VARCHAR(42) NOT NULL,
    nonce NUMERIC NOT NULL CHECK (nonce > 0),   
    evm_transaction_request_id INTEGER NOT NULL,
    FOREIGN KEY (evm_transaction_request_id) REFERENCES evm_transaction_request(id),
    UNIQUE (eth_address, nonce)
);

CREATE TABLE contract_call (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    contract_address VARCHAR(42) NOT NULL,
    method TEXT NOT NULL,
    params_hex TEXT NOT NULL,
    evm_transaction_request_id INTEGER NOT NULL,
    FOREIGN KEY (evm_transaction_request_id) REFERENCES evm_transaction_request(id),
    UNIQUE (contract_address, method, params_hex)
);

-- +migrate Down

DROP TABLE contract_call;
DROP TABLE transaction_nonce;
DROP TABLE evm_transaction_request;
