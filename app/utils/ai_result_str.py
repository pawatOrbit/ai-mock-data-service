from app.model.http_response.ai_model_response import AiMockDataResponseBase
import re
import logging

async def extract_thinking_and_context(result: str)-> AiMockDataResponseBase:

    logging.info(f"Extracting thinking and context from result: {result}")
    # Extract content between <think> and </think>
    resultContext = await extract_response(result)

    return AiMockDataResponseBase(
        response=resultContext,
    )

async def extract_response(result: str) -> str:
    result = result.replace("\n", "")
    pattern = r"(INSERT INTO .*?;)"
    matches = re.findall(pattern, result, re.DOTALL)
    return "".join(matches)
