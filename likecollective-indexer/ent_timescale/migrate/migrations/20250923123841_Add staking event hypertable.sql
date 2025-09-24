-- Create "staking_events_hyper_table" table
CREATE TABLE "public"."staking_events_hyper_table" (
  "transaction_hash" character varying NOT NULL,
  "transaction_index" bigint NOT NULL,
  "block_number" numeric NOT NULL,
  "log_index" bigint NOT NULL,
  "event_type" character varying NOT NULL DEFAULT 'staked',
  "nft_class_address" character varying NOT NULL,
  "account_evm_address" character varying NOT NULL,
  "staked_amount_added" numeric NOT NULL DEFAULT 0,
  "staked_amount_removed" numeric NOT NULL DEFAULT 0,
  "pending_reward_amount_added" numeric NOT NULL DEFAULT 0,
  "pending_reward_amount_removed" numeric NOT NULL DEFAULT 0,
  "claimed_reward_amount_added" numeric NOT NULL DEFAULT 0,
  "claimed_reward_amount_removed" numeric NOT NULL DEFAULT 0,
  "datetime" timestamptz NOT NULL
)
WITH (
  timescaledb.hypertable,
  timescaledb.partition_column='datetime',
  timescaledb.chunk_interval='7 day',
  timescaledb.segmentby='event_type'
);

CREATE OR REPLACE FUNCTION index_staking_events_to_hyper_table()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO public.staking_events_hyper_table (
        transaction_hash,
        transaction_index,
        block_number,
        log_index,
        event_type,
        nft_class_address,
        account_evm_address,
        staked_amount_added,
        staked_amount_removed,
        pending_reward_amount_added,
        pending_reward_amount_removed,
        claimed_reward_amount_added,
        claimed_reward_amount_removed,
        datetime)
    VALUES (NEW.transaction_hash,
        NEW.transaction_index,
        NEW.block_number,
        NEW.log_index,
        NEW.event_type,
        NEW.nft_class_address,
        NEW.account_evm_address,
        NEW.staked_amount_added,
        NEW.staked_amount_removed,
        NEW.pending_reward_amount_added,
        NEW.pending_reward_amount_removed,
        NEW.claimed_reward_amount_added,
        NEW.claimed_reward_amount_removed,
        NEW.datetime);
        RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER index_staking_events_to_hyper_table_insert_trigger
AFTER INSERT ON staking_events
FOR EACH ROW
EXECUTE FUNCTION index_staking_events_to_hyper_table();

ALTER TABLE staking_events ENABLE REPLICA TRIGGER index_staking_events_to_hyper_table_insert_trigger;
