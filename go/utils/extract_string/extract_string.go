package extractstring

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
)

type ExtractStringUtils interface{
	ExtractForeignKeyInfo(input string) ([]string, map[string][2]string)
	ExtractInsertValues(sql string) (map[string]string, error)
}

type ExtractStringUtilsImpl struct {
}

func NewExtractStringUtils() ExtractStringUtils {
	return &ExtractStringUtilsImpl{}
}

// ExtractForeignKeyInfo parses XML-style metadata from a string and returns linked fields and a mapping
func (s ExtractStringUtilsImpl) ExtractForeignKeyInfo(input string) ([]string, map[string][2]string) {
	// Helper to extract content inside specific tags
	extractTagContent := func(tag string, input string) []string {
		re := regexp.MustCompile(fmt.Sprintf(`<%s>\s*(.*?)\s*</%s>`, tag, tag))
		matches := re.FindStringSubmatch(input)
		if len(matches) < 2 {
			return []string{}
		}
		parts := strings.Split(matches[1], ",")
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}
		return parts
	}

	linkedFields := extractTagContent("linked_field", input)
	foreignTables := extractTagContent("foreign_key_table", input)
	foreignFields := extractTagContent("foreign_key_field", input)

	fkDict := make(map[string][2]string)
	for i := range linkedFields {
		// Only map if all slices are of equal length
		if i < len(foreignTables) && i < len(foreignFields) {
			fkDict[linkedFields[i]] = [2]string{foreignTables[i], foreignFields[i]}
		}
	}

	return linkedFields, fkDict
}

// ExtractInsertValues parses SQL INSERT statement into a map of column-value pairs
func (s ExtractStringUtilsImpl) ExtractInsertValues(sql string) (map[string]string, error) {
	log.Printf("Extracting values from SQL: %s", sql)

	re := regexp.MustCompile(`(?i)INSERT INTO \w+\s*\(([^)]+)\)\s*VALUES\s*\(([^)]+)\)`)
	matches := re.FindStringSubmatch(sql)
	if len(matches) != 3 {
		return nil, errors.New("invalid INSERT statement format")
	}

	columns := strings.Split(matches[1], ",")
	valuesRaw := matches[2]

	// Extract string values including quoted strings with commas
	valueRegex := regexp.MustCompile(`'((?:[^']|\\')*)'`)
	valueMatches := valueRegex.FindAllStringSubmatch(valuesRaw, -1)

	if len(columns) != len(valueMatches) {
		return nil, errors.New("number of columns and values do not match")
	}

	result := make(map[string]string)
	for i, col := range columns {
		result[strings.TrimSpace(col)] = valueMatches[i][1]
	}

	return result, nil
}
