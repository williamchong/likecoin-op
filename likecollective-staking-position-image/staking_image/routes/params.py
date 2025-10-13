from fastapi import APIRouter

router = APIRouter()


@router.get("/")
def get_image_from_params(
    book_nft_address: str,
    staked_amount: str,
    reward_index: str,
    initial_staker: str,
    book_nft_name: str | None = None,
):
    return {
        "book_nft_address": book_nft_address,
        "staked_amount": staked_amount,
        "reward_index": reward_index,
        "initial_staker": initial_staker,
        "book_nft_name": book_nft_name,
    }
