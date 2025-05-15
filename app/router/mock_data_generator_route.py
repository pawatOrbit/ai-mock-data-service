from fastapi import APIRouter
from app.model.http.mock_data_generator_req_resp import MockDataGeneratorRequest, MockDataGeneratorResponse
from app.service.mock_data_generator_service import mock_data_generator_service

router = APIRouter(
    prefix="/generator",
    tags=["Mock Data Generator"],
)


@router.post("/mock-data", response_model=MockDataGeneratorResponse)
async def generate_mock_data(request: MockDataGeneratorRequest):
    """
    Generate mock data for a given table name.
    """
    table_name = request.table_name
    num_samples = request.num_samples
    result = await mock_data_generator_service(table_name,num_samples)
    if result:
        return result
    else:
        return MockDataGeneratorResponse(status=500, data=None)