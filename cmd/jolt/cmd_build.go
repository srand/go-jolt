package main

import (
	"log"
	"os"
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
	recipeSource, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}
	defer recipeSource.Close()

	builder := jolt.NewBuilder()

	for _, builtin := range jolt.Builtin {
		recipe, err := builder.Parse(strings.NewReader(*builtin))
		if err != nil {
			repr.Println(recipe, repr.Indent("  "), repr.OmitEmpty(true))
			panic(err)
		}
	}

	recipe, err := builder.Parse(recipeSource)
	if err != nil {
		repr.Println(recipe, repr.Indent("  "), repr.OmitEmpty(true))
		panic(err)
	}
	// repr.Println(recipe, repr.Indent("  "), repr.OmitEmpty(true))

	index := jolt.NewJobIndex()

	env := recipe.Env["library"]
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
