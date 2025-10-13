from fastapi import APIRouter

router = APIRouter()


@router.get("/{encoded_string}")
def get_image_from_encoded_string(
    encoded_string: str,
):
    return {
        "encoded_string": encoded_string,
    }
