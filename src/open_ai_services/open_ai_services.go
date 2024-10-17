package openaiservice

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/mzamani18/rapd_solutions_challenge/config"
	"github.com/mzamani18/rapd_solutions_challenge/entity"
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

var client *openai.Client
var ctx context.Context

func InitilizeClient() {
	if(client == nil){
		client = openai.NewClient(config.Config.OpenAI.ApiKey)
		ctx = context.Background()
	}
}


func GetLaptopDetail(text string) (*entity.LaptopDetail, error) {
	var laptop_detail entity.LaptopDetail
	schema, err := jsonschema.GenerateSchemaForType(laptop_detail)

	if err != nil {
		log.Fatalf("GenerateSchemaForType error: %v", err)
		return nil, err
	}

	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT4oMini,
		Temperature: config.Config.OpenAI.Temperature,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "Please extract these attributes of a laptop from text that I send to you, it is possible to there exist missing valuse that you must use your knowledge about laptops, note that battery status should be YES or NO",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: text,
			},
		},
		ResponseFormat: &openai.ChatCompletionResponseFormat{
			Type: openai.ChatCompletionResponseFormatTypeJSONSchema,
			JSONSchema: &openai.ChatCompletionResponseFormatJSONSchema{
				Name:   "convert_unstructured_laptopdata_to_structured",
				Schema: schema,
				Strict: true,
			},
		},
	})

	if err != nil {
		log.Fatalf("CreateChatCompletion error: %v", err)
		return nil, err
	}

	err = schema.Unmarshal(resp.Choices[0].Message.Content, &laptop_detail)
	if err != nil {
		log.Fatalf("Unmarshal schema error: %v", err)
		return nil, err
	}

	return &laptop_detail, nil
}
