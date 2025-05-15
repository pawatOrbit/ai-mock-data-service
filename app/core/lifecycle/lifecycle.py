from contextlib import asynccontextmanager
from fastapi import FastAPI
from app.core.pgdb.db import database
import logging

@asynccontextmanager
async def lifespan(app: FastAPI):
    logging.info("Starting up the application...")
    await database.connect()

    yield

    logging.info("Shutting down the application...")
    await database.disconnect()
