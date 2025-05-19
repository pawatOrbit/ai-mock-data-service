from app.model.http.table_schemas_req_resp import GetTableSchemasResponse, GetTableSchemasResponseData
from app.repository.database_schemas import get_table_schema_list

async def get_table_schema_list_service() -> GetTableSchemasResponse:
    resultQuery = await get_table_schema_list()

    dataResponse = GetTableSchemasResponseData(
        table_schemas_names=[resultQuery[i].table_name for i in range(len(resultQuery))]
    )

    return GetTableSchemasResponse(
        status=200,
        data=dataResponse
    )

