from fastapi import FastAPI
from app.router.mock_data_generator_route import router as mock_data_generator_route
from app.router.table_schemas_route import router as table_schemas_router
from app.core.lifecycle.lifecycle import lifespan
from app.core.log.logger import setup_logging
from app.core.middleware.logging_middleware import LoggingMiddleware

setup_logging()

app = FastAPI(
    title="AI Mock Data Service",
    version="1.0.0",
    lifespan=lifespan
)

app.add_middleware(
    LoggingMiddleware,
)

API_PREFIX = "/api/v1"

app.include_router(mock_data_generator_route, prefix=API_PREFIX, tags=["Mock Data Generator"])
app.include_router(table_schemas_router, prefix=API_PREFIX, tags=["Table Schemas"])
