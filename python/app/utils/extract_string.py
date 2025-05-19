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

def extract_foreign_key_info(input_str: str) -> tuple[list[str], dict[str, str]]:
    """
    Extracts foreign key mapping from the input string.
    Example:
      <linked_field> updated_by, category_id </linked_field>
      <foreign_key_table> users, categories </foreign_key_table>
      <foreign_key_field> id, id </foreign_key_field>
    Returns:
      fields: list of linked fields
      fk_dict: dict mapping each linked field to (foreign_key_table, foreign_key_field)
    """
    import re

    matches = re.findall(r"<([^>]+)>([^<]+)</\1>", input_str)
    if len(matches) != 3:
        raise ValueError("Input must contain exactly three tag pairs.")

    fields = [f.strip() for f in matches[0][1].split(",")]
    tables = [t.strip() for t in matches[1][1].split(",")]
    fk_fields = [fk.strip() for fk in matches[2][1].split(",")]

    if not (len(fields) == len(tables) == len(fk_fields)):
        raise ValueError("Number of fields, tables, and foreign key fields must match.")

    return fields, dict(zip(tables, fk_fields))

def extract_insert_values(insert_sql: str) -> dict:
    """
    Extracts column-value pairs from a single-line SQL INSERT statement.
    Example:
      INSERT INTO tasks (id, title, description, category_id) VALUES ('a5f89c0d-e4b2-46ae-8716-11431ddad3af', 'Design User Interface', 'Create mockups for the new user interface', 'b2e7dcfa-e1ab-460a-8a5a-f9ce555d1234');
      ->
      {
        "id": "a5f89c0d-e4b2-46ae-8716-11431ddad3af",
        "title": "Design User Interface",
        ...
      }
    """
    logging.info(f"Extracting values from SQL: {insert_sql}")

    # Extract columns and values
    match = re.search(
        r"INSERT INTO \w+\s*\(([^)]+)\)\s*VALUES\s*\(([^)]+)\)", insert_sql, re.IGNORECASE
    )
    if not match:
        raise ValueError("Invalid INSERT statement format.")

    columns = [col.strip() for col in match.group(1).split(",")]
    # Split values, handling quoted strings with commas
    values = re.findall(r"'((?:[^']|\\')*)'", match.group(2))

    print(f"Columns: {columns}")
    print(f"Values: {values}")

    if len(columns) != len(values):
        raise ValueError("Number of columns and values do not match.")

    return dict(zip(columns, values))