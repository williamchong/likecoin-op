from PIL import Image, ImageDraw

IMAGE_SIZE = (290, 500)
IMAGE_BACKGROUND_COLOR = (0, 0, 0, 0)


def gen_image(
    book_nft_address: str,
    staked_amount: str,
    reward_index: str,
    initial_staker: str,
    book_nft_name: str | None = None,
) -> Image.Image:
    """
    Generate an image with the given parameters.

    Args:
        `book_nft_address`: The address of the book NFT.
            It is a 42 characters long string e.g. `0x1234567890123456789012345678901234567890`
        `staked_amount`: The amount of staked tokens.
            It is a string representing a number e.g. `123`
        `reward_index`: The reward index.
            It is a string representing a number e.g. `123`
        `initial_staker`: The address of the initial staker.
            It is a 42 characters long string e.g. `0x1234567890123456789012345678901234567890`
        `book_nft_name`: The name of the book NFT.

    Returns:
        An image with the given parameters.
    """

    image = Image.new("RGBA", IMAGE_SIZE, IMAGE_BACKGROUND_COLOR)

    draw = ImageDraw.Draw(image)

    if book_nft_name:
        draw.text((10, 10), book_nft_name, fill=(0, 0, 0, 255))
    draw.text(
        (10, 30), f"Book NFT Address: {book_nft_address}", fill=(255, 255, 255, 255)
    )
    draw.text((10, 50), f"Staked Amount: {staked_amount}", fill=(255, 255, 255, 255))
    draw.text((10, 70), f"Reward Index: {reward_index}", fill=(255, 255, 255, 255))
    draw.text((10, 90), f"Initial Staker: {initial_staker}", fill=(255, 255, 255, 255))

    return image
