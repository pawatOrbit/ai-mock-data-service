from pydantic import BaseModel

class LmStudioContentResponse(BaseModel):
    response: str
    prompt_tokens: int = 0
    completion_tokens: int = 0
    total_tokens: int = 0