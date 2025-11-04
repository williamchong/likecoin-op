from pydantic import ConfigDict, Field, HttpUrl
from pydantic_settings import BaseSettings


class Config(BaseSettings, frozen=True):
    base_url: HttpUrl = Field(default=HttpUrl("http://localhost:8000"))
    rpc_url: str
    likecoin_contract_address: str

    model_config = ConfigDict(env_file=".env", env_file_encoding="utf-8")


config = Config()
