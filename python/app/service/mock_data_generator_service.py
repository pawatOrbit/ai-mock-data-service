from app.repository.database_schemas import get_schema_by_table_name
from app.prompt.mock_data_generator_prompt import generate_prompt_without_key,generate_prompt_for_get_fk_table_and_fk_field_and_linked_field_from_table_script,generate_prompt_for_mock_data_with_values_and_fields
from app.client.lm_studio import query_lm
from app.core.config.config import lm_deepseek_model
from app.model.http.mock_data_generator_req_resp import MockDataGeneratorResponse,MockDataGeneratorResponseData
from datetime import datetime
from app.utils.extract_string import extract_foreign_key_info, extract_insert_values
import logging

async def mock_data_generator_service(table_name: str, num_sample: int)-> MockDataGeneratorResponse| None:
    if num_sample <= 0:
        logging.error("Number of samples must be greater than 0.")
        return None
    
    datetime_before = datetime.now()
    
    data = await get_schema_by_table_name(table_name)

    database_script = data.table_script

    prompt_for_service = generate_prompt_without_key(table_name, database_script, num_sample)

    result_from_aI = await query_lm(prompt=prompt_for_service, model=lm_deepseek_model, max_tokens=12000)

    time_taken = datetime.now() - datetime_before
    resultService = MockDataGeneratorResponseData(
        query=result_from_aI.response,
        prompt_tokens=result_from_aI.prompt_tokens,
        completion_tokens=result_from_aI.completion_tokens,
        total_tokens=result_from_aI.total_tokens,
        time_taken=time_taken.total_seconds()
    )

    if result_from_aI:
        return MockDataGeneratorResponse(
                status=200,
                data=resultService
            )
    else:
        print("Failed to generate mock data.")
        return None
    
async def generate_mock_data_fk_version_service(table_name: str, num_sample: int)-> MockDataGeneratorResponse| None:
    if num_sample <= 0:
        logging.error("Number of samples must be greater than 0.")
        return None
    
    datetime_before = datetime.now()
    
    data = await get_schema_by_table_name(table_name)

    database_script = data.table_script

    prompt_for_fk = generate_prompt_for_get_fk_table_and_fk_field_and_linked_field_from_table_script(table_name, database_script)

    prompt_token = 0
    completion_token = 0
    total_token = 0

    get_table_join_and_field = await query_lm(prompt=prompt_for_fk, model=lm_deepseek_model, max_tokens=12000)

    link_fields, table_and_field = extract_foreign_key_info(get_table_join_and_field.response)

    prompt_token += get_table_join_and_field.prompt_tokens
    completion_token += get_table_join_and_field.completion_tokens
    total_token += get_table_join_and_field.total_tokens

    response_insert_script = []

    values = []

    index = 0
    for fk_table_name, fk_table_field in table_and_field.items():
        logging.info(f"Table: {fk_table_name}, Field: {fk_table_field}")
        logging.info(f"Link Field by index: {link_fields[index]}")
        index += 1

        fk_data_schema = await get_schema_by_table_name(fk_table_name)

        prompt_for_generate_fk_data = generate_prompt_without_key(table_name=fk_table_name, table_script=fk_data_schema.table_script, num_sample=1)

        result_from_aI = await query_lm(prompt=prompt_for_generate_fk_data, model=lm_deepseek_model, max_tokens=12000)
        prompt_token += result_from_aI.prompt_tokens
        completion_token += result_from_aI.completion_tokens
        total_token += result_from_aI.total_tokens

        if result_from_aI:
            response_insert_script.append(result_from_aI.response)
        else:
            logging.error(f"Failed to generate mock data for {fk_table_name}.")
            raise Exception(f"Failed to generate mock data for {fk_table_name}.")
        
        insert_value_and_field = extract_insert_values(result_from_aI.response)
        values.append(insert_value_and_field[fk_table_field])

    # Generate the main table mock data
    prompt_for_service = generate_prompt_for_mock_data_with_values_and_fields(table_name=table_name, table_script=database_script, num_sample=num_sample, fields_name=link_fields, fields_value=values)

    result_from_main_table = await query_lm(prompt=prompt_for_service, model=lm_deepseek_model, max_tokens=12000)

    # Extract the main table mock data
    query_fk = ' '.join(response_insert_script)

    prompt_token += result_from_main_table.prompt_tokens
    completion_token += result_from_main_table.completion_tokens
    total_token += result_from_main_table.total_tokens
    time_taken = datetime.now() - datetime_before

    resultService = MockDataGeneratorResponseData(
        query=result_from_main_table.response+ " " + query_fk,
        prompt_tokens=prompt_token,
        completion_tokens=completion_token,
        total_tokens=total_token,
        time_taken=time_taken.total_seconds()
    )

    if result_from_main_table:
        return MockDataGeneratorResponse(
                status=200,
                data=resultService
            )
    else:
        print("Failed to generate mock data.")
        return None


