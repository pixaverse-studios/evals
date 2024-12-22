# Pixa Evals: Model Evaluation Framework for Child-Safe AI

An advanced benchmarking system designed to assess and compare AI language models for child-friendly interactions, ensuring safe and developmentally appropriate digital experiences.

## Overview

This framework provides comprehensive testing to evaluate AI models on their ability to interact safely and effectively with children. Using OpenAI's o1-preview as the evaluation engine, it analyzes models across five key dimensions:

- 🧠 Intelligence (IQ) - Assesses knowledge, reasoning, and problem-solving abilities
- ❤️ Emotional Intelligence (EQ) - Measures empathy, social awareness, and emotional understanding  
- 📚 Educational Value - Evaluates learning outcomes and teaching effectiveness
- 🎨 Creativity - Rates imaginative thinking and engaging interactions
- 🛡️ Safety & Privacy - Validates content filtering and child protection measures

## Key Features

- Real-time parallel evaluation of multiple AI models
- Context-aware testing incorporating:
  - Age-appropriate content and language
  - Child development milestones
  - Individual interests and learning styles
  - Educational objectives
- Detailed performance analytics and reporting
- Easy-to-use command line interface
- Flexible scoring system with both objective and subjective metrics
- Customizable evaluation parameters

## Latest Benchmark Results (v1)

| Model         | Intelligence | Emotional Intelligence | Educational Value | Creativity | Safety & Privacy |
|--------------|--------------|----------------------|------------------|------------|-----------------|
| OPENAI GPT-4o | 19/40       | 24/40                | 24/40           | 18/40     | 27/40          |
| Sonnet3.5     | 20/40       | 36/40                | 23/40           | 23.5/40   | 33/40          |
| Qwen2.5-70b   | 24/40       | 31/40                | 15/40           | 23/40     | 31/40          |
| Llama3.2-70b  | Pending     | Pending              | Pending         | Pending    | Pending        |
| Nemotron-70b  | 20/40       | 34/40                | 10/40           | 25/40     | 31/40          |
| Mixtral-8x22B | 22/40       | 29/40                | 17/40           | 21/40     | 30/40          |
| Qwen-QwQ      | 23/40       | 19/40                | 14/40           | 25/40     | 28/40          |
| Gemini        | Coming Soon | Coming Soon          | Coming Soon     | Coming Soon| Coming Soon    |


## Roadmap

- Standalone Conversation Benchmarks
  - Evaluate fabricated child-AI dialogues in isolation
  - Assess conversation consistency and coherence
  - Measure contextual memory and relationship building
  - Score natural dialogue flow and age-appropriate responses
  - Analyze conversation safety across multiple exchanges
  - Compare models on extended interaction capabilities
