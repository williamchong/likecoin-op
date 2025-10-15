from fastapi import FastAPI

from staking_image.routes import router

app = FastAPI()
app.include_router(router)
