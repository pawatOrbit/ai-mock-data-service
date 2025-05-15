from pydantic import BaseModel

class AiMockDataResponseBase(BaseModel):
    response: str