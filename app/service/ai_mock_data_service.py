from app.repository.database_table import get_database_table
from app.prompt.ai_mock_data_service import generate_mock_data_prompt
from app.client.ai_client import query_lm
from app.config import lm_deepseek_model
from app.model.http_response.ai_model_response import AiMockDataResponseBase

async def ai_mock_data_service(table_name: str, num_sample: int)-> AiMockDataResponseBase| None:
    if num_sample <= 0:
        print("Number of samples must be greater than 0.")
        return None
    
    data = await get_database_table(table_name)

    database_script = data.table_script

    promptForService = generate_mock_data_prompt(table_name, database_script, num_sample)

    resultFromAI = await query_lm(prompt=promptForService, model=lm_deepseek_model, max_tokens=12000)

    if resultFromAI:
        return resultFromAI
    else:
        print("No result from AI model.")
        return None