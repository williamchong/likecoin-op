import decimal
import logging

from fastapi import APIRouter, Response

from staking_image.deps import EVMClientDep
from staking_image.gen_image import gen_image_by_params
from staking_image.gen_image.params import parse_base64_encoded_params

logger = logging.getLogger(__name__)

router = APIRouter()


@router.get(
    "/{encoded_string}",
    response_class=Response,
)
def get_image_from_encoded_string(
    evm_client: EVMClientDep,
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

    content = gen_image_by_params(params, format="PNG")

    return Response(
        content=content,
        media_type="image/png",
    )
