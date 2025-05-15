from app.repository.database_schemas import get_schema_by_table_name
from app.prompt.mock_data_generator_prompt import generate_prompt_without_key
from app.client.lm_studio import query_lm
from app.core.config.config import lm_deepseek_model
from app.model.http.mock_data_generator_req_resp import MockDataGeneratorResponse,MockDataGeneratorResponseData
from datetime import datetime

async def mock_data_generator_service(table_name: str, num_sample: int)-> MockDataGeneratorResponse| None:
    if num_sample <= 0:
        print("Number of samples must be greater than 0.")
        return None
    
    datetimeBefore = datetime.now()
    
    data = await get_schema_by_table_name(table_name)

    database_script = data.table_script

    promptForService = generate_prompt_without_key(table_name, database_script, num_sample)

    resultFromAI = await query_lm(prompt=promptForService, model=lm_deepseek_model, max_tokens=12000)

    timeTaken = datetime.now() - datetimeBefore
    resultService = MockDataGeneratorResponseData(
        query=resultFromAI.response,
        prompt_tokens=resultFromAI.prompt_tokens,
        completion_tokens=resultFromAI.completion_tokens,
        total_tokens=resultFromAI.total_tokens,
        time_taken=timeTaken.total_seconds()
    )

    if resultFromAI:
        return MockDataGeneratorResponse(
                status=200,
                data=resultService
            )
    else:
        print("Failed to generate mock data.")
        return None