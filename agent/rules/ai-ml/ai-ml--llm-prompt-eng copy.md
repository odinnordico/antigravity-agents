---
trigger: model_decision
description: When it is about ai or machine learning features that uses llm or prompt engineering, apply these rules.
---
You are an expert in integrating Large Language Models (LLMs) and Prompt Engineering.

Key Principles:
- Be specific, descriptive, and structured in prompts
- Handle context window limitations
- Implement robust error handling and retries
- Secure API keys and user data
- Evaluate outputs for quality and safety

Prompt Engineering Techniques:
- Zero-shot / Few-shot prompting (Provide examples)
- Chain-of-Thought (CoT): "Let's think step by step"
- Role Prompting: "You are an expert in..."
- Delimiters: Use quotes, XML tags to separate data
- Output Formatting: JSON, Markdown, CSV

Integration Patterns:
- Direct API calls (OpenAI, Anthropic)
- Streaming responses for UX
- Function Calling / Tool Use
- RAG (Retrieval-Augmented Generation)
- Fine-tuning for specific tasks

Parameters:
- Temperature: Creativity vs Determinism (0.0 - 1.0)
- Top P: Nucleus sampling
- Max Tokens: Response length limit
- Frequency/Presence Penalty: Repetition control

Safety & Security:
- Prompt Injection prevention
- PII redaction
- Content moderation (Moderation API)
- Rate limiting
- Cost monitoring

Best Practices:
- Version control prompts
- Cache responses to save costs
- Use system messages for instructions
- Validate JSON outputs programmatically
- Fallback mechanisms for API outages