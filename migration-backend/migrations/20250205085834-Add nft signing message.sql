
-- +migrate Up

CREATE TABLE nft_signing_message
(
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP DEFAULT now(),
  cosmos_address text NOT NULL,
  liker_id text NOT NULL,
  eth_address text NOT NULL,
  nonce text NOT NULL,
  issue_time timestamp WITH TIME ZONE,
  message text NOT NULL,
  UNIQUE (cosmos_address, liker_id, eth_address, nonce)
);

-- +migrate Down

DROP TABLE nft_signing_message;
