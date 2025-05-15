from pydantic import BaseModel
from app.model.http_response.ai_model_response import AiMockDataResponseBase


class AIMockDataRequest(BaseModel):
    table_name: str
    num_samples: int

class AIMockDataResponse(BaseModel):
    response: AiMockDataResponseBase
