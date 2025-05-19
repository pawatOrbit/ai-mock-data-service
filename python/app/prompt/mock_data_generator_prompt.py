def generate_prompt_without_key(table_name: str, table_script: str, num_sample:int ) -> str:
    prompt = f"""
    You are a data generation expert. Your task is to create mock data for a database table.
    Table Name: {table_name}
    Table Structure:
    {table_script}
    Please generate {num_sample} rows of mock data for this table. Ensure that the data is realistic and adheres to the structure defined above.
    create mock data for all fields in the table not use function like NOW() or UUID() using the following format:
    Please not thinking, just give me the result (the SQL insert statement) in the following format:
    Please give me only the SQL insert statement!
    Example of format this use function:
    uuid 'a5f89c0d-e4b2-46ae-8716-11431ddad3af', 'b2e7dcfa-e1ab-460a-8a5a-f9ce555d1234'
    timestamp '2023-10-01 12:00:00' 
    timestampz '2023-10-01 12:00:00+00'
    Example of the mock data format:
    INSERT INTO {table_name} (column1, column2, column3) VALUES (value1, value2, value3);
    """
    
    return prompt


def generate_prompt_for_get_fk_table_and_fk_field_and_linked_field_from_table_script(table_name: str, table_script: str) -> str:
    prompt = f"""
    You are a data generation expert. Your task is to extract foreign key table names and their corresponding fields from a database table script.
    Table Name: {table_name}
    Table Structure:
    {table_script}
    Please extract the foreign key table names and their corresponding fields from the table script.
    Please not thinking, just give me the result (the foreign key table names and fields) in the following format:
    Example of table schema format:
    updated_by uuid REFERENCES users(id),
    category_id uuid REFERENCES categories(id),
    Example of the foreign key format:
    <linked_field> updated_by, category_id </linked_field>
    <foreign_key_table> users, categories </foreign_key_table>
    <foreign_key_field> id, id </foreign_key_field>
    """
    return prompt

def generate_prompt_for_mock_data_with_values_and_fields(table_name: str, table_script: str, num_sample:int, fields_name: list[str], fields_value: list[str]) -> str:
    prompt = f"""
    You are a data generation expert. Your task is to create mock data for a database table.
    Table Name: {table_name}
    Table Structure:
    {table_script}
    with the following fields and values:
    FieldsName: {fields_name}
    FieldsValue: {fields_value}
    !!! Please not thinking, and not describe of result like I have generated mock data ... !!!
    give me the result (the SQL insert statement) in the following format:
    Please give me only the SQL insert statement!
    Example of the mock data format:
    INSERT INTO {table_name} (column1, column2, column3) VALUES (value1, value2, value3);
    if the field name is not in the fields name and value
    Please generate mock data for all fields in the table not use function like NOW() or UUID() using the following format:
    but the field name and value is not in the fields name and value
    example of the mock data format that have the field name and value
    field name is column1 and value is valueExample
    INSEERT INTO {table_name} (column1, column2, column3) VALUES (valueExample, value2, value3);
    (mock data only column2 and column3 that are not in the fields name and value)
    Please generate {num_sample} rows of mock data for this table. Ensure that the data is realistic and adheres to the structure defined above.
    """
    
    return prompt