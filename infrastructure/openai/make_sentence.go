package openai

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"math"
	"strings"
)

func (c *Client) MakeJapaneseSentence(rememberVocabularyList []string, vocabulary, meaning string) (japaneseOriginSentence, japaneseHiraganaSentence, traditionalChineseSentence string, err error) {
	resp, err := c.openai.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleSystem,
					Content: `
接下來的輸入輸出內容我用三個米字號包起來

您現在是一個日文老師
我會提供你一個日文的單詞，平假名以及他的中文意思
輸入
***
両親[りょうしん]
雙親
***
要請你為一個JLPT程度為N5的同學造句，並在漢字後面附上相對應的平假名
附上漢字的方式請依照以下格式
在漢字前面加上一個半形空白，在漢字後面用中括號將平假名填入

漢字:
***山***
平假名: 
***やま***

輸出:
*** 山[やま]***


完整句子
***私は山でハイキングを楽しんでいます***
輸出
*** 私[わたし]は 山[やま]でハイキングを 楽[たの]しんでいます***

也在造句的後面附上繁體中文的翻譯

另外，最後我會附上曾經背過的日文單詞，在接下來的造句中，請盡可能的使用這些單詞

接下來的所有回應請用日文來進行

以下為曾經背過的日文單詞
` + strings.Join(rememberVocabularyList, ","),
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "両親[りょうしん]\n雙親",
				},
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: "私[わたし]の 両親[りょうしん]は 旅行[りょこう]に 行[い]っています。\n私の両親は旅行に行っています。\n我的父母正在旅行中。",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "月餅[げっぺい]\n月餅",
				},
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: "友達[ともだち]と 一緒[いっしょ]に 月餅[げっぺい]を 食[た]べました。\n友達と一緒に月餅を食べました。\n我和朋友一起吃月餅。",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "彼の方[あのかた]\nあの人　的敬語 他 她 那個人",
				},
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: " 彼の方[あのかた]が 日本語[にほんご]が 上手[じょうず]です。\n彼の方が日本語が上手です。\n他的日語比較流利。",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "八つ[やっつ]\n八個",
				},
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: " 私[わたし]の 妹[いもうと]は 八つ[やっつ]の 時[とき]にピアノを 始[はじ]めました。\n私の妹は八つの時にピアノを始めました。\n我的妹妹八歲的時候開始學鋼琴。",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "高速[こうそく]バス\n高速巴士",
				},
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: " 私[わたし]は 高速[こうそく]バスで 東京[とうきょう]へ 行[い]きます。\n私は高速バスで東京へ行きます。\n我會坐高速巴士去東京。",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "読[よ]み 方[かた]\n讀法 念法",
				},
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: "この 漢字[かんじ]の 読[よ]み 方[かた]を 教[おし]えてください。\nこの漢字の読み方を教えてください。\n請告訴我這個漢字的讀法。",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "砂糖[さとう]\n砂糖",
				},
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: "お 茶[ちゃ]に 砂糖[さとう]を 入[い]れると 甘[あま]くなります。\nお茶に砂糖を入れると甘くなります。\n在茶裡加糖會變甜。",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "エアコン\n冷氣",
				},
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: "夏[なつ]になると エアコンをつけて 部屋[へや]を 冷[さ]やします。\n夏になるとエアコンをつけて部屋を冷やします。\n夏天的時候，我會開冷氣來降溫房間。",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "電気[でんき]\n電燈",
				},
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: "夜[よる]になると 電気[でんき]をつけて 部屋[へや]を明[あか]るくします。\n夜になると電気をつけて部屋を明るくします。\n晚上的時候，我會開燈讓房間變亮。",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "間[あいだ]\n兩者之間",
				},
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: "私[わたし]と 友達[ともだち]の 間[あいだ]には 深[ふか]い 絆[きずな]があります。\n私と友達の間には深い絆があります。\n我和朋友之間有著深厚的羈絆。",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("%s\n%s", vocabulary, meaning),
				},
			},
			Temperature: math.SmallestNonzeroFloat32,
		},
	)
	result := resp.Choices[0].Message.Content
	splitResult := strings.Split(result, "\n")
	japaneseHiraganaSentence = splitResult[0]
	japaneseOriginSentence = splitResult[1]
	traditionalChineseSentence = splitResult[2]
	return japaneseOriginSentence, japaneseHiraganaSentence, traditionalChineseSentence, err
}
