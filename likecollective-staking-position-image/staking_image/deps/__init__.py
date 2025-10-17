from functools import lru_cache
from typing import Annotated

from fastapi import Depends

from staking_image.config import Config, config
from staking_image.evm import EVMClient


@lru_cache(maxsize=1)
def get_config() -> Config:
    return config


ConfigDep = Annotated[Config, Depends(get_config)]


@lru_cache(maxsize=1)
def get_evm_client(config: ConfigDep) -> EVMClient:
    return EVMClient(
        rpc_url=config.rpc_url,
        likecoin_contract_address=config.likecoin_contract_address,
    )


EVMClientDep = Annotated[EVMClient, Depends(get_evm_client)]


__all__ = ["ConfigDep", "EVMClientDep"]
