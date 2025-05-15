import httpx
from app.config import lm_path, lm_deepseek_timeout
from app.model.http_response.ai_model_response import AiMockDataResponseBase
from app.utils.ai_result_str import extract_thinking_and_context


async def query_lm(prompt: str,model: str, temperature: float = 0.7, max_tokens: int = 512) -> AiMockDataResponseBase:
    payload = {
        "model": model,
        "messages": [
            {"role": "user", "content": prompt}
        ],
        "temperature": temperature,
        "max_tokens": max_tokens,
    }

    async with httpx.AsyncClient(timeout=lm_deepseek_timeout) as client:
        response = await client.post(lm_path, json=payload)
        response.raise_for_status()
        result = response.json()

        contentResult = result["choices"][0]["message"]["content"]

        return await extract_thinking_and_context(contentResult)