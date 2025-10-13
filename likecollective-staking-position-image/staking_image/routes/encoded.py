from fastapi import APIRouter, Response

from staking_image.gen_image import gen_image_by_encoded_string

router = APIRouter()


@router.get(
    "/{encoded_string}",
    response_class=Response,
)
def get_image_from_encoded_string(
    encoded_string: str,
):
    content = gen_image_by_encoded_string(encoded_string, format="PNG")

    return Response(
        content=content,
        media_type="image/png",
    )
