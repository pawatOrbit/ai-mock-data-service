from pydantic import BaseModel


class MockDataGeneratorRequest(BaseModel):
    table_name: str
    num_samples: int

class MockDataGeneratorResponseData(BaseModel):
    query: str

class MockDataGeneratorResponse(BaseModel):
    status: int
    data: MockDataGeneratorResponseData
