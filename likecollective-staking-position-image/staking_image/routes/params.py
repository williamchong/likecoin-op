import logging
from urllib.parse import urlencode

from fastapi import APIRouter, Response
from pydantic import BaseModel, HttpUrl

from staking_image.deps import ConfigDep, EVMClientDep
from staking_image.gen_image.params import Params

logger = logging.getLogger(__name__)

router = APIRouter()


class Response(BaseModel):
    name: str | None = None
    description: str
    image: HttpUrl
    stakeholder: str


@router.get("/", response_model=Response)
def get_image_from_params(
    evm_client: EVMClientDep,
    config: ConfigDep,
    book_nft_address: str,
    staked_amount: str,
    reward_index: str,
    initial_staker: str,
    book_nft_name: str | None = None,
):
    if book_nft_name is None:
        try:
            book_nft_name = evm_client.get_likenft_name(book_nft_address)
        except Exception as e:
            logger.warning(f"Failed to get book NFT name: {e}")

    params = Params(
        book_nft_address=book_nft_address,
        staked_amount=staked_amount,
        reward_index=reward_index,
        initial_staker=initial_staker,
        book_nft_name=book_nft_name,
    )

    image_url_params = urlencode(params.model_dump())

    return Response(
        name=book_nft_name,
        description=f"NFT for {staked_amount} LIKE stake in {book_nft_name}",
        image=config.base_url.build(
            scheme=config.base_url.scheme,
            host=config.base_url.host,
            port=config.base_url.port,
            path="image/",
            query=image_url_params,
        ),
        stakeholder=initial_staker,
    )
