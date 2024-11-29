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
	IQ       IQCmd       `cmd:"" help:"Run IQ benchmark evaluation"`
	EQ       EQCmd       `cmd:"" help:"Run EQ benchmark evaluation"` 
	Learning LearningCmd `cmd:"" help:"Run learning benchmark evaluation"`
	Safety   SafetyCmd   `cmd:"" help:"Run safety benchmark evaluation"`
	Curiosity CuriosityCmd `cmd:"" help:"Run curiosity benchmark evaluation"`
}

type IQCmd struct {
	Model1    string `help:"First model to evaluate (default: OpenAI)" default:"OpenAI"`
	Model2    string `help:"Second model to evaluate (default: Claude)" default:"Claude"`
	Name      string `help:"Child's name"`
	Age       string `help:"Child's age"`
	Interests string `help:"Child's interests"`
	Goals     string `help:"Child's learning goals"`
}

type EQCmd struct {
	Model1    string `help:"First model to evaluate (default: OpenAI)" default:"OpenAI"`
	Model2    string `help:"Second model to evaluate (default: Claude)" default:"Claude"`
	Name      string `help:"Child's name"`
	Age       string `help:"Child's age"`
	Interests string `help:"Child's interests"`
	Goals     string `help:"Child's learning goals"`
}

type LearningCmd struct {
	Model1    string `help:"First model to evaluate (default: OpenAI)" default:"OpenAI"`
	Model2    string `help:"Second model to evaluate (default: Claude)" default:"Claude"`
	Name      string `help:"Child's name"`
	Age       string `help:"Child's age"`
	Interests string `help:"Child's interests"`
	Goals     string `help:"Child's learning goals"`
}

type SafetyCmd struct {
	Model1    string `help:"First model to evaluate (default: OpenAI)" default:"OpenAI"`
	Model2    string `help:"Second model to evaluate (default: Claude)" default:"Claude"`
	Name      string `help:"Child's name"`
	Age       string `help:"Child's age"`
	Interests string `help:"Child's interests"`
	Goals     string `help:"Child's learning goals"`
}

type CuriosityCmd struct {
	Model1    string `help:"First model to evaluate (default: OpenAI)" default:"OpenAI"`
	Model2    string `help:"Second model to evaluate (default: Claude)" default:"Claude"`
	Name      string `help:"Child's name"`
	Age       string `help:"Child's age"`
	Interests string `help:"Child's interests"`
	Goals     string `help:"Child's learning goals"`
}

func (r *IQCmd) Run() error {
	model1 := models.FromString(r.Model1)
	model2 := models.FromString(r.Model2)
	childData := eval.ChildData{
		Name:      r.Name,
		Age:       r.Age,
		Interests: r.Interests,
		Goals:     r.Goals,
	}
	client := eval.Engine(model1, model2, childData)
	score1, score2, msg, err := client.EvaluateIqBenchmarks()
	if err != nil {
		return err
	}
	fmt.Printf("Model 1 Score: %s\nModel 2 Score: %s\nMessage: %s\n", score1, score2, msg)
	return nil
}

func (r *EQCmd) Run() error {
	model1 := models.FromString(r.Model1)
	model2 := models.FromString(r.Model2)
	childData := eval.ChildData{
		Name:      r.Name,
		Age:       r.Age,
		Interests: r.Interests,
		Goals:     r.Goals,
	}
	client := eval.Engine(model1, model2, childData)
	score1, score2, msg, err := client.EvaluateEqBenchmarks()
	if err != nil {
		return err
	}
	fmt.Printf("Model 1 Score: %s\nModel 2 Score: %s\nMessage: %s\n", score1, score2, msg)
	return nil
}

func (r *LearningCmd) Run() error {
	model1 := models.FromString(r.Model1)
	model2 := models.FromString(r.Model2)
	childData := eval.ChildData{
		Name:      r.Name,
		Age:       r.Age,
		Interests: r.Interests,
		Goals:     r.Goals,
	}
	client := eval.Engine(model1, model2, childData)
	score1, score2, msg, err := client.EvaluateLearningBenchmarks()
	if err != nil {
		return err
	}
	fmt.Printf("Model 1 Score: %s\nModel 2 Score: %s\nMessage: %s\n", score1, score2, msg)
	return nil
}

func (r *SafetyCmd) Run() error {
	model1 := models.FromString(r.Model1)
	model2 := models.FromString(r.Model2)
	childData := eval.ChildData{
		Name:      r.Name,
		Age:       r.Age,
		Interests: r.Interests,
		Goals:     r.Goals,
	}
	client := eval.Engine(model1, model2, childData)
	score1, score2, msg, err := client.EvaluateSafetyBenchmarks()
	if err != nil {
		return err
	}
	fmt.Printf("Model 1 Score: %s\nModel 2 Score: %s\nMessage: %s\n", score1, score2, msg)
	return nil
}

func (r *CuriosityCmd) Run() error {
	model1 := models.FromString(r.Model1)
	model2 := models.FromString(r.Model2)
	childData := eval.ChildData{
		Name:      r.Name,
		Age:       r.Age,
		Interests: r.Interests,
		Goals:     r.Goals,
	}
	client := eval.Engine(model1, model2, childData)
	score1, score2, msg, err := client.EvaluateCuriosityBenchmarks()
	if err != nil {
		return err
	}
	fmt.Printf("Model 1 Score: %s\nModel 2 Score: %s\nMessage: %s\n", score1, score2, msg)
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