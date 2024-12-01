package main

import (
	"evals/internal/models"
	"evals/pkg/eval"
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/joho/godotenv"
	"os"
)

var CLI struct {
	IQ        IQCmd        `cmd:"" help:"Run IQ benchmark evaluation"`
	EQ        EQCmd        `cmd:"" help:"Run EQ benchmark evaluation"` 
	Learning  LearningCmd  `cmd:"" help:"Run learning benchmark evaluation"`
	Safety    SafetyCmd    `cmd:"" help:"Run safety benchmark evaluation"`
	Curiosity CuriosityCmd `cmd:"" help:"Run curiosity benchmark evaluation"`
	All       AllCmd       `cmd:"" help:"Run all benchmark evaluations"`
}

type IQCmd struct {
	Model     string `help:"Model to evaluate (OpenAI, Claude, Llama, Qwen, Nemotron, Gemini, Phi, Mistral)" default:"OpenAI"`
	Name      string `help:"Child's name"`
	Age       string `help:"Child's age"`
	Interests string `help:"Child's interests"`
	Goals     string `help:"Child's learning goals"`
}

type EQCmd struct {
	Model     string `help:"Model to evaluate (OpenAI, Claude, Llama, Qwen, Nemotron, Gemini, Phi, Mistral)" default:"OpenAI"`
	Name      string `help:"Child's name"`
	Age       string `help:"Child's age"`
	Interests string `help:"Child's interests"`
	Goals     string `help:"Child's learning goals"`
}

type LearningCmd struct {
	Model     string `help:"Model to evaluate (OpenAI, Claude, Llama, Qwen, Nemotron, Gemini, Phi, Mistral)" default:"OpenAI"`
	Name      string `help:"Child's name"`
	Age       string `help:"Child's age"`
	Interests string `help:"Child's interests"`
	Goals     string `help:"Child's learning goals"`
}

type SafetyCmd struct {
	Model     string `help:"Model to evaluate (OpenAI, Claude, Llama, Qwen, Nemotron, Gemini, Phi, Mistral)" default:"OpenAI"`
	Name      string `help:"Child's name"`
	Age       string `help:"Child's age"`
	Interests string `help:"Child's interests"`
	Goals     string `help:"Child's learning goals"`
}

type CuriosityCmd struct {
	Model     string `help:"Model to evaluate (OpenAI, Claude, Llama, Qwen, Nemotron, Gemini, Phi, Mistral)" default:"OpenAI"`
	Name      string `help:"Child's name"`
	Age       string `help:"Child's age"`
	Interests string `help:"Child's interests"`
	Goals     string `help:"Child's learning goals"`
}

type AllCmd struct {
	Model     string `help:"Model to evaluate (OpenAI, Claude, Llama, Qwen, Nemotron, Gemini, Phi, Mistral)" default:"OpenAI"`
	Name      string `help:"Child's name"`
	Age       string `help:"Child's age"`
	Interests string `help:"Child's interests"`
	Goals     string `help:"Child's learning goals"`
}

func (r *IQCmd) Run() error {
	model := models.FromString(r.Model)
	childData := eval.ChildData{
		Name:      r.Name,
		Age:       r.Age,
		Interests: r.Interests,
		Goals:     r.Goals,
	}
	client := eval.Engine(model, childData)
	score, msg, err := client.EvaluateIqBenchmarks()
	if err != nil {
		return err
	}
	fmt.Printf("Score: %s\nMessage: %s\n", score, msg)
	return nil
}

func (r *EQCmd) Run() error {
	model := models.FromString(r.Model)
	childData := eval.ChildData{
		Name:      r.Name,
		Age:       r.Age,
		Interests: r.Interests,
		Goals:     r.Goals,
	}
	client := eval.Engine(model, childData)
	score, msg, err := client.EvaluateEqBenchmarks()
	if err != nil {
		return err
	}
	fmt.Printf("Score: %s\nMessage: %s\n", score, msg)
	return nil
}

func (r *LearningCmd) Run() error {
	model := models.FromString(r.Model)
	childData := eval.ChildData{
		Name:      r.Name,
		Age:       r.Age,
		Interests: r.Interests,
		Goals:     r.Goals,
	}
	client := eval.Engine(model, childData)
	score, msg, err := client.EvaluateLearningBenchmarks()
	if err != nil {
		return err
	}
	fmt.Printf("Score: %s\nMessage: %s\n", score, msg)
	return nil
}

func (r *SafetyCmd) Run() error {
	model := models.FromString(r.Model)
	childData := eval.ChildData{
		Name:      r.Name,
		Age:       r.Age,
		Interests: r.Interests,
		Goals:     r.Goals,
	}
	client := eval.Engine(model, childData)
	score, msg, err := client.EvaluateSafetyBenchmarks()
	if err != nil {
		return err
	}
	fmt.Printf("Score: %s\nMessage: %s\n", score, msg)
	return nil
}

func (r *CuriosityCmd) Run() error {
	model := models.FromString(r.Model)
	childData := eval.ChildData{
		Name:      r.Name,
		Age:       r.Age,
		Interests: r.Interests,
		Goals:     r.Goals,
	}
	client := eval.Engine(model, childData)
	score, msg, err := client.EvaluateCuriosityBenchmarks()
	if err != nil {
		return err
	}
	fmt.Printf("Score: %s\nMessage: %s\n", score, msg)
	return nil
}

func (r *AllCmd) Run() error {
	model := models.FromString(r.Model)
	childData := eval.ChildData{
		Name:      r.Name,
		Age:       r.Age,
		Interests: r.Interests,
		Goals:     r.Goals,
	}
	client := eval.Engine(model, childData)

	fmt.Println("\n=== IQ Evaluation ===")
	iqScore, iqMsg, err := client.EvaluateIqBenchmarks()
	if err != nil {
		return err
	}
	fmt.Printf("Score: %s\nMessage: %s\n", iqScore, iqMsg)

	fmt.Println("\n=== EQ Evaluation ===")
	eqScore, eqMsg, err := client.EvaluateEqBenchmarks()
	if err != nil {
		return err
	}
	fmt.Printf("Score: %s\nMessage: %s\n", eqScore, eqMsg)

	fmt.Println("\n=== Learning Evaluation ===")
	learnScore, learnMsg, err := client.EvaluateLearningBenchmarks()
	if err != nil {
		return err
	}
	fmt.Printf("Score: %s\nMessage: %s\n", learnScore, learnMsg)

	fmt.Println("\n=== Safety Evaluation ===")
	safetyScore, safetyMsg, err := client.EvaluateSafetyBenchmarks()
	if err != nil {
		return err
	}
	fmt.Printf("Score: %s\nMessage: %s\n", safetyScore, safetyMsg)

	fmt.Println("\n=== Curiosity Evaluation ===")
	curiosityScore, curiosityMsg, err := client.EvaluateCuriosityBenchmarks()
	if err != nil {
		return err
	}
	fmt.Printf("Score: %s\nMessage: %s\n", curiosityScore, curiosityMsg)

	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Warning: Error loading .env file: %v\n", err)
	}

	ctx := kong.Parse(&CLI,
		kong.Name("evals"),
		kong.Description("A CLI tool for evaluating and comparing AI models"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: true,
		}))
	err = ctx.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}