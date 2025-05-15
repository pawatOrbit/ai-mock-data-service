from fastapi import FastAPI
from app.router.mock_data_generator_route import router as ai_mock_data_router
from app.core.lifecycle.lifecycle import lifespan
from app.core.log.logger import setup_logging

# Initialize logging
setup_logging()

app = FastAPI(
    title="AI Mock Data Service",
    version="1.0.0",
    lifespan=lifespan
)


app.include_router(ai_mock_data_router, prefix="/v1")