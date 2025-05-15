def generate_mock_data_prompt(table_name: str, table_script: str, num_sample:int ) -> str:
    prompt = f"""
    You are a data generation expert. Your task is to create mock data for a database table.
    Table Name: {table_name}
    Table Structure:
    {table_script}
    Please generate {num_sample} rows of mock data for this table. Ensure that the data is realistic and adheres to the structure defined above.
    Please not thinking, just give me the result (the SQL insert statement) in the following format:
    Example of the mock data format:
    INSERT INTO {table_name} (column1, column2, column3) VALUES (value1, value2, value3);
    """
    
    return prompt


def generate_prompt_for_get_fk_table(table_name: str, table_script: str) -> str:
    prompt = f"""
    You are a data generation expert. Your task is list all the foreign key tables for a database table.
    Table Name: {table_name}
    Table Structure:
    {table_script}
    Please list all the foreign key tables for this table. Ensure that the data is realistic and adheres to the structure defined above.
    Please not thinking, just give me the result (all the foreign key tables) in the following format:
    if the table has no foreign key tables, please return "NONE"
    Example of the foreign key tables format:
    table1, table2, table3
    """
    return prompt

def generate_prompt_for_mock_data_with_values_and_fields(table_name: str, table_script: str, num_sample:int, fields_name: list[str], fields_value: list[str]) -> str:
    prompt = f"""
    You are a data generation expert. Your task is to create mock data for a database table.
    Table Name: {table_name}
    Table Structure:
    {table_script}
    Please generate {num_sample} rows of mock data for this table. Ensure that the data is realistic and adheres to the structure defined above.
    Please use the following fields and values:
    FieldsName: {fields_name}
    FieldsValue: {fields_value}
    Please not thinking, just give me the result (the SQL insert statement) in the following format:
    Example of the mock data format:
    FieldsName: column1
    FieldsValue: valueExample
    INSERT INTO {table_name} (column1, column2, column3) VALUES (valueExample, value2, value3);
    mock data only column2 and column3 that are not in the fields name and value
    column1 is the field name and valueExample is the field value
    """
    
    return prompt