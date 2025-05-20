package prompt

import (
	"fmt"
	"strings"
)

type GeneratePromptUtils interface{
	GeneratePromptWithoutKey(tableName, tableScript string, numSample int) string
	GeneratePromptForFKExtraction(tableName, tableScript string) string
	GeneratePromptForMockDataWithValues(tableName, tableScript string, numSample int, fieldsName, fieldsValue []string) string

}

type GeneratePromptUtilsImpl struct {
}

func NewGeneratePromptUtils() GeneratePromptUtils {
	return &GeneratePromptUtilsImpl{}
}

// GeneratePromptWithoutKey generates a prompt to create mock data without using functions like NOW() or UUID().
func (s GeneratePromptUtilsImpl) GeneratePromptWithoutKey(tableName, tableScript string, numSample int) string {
	return fmt.Sprintf(`
You are a data generation expert. Your task is to create mock data for a database table.
Table Name: %s
Table Structure:
%s
Please generate %d rows of mock data for this table. Ensure that the data is realistic and adheres to the structure defined above.
Create mock data for all fields in the table without using functions like NOW() or UUID(). Use the following format:
Please do not think or describe, just provide the SQL insert statements in the following format:
Example of format to use:
uuid 'a5f89c0d-e4b2-46ae-8716-11431ddad3af', 'b2e7dcfa-e1ab-460a-8a5a-f9ce555d1234'
timestamp '2023-10-01 12:00:00'
timestampz '2023-10-01 12:00:00+00'
Example of the mock data format:
INSERT INTO %s (column1, column2, column3) VALUES (value1, value2, value3);
`, tableName, tableScript, numSample, tableName)
}

// GeneratePromptForFKExtraction generates a prompt to extract foreign key tables and fields from a table script.
func (s GeneratePromptUtilsImpl) GeneratePromptForFKExtraction(tableName, tableScript string) string {
	return fmt.Sprintf(`
You are a data generation expert. Your task is to extract foreign key table names and their corresponding fields from a database table script.
Table Name: %s
Table Structure:
%s
Please extract the foreign key table names and their corresponding fields from the table script.
Do not think or describe, just provide the result in the following format:
Example of table schema format:
updated_by uuid REFERENCES users(id),
category_id uuid REFERENCES categories(id),
Example of the foreign key format:
<linked_field> updated_by, category_id </linked_field>
<foreign_key_table> users, categories </foreign_key_table>
<foreign_key_field> id, id </foreign_key_field>
`, tableName, tableScript)
}

// GeneratePromptForMockDataWithValues generates a prompt to create mock data with specified fields and values.
func (s GeneratePromptUtilsImpl) GeneratePromptForMockDataWithValues(tableName, tableScript string, numSample int, fieldsName, fieldsValue []string) string {
	fieldsNameStr := strings.Join(fieldsName, ", ")
	fieldsValueStr := strings.Join(fieldsValue, ", ")

	return fmt.Sprintf(`
You are a data generation expert. Your task is to create mock data for a database table.
Table Name: %s
Table Structure:
%s
With the following fields and values:
FieldsName: [%s]
FieldsValue: [%s]
!!! Please do not think or describe the result like "I have generated mock data..." !!!
Provide only the SQL insert statements in the following format:
Example of the mock data format:
INSERT INTO %s (column1, column2, column3) VALUES (value1, value2, value3);
If the field name is not in the provided fields name and value, generate mock data for those fields without using functions like NOW() or UUID().
Example when field name and value are provided:
Field name is column1 and value is valueExample:
INSERT INTO %s (column1, column2, column3) VALUES (valueExample, value2, value3);
(mock data only for column2 and column3 that are not in the provided fields name and value)
Please generate %d rows of mock data for this table. Ensure that the data is realistic and adheres to the structure defined above.
`, tableName, tableScript, fieldsNameStr, fieldsValueStr, tableName, tableName, numSample)
}
