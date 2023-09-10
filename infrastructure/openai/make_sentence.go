package openai

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"math"
	"strings"
)

func (c *Client) makeJapaneseSentence(rememberVocabularyList []string, vocabulary string) (japaneseOriginSentence, japaneseHiraganaSentence, traditionalChineseSentence string, err error) {
	resp, err := c.openai.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleSystem,
					Content: "您現在是一個日文老師\n" +
						"我會提供你一個日文的單詞\n" +
						"要請你為一個JLPT程度為N5的同學造句，並在漢字後面附上相對應的平假名(請不要隨意的拆解我提供的單詞)\n" +
						"也在造句的後面附上繁體中文的翻譯\n\n" +
						"另外，最後我會附上曾經背過的日文單詞，在接下來的造句中，請盡可能的使用這些單詞\n\n" +
						"接下來的所有回應請用日文來進行\n\n" +
						"以下為曾經背過的日文單詞 \n\n" +
						strings.Join(rememberVocabularyList, ","),
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "みかん",
				},
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: "彼女はみかんを食べています。\n彼女[かのじょ]はみかんを 食[た]べています。\n她正在吃橘子。",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "机",
				},
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: "私の机は木製です。\n私[わたし]の 机[つくえ]は 木製[もくせい]です。\n我的桌子是木制的。",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "パソコン",
				},
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: "私はパソコンで日本語を勉強します。\n私[わたし]はパソコンで日本語[にほんご]を勉強[べんきょう]します。\n我用電腦學習日語。",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "携帯",
				},
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: "彼は携帯で友達とメッセージを送っています。\n彼[かれ]は携帯[けいたい]で友達[ともだち]とメッセージを送[おく]っています。\n他正在用手機和朋友發送訊息。",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: vocabulary,
				},
			},
			Temperature: math.SmallestNonzeroFloat32,
		},
	)
	result := resp.Choices[0].Message.Content
	splitResult := strings.Split(result, "\n")
	japaneseOriginSentence = splitResult[0]
	japaneseHiraganaSentence = splitResult[1]
	traditionalChineseSentence = splitResult[2]
	return japaneseOriginSentence, japaneseHiraganaSentence, traditionalChineseSentence, err
}
