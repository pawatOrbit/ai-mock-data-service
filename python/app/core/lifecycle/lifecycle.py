from contextlib import asynccontextmanager
from fastapi import FastAPI
from app.core.pgdb.db import database
import logging

@asynccontextmanager
async def lifespan(app: FastAPI):
    logging.info("Starting up the application...")
    await database.connect()
    logging.info("Connected to the database successfully.")

    yield

    logging.info("Shutting down the application...")
    await database.disconnect()
    logging.info("Disconnected from the database successfully.")
    logging.info("The application has been shut down.")
