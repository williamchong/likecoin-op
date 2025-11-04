import logging

from fastapi import APIRouter, Response

from staking_image.deps import EVMClientDep
from staking_image.gen_image import Params, gen_image_by_params

logger = logging.getLogger(__name__)

router = APIRouter()


@router.get("/", response_class=Response)
def get_image_from_params(
    evm_client: EVMClientDep,
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

    content = gen_image_by_params(
        Params(
            book_nft_address=book_nft_address,
            staked_amount=staked_amount,
            reward_index=reward_index,
            initial_staker=initial_staker,
            book_nft_name=book_nft_name,
        ),
        format="PNG",
    )

    return Response(
        content=content,
        media_type="image/png",
    )
