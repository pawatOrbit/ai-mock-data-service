from pydantic import BaseModel
from datetime import datetime
class MockDataGeneratorRequest(BaseModel):

    table_name: str
    num_samples: int

class MockDataGeneratorResponseData(BaseModel):
    query: str
    prompt_tokens: int = 0
    completion_tokens: int = 0
    total_tokens: int = 0
    time_taken: float = 0.0

class MockDataGeneratorResponse(BaseModel):
    status: int
    data: MockDataGeneratorResponseData
