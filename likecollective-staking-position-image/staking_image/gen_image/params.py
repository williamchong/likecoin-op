import base64
from typing import Any

from eth_abi import decode
from pydantic import BaseModel


class Params(BaseModel):
    book_nft_address: str
    staked_amount: str
    reward_index: str
    initial_staker: str
    book_nft_name: str | None = None


def parse_base64_encoded_params(v: Any) -> Params:
    if not isinstance(v, str):
        raise ValueError("Invalid base64 encoded string")

    try:
        data = base64.b64decode(v)

        # Address is 20 bytes
        address_bytes = data[:20].rjust(32, b"\x00")
        # Uint256 is 32 bytes
        staked_amount_bytes = data[20:52]
        # Uint256 is 32 bytes
        reward_index_bytes = data[52:84]
        # Address is 20 bytes
        initial_staker_bytes = data[84:104].rjust(32, b"\x00")

        (book_nft_address, staked_amount, reward_index, initial_staker) = decode(
            ["address", "uint256", "uint256", "address"],
            address_bytes
            + staked_amount_bytes
            + reward_index_bytes
            + initial_staker_bytes,
        )

        return Params(
            book_nft_address=book_nft_address,
            staked_amount=f"{staked_amount}",
            reward_index=f"{reward_index}",
            initial_staker=initial_staker,
            book_nft_name=None,
        )
    except Exception as e:
        raise ValueError(f"Invalid base64 encoded string: {e}")
