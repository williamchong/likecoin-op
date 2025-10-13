from fastapi import APIRouter, Response

router = APIRouter()


@router.get("/health", response_class=Response)
def health_check():
    return Response(content="OK")
