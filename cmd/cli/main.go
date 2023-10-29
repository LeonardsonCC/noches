package main

import (
	"fmt"
	"os"
	"time"

	"github.com/LeonardsonCC/noches/internal/noches"
	"github.com/LeonardsonCC/noches/internal/noches/domain"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func main() {
	noches, err := noches.NewNoches("./my-notes.nch")
	if err != nil {
		panic(err)
	}

	Execute(noches)
}

func Execute(noches noches.Noches) {
	rootCmd := &cobra.Command{
		Use:   "noches",
		Short: "Noches - Note taking app",
		Long:  `Note taking app - https://github.com/LeonardsonCC/noches`,
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
		},
	}

	addNote := &cobra.Command{
		Use:   "add [description]",
		Short: "add note",
		Run: func(_ *cobra.Command, args []string) {
			err := noches.Create(
				domain.NewNote(
					domain.WithText(args[0]),
				),
			)
			if err != nil {
				fmt.Printf("failed to create note: %v\n", err)
				return
			}

			fmt.Printf("note created successfuly\n")
		},
	}

	listNotes := &cobra.Command{
		Use:   "list",
		Short: "list notes",
		Run: func(_ *cobra.Command, _ []string) {
			n, err := noches.List()
			if err != nil {
				fmt.Printf("failed to list notes: %v\n", err)
				return
			}

			for _, note := range n {
				fmt.Printf("Note:\n")
				fmt.Printf("ID: %s\n", note.ID())
				fmt.Printf("Text: %s\n", note.Text())
				fmt.Printf("CreatedAt: %s\n", note.CreatedAt().Format(time.UnixDate))
				fmt.Printf("UpdatedAt: %s\n", note.UpdatedAt().Format(time.UnixDate))
				fmt.Printf("==============================================\n")
			}
		},
	}

	deleteNote := &cobra.Command{
		Use:   "delete [note-id]",
		Short: "delete note by id",
		Run: func(_ *cobra.Command, args []string) {
			id, err := uuid.Parse(args[0])
			if err != nil {
				fmt.Printf("invalid id: %v\n", err)
				return
			}

			err = noches.Delete(id)
			if err != nil {
				fmt.Printf("failed to delete note: %v\n", err)
				return
			}

			fmt.Printf("note %s deleted successfuly", id.String())
		},
	}

	rootCmd.AddCommand(addNote)
	rootCmd.AddCommand(listNotes)
	rootCmd.AddCommand(deleteNote)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
