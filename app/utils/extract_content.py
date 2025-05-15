from app.model.client.ai_model_response import LmStudioContentResponse
import re
import logging

async def extract_thinking_and_context(result: str)-> LmStudioContentResponse:

    logging.info(f"Extracting thinking and context from result: {result}")
    # Extract content between <think> and </think>
    resultContext = await extract_response(result)

    return LmStudioContentResponse(
        response=resultContext,
    )

async def extract_response(result: str) -> str:
    result = result.replace("\n", "")
    pattern = r"(INSERT INTO .*?;)"
    matches = re.findall(pattern, result, re.DOTALL)
    return "".join(matches)
