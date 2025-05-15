from fastapi import APIRouter
from app.model.request_response.ai_mock_data_req_resp import AIMockDataRequest, AIMockDataResponse
from app.service.ai_mock_data_service import ai_mock_data_service

router = APIRouter()


@router.post("/ai/mock_data", response_model=AIMockDataResponse)
async def ai_mock_data(request: AIMockDataRequest):
    """
    Generate mock data for a given table name.
    """
    table_name = request.table_name
    num_samples = request.num_samples
    result = await ai_mock_data_service(table_name,num_samples)
    if result:
        return AIMockDataResponse(response=result)
    else:
        return AIMockDataResponse(response=None)