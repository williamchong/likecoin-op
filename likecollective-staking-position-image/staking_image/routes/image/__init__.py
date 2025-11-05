from fastapi import APIRouter

from .encoded import router as encoded_router
from .params import router as params_router

router = APIRouter(tags=["image"], prefix="/image")

router.include_router(encoded_router)
router.include_router(params_router)
