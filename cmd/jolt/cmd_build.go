package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/alecthomas/repr"
	"github.com/spf13/cobra"
	jolt "github.com/srand/go-jolt/pkg"
)

var cmdBuild = &cobra.Command{
	Use:   "build [task(s)]",
	Short: "Build task artifact",
	Args:  cobra.MinimumNArgs(1),
	Run:   Build,
}

func init() {
	rootCmd.AddCommand(cmdBuild)
}

func Build(cmd *cobra.Command, args []string) {
	builder := jolt.NewBuilder()
	for _, builtin := range jolt.Builtin {
		recipe, err := builder.Parse(strings.NewReader(*builtin))
		if err != nil {
			repr.Println(recipe, repr.Indent("  "), repr.OmitEmpty(true))
			panic(err)
		}
	}

	matches, err := filepath.Glob("*.jolt")
	if err != nil {
		panic("No .jolt files found in the current working directory")
	}

	for _, match := range matches {
		file, err := os.Open(match)
		if err != nil {
			panic("Failed to open recipe:" + err.Error())
		}
		defer file.Close()

		recipe, err := builder.Parse(file)
		if err != nil {
			repr.Println(recipe, repr.Indent("  "), repr.OmitEmpty(true))
			panic(err)
		}
		// repr.Println(recipe, repr.Indent("  "), repr.OmitEmpty(true))

		index := jolt.NewJobIndex()

		for _, task := range args {
			env, ok := recipe.Env[task]
			if !ok {
				panic("Task not found")
			}

			for _, task := range env.Tasks {
				for _, job := range task.Jobs {
					index.Add(job)
				}
			}

			schedule := jolt.NewJobSchedule(index)
			go schedule.Dispatch()
			<-schedule.Done()

			if !index.IsEmpty() {
				log.Fatal("Unprocessed jobs remain: ", index.RefToJob)
			}
		}

	}
}
