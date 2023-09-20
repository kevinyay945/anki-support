/*
Copyright © 2023 Kevin Chen
*/
package cmd

import (
	"anki-support/helper"
	"anki-support/infrastructure/anki"
	"anki-support/infrastructure/gcp"
	"anki-support/infrastructure/openai"
	"fmt"
	"github.com/atselvan/ankiconnect"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"path/filepath"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ankiClient := anki.NewClient(log.New())
		//deckName := "製作中日語卡片"
		deckName := "日語單字"
		noteList, err := ankiClient.GetTodoNoteFromDeck(deckName)
		if err != nil {
			fmt.Printf("GetTodoNoteFromDeck, %s", err.Error())
			return
		}
		allNoteList, err := ankiClient.GetTodoNoteFromDeck(deckName)
		if err != nil {
			fmt.Printf("GetTodoNoteFromDeck, %s", err.Error())
			return
		}
		allVocabularyList := []string{}
		for _, note := range allNoteList {
			allVocabularyList = append(allVocabularyList, note.Fields["Expression"].Value)
		}

		for _, note := range noteList {
			switch note.ModelName {
			case "Japanese (recognition&recall)":
				gcpClient := gcp.NewGCP()
				defer gcpClient.Close()
				fmt.Println(note.ModelName)
				expression := note.Fields["Expression"].Value
				reading := note.Fields["Reading"].Value
				meaning := note.Fields["Meaning"].Value
				path, _ := helper.GetCurrentExecutableFolderPath()
				tempAssetPath := filepath.Join(path, "temp")
				expressionVoicePath, _ := gcpClient.GenerateAudioByText(expression, tempAssetPath, expression)

				openAIClient := openai.NewClient()
				sentence, hiraganaSentence, chineseSentence, err := openAIClient.MakeJapaneseSentence(allVocabularyList, reading, meaning)
				if err != nil {
					continue
				}
				sentenceVoicePath, _ := gcpClient.GenerateAudioByText(sentence, tempAssetPath, sentence)
				note.Fields["JapaneseNote"] = ankiconnect.FieldData{
					Value: hiraganaSentence,
					Order: note.Fields["JapaneseNote"].Order,
				}
				note.Fields["ChineseNote"] = ankiconnect.FieldData{
					Value: chineseSentence,
					Order: note.Fields["ChineseNote"].Order,
				}

				_ = ankiClient.EditNoteById(note, []ankiconnect.Audio{{
					Path:     expressionVoicePath,
					Filename: filepath.Base(expressionVoicePath),
					Fields:   []string{"Sound"},
				}, {
					Path:     sentenceVoicePath,
					Filename: filepath.Base(sentenceVoicePath),
					Fields:   []string{"JapaneseSound"},
				}}, nil, nil)
				_ = ankiClient.DeleteTagFromNote(note.NoteId, "anki-helper-vocabulary-todo")
				_ = ankiClient.AddTagFromNote(note.NoteId, "anki-helper-vocabulary-done")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
