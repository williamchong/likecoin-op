from pydantic import ConfigDict
from pydantic_settings import BaseSettings


class Config(BaseSettings, frozen=True):
    rpc_url: str
    likecoin_contract_address: str

    model_config = ConfigDict(env_file=".env", env_file_encoding="utf-8")


config = Config()
