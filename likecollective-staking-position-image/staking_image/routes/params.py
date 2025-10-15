from fastapi import APIRouter, Response

from staking_image.gen_image import Params, gen_image_by_params

router = APIRouter()


@router.get("/", response_class=Response)
def get_image_from_params(
    book_nft_address: str,
    staked_amount: str,
    reward_index: str,
    initial_staker: str,
    book_nft_name: str | None = None,
):
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
