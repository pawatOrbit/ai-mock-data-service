from fastapi import FastAPI
from app.router.ai_mock_data_route import router as ai_mock_data_router
from app.database.db import database
import logging

app = FastAPI(title="AI Mock Data Service", version="1.0.0")


@app.on_event("startup")
async def startup():
    logging.info("Starting up the application...")
    await database.connect()

@app.on_event("shutdown")
async def shutdown():
    logging.info("Shutting down the application...")
    await database.disconnect()


app.include_router(ai_mock_data_router, prefix="/api")