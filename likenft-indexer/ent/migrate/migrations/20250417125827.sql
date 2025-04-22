-- Modify "evm_events" table
ALTER TABLE "evm_events"
    ADD COLUMN "chain_id" numeric NULL,
    ADD COLUMN "name" character varying NULL,
    ADD COLUMN "signature" character varying NULL,
    ADD COLUMN "indexed_params" jsonb NULL,
    ADD COLUMN "non_indexed_params" jsonb NULL;
UPDATE evm_events SET
    chain_id='11155420',
    name=topic0,
    signature=topic0,
    indexed_params='{}',
    non_indexed_params='{}';
ALTER TABLE "evm_events"
    ALTER COLUMN "chain_id" SET NOT NULL,
    ALTER COLUMN "name" SET NOT NULL,
    ALTER COLUMN "signature" SET NOT NULL,
    ALTER COLUMN "indexed_params" SET NOT NULL,
    ALTER COLUMN "non_indexed_params" SET NOT NULL;
CREATE INDEX "evmevent_signature" ON "evm_events" ("signature");
