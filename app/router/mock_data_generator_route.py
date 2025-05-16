from fastapi import APIRouter
from app.model.http.mock_data_generator_req_resp import MockDataGeneratorRequest, MockDataGeneratorResponse
from app.service.mock_data_generator_service import (
    mock_data_generator_service,
    generate_mock_data_fk_version_service
)

router = APIRouter(
    prefix="/generator",
    tags=["Mock Data Generator"],
)

@router.post("/mock-data/basic", response_model=MockDataGeneratorResponse)
async def generate_mock_data_basic(request: MockDataGeneratorRequest):
    """
    Generate mock data for a table without handling foreign key relations.
    """
    result = await mock_data_generator_service(request.table_name, request.num_samples)
    if result:
        return result
    return MockDataGeneratorResponse(status=500, data=None)


@router.post("/mock-data/with-fk", response_model=MockDataGeneratorResponse)
async def generate_mock_data_with_fk(request: MockDataGeneratorRequest):
    """
    Generate mock data for a table including handling of foreign key relations.
    """
    result = await generate_mock_data_fk_version_service(request.table_name, request.num_samples)
    if result:
        return result
    return MockDataGeneratorResponse(status=500, data=None)
