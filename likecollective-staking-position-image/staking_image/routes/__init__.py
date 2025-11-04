from fastapi import APIRouter

from .health import router as health_router
from .image import router as image_router

router = APIRouter(tags=["image"])

router.include_router(health_router)
router.include_router(image_router)
