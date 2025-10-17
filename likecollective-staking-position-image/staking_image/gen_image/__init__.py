import io

from .gen_image import gen_image
from .params import Params


def gen_image_by_params(params: Params, format: str = "PNG") -> bytes:
    image = gen_image(
        book_nft_address=params.book_nft_address,
        staked_amount=params.staked_amount,
        reward_index=params.reward_index,
        initial_staker=params.initial_staker,
        book_nft_name=params.book_nft_name,
    )

    img_byte_arr = io.BytesIO()
    image.save(img_byte_arr, format=format)
    return img_byte_arr.getvalue()
