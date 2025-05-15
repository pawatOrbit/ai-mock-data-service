from app.model.client.ai_model_response import LmStudioContentResponse
import re
import logging

async def extract_thinking_and_context(result: str, prompt_tokens: int = 0, completion_tokens: int = 0, total_tokens: int = 0) -> LmStudioContentResponse:

    logging.info(f"Extracting thinking and context from result: {result}")
    # Extract content between <think> and </think>
    result = result.replace("\n", "")

    return LmStudioContentResponse(
        response=result,
        prompt_tokens=prompt_tokens,
        completion_tokens=completion_tokens,
        total_tokens=total_tokens
    )