import decimal
import logging

from fastapi import APIRouter
from pydantic import BaseModel, HttpUrl

from staking_image.deps import ConfigDep, EVMClientDep
from staking_image.gen_image.params import parse_base64_encoded_params

logger = logging.getLogger(__name__)

router = APIRouter()


class Response(BaseModel):
    name: str | None = None
    description: str
    image: HttpUrl
    stakeholder: str


@router.get(
    "/{encoded_string}",
    response_model=Response,
)
def get_image_from_encoded_string(
    evm_client: EVMClientDep,
    config: ConfigDep,
    encoded_string: str,
):
    params = parse_base64_encoded_params(encoded_string)

    likecoin_decimals = evm_client.get_likecoin_decimals()
    params.staked_amount = str(
        decimal.Decimal(params.staked_amount)
        / (decimal.Decimal(10) ** decimal.Decimal(likecoin_decimals))
    )

    if params.book_nft_name is None:
        try:
            params.book_nft_name = evm_client.get_likenft_name(params.book_nft_address)
        except Exception as e:
            logger.warning(f"Failed to get book NFT name: {e}")

    return Response(
        name=params.book_nft_name,
        description=f"NFT for {params.staked_amount} LIKE stake in {params.book_nft_name or params.book_nft_address}",
        image=config.base_url.build(
            scheme=config.base_url.scheme,
            host=config.base_url.host,
            port=config.base_url.port,
            path=f"image/{encoded_string}",
        ),
        stakeholder=params.initial_staker,
    )
