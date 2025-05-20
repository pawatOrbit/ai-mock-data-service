package utils

import (
	extractstring "github.com/pawatOrbit/ai-mock-data-service/go/utils/extract_string"
	"github.com/pawatOrbit/ai-mock-data-service/go/utils/prompt"
)

type Utils struct {
	ExtractStringUtils  extractstring.ExtractStringUtils
	GeneratePromptUtils prompt.GeneratePromptUtils
}

func NewUtils() *Utils {
	return &Utils{
		ExtractStringUtils:  extractstring.NewExtractStringUtils(),
		GeneratePromptUtils: prompt.NewGeneratePromptUtils(),
	}
}
