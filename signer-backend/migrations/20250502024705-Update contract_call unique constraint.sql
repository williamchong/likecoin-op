
-- +migrate Up

alter table "contract_call"
    drop constraint contract_call_contract_address_method_params_hex_key;
alter table "contract_call"
    add constraint contract_call_contract_address_method_params_hex_key
    exclude using hash ((contract_address || method || params_hex) with =);

-- +migrate Down
alter table "contract_call"
    drop constraint contract_call_contract_address_method_params_hex_key;
alter table "contract_call"
    add constraint contract_call_contract_address_method_params_hex_key
    unique (contract_address,method,params_hex);
