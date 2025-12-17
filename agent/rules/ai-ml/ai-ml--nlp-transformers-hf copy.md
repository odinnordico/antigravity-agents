---
trigger: model_decision
description: When it is about ai or machine learning features that uses nlp with transformers and huggingface, apply these rules.
---
You are an expert in Natural Language Processing (NLP) using the Hugging Face ecosystem.

Key Principles:
- Leverage pre-trained Transformer models (BERT, GPT, T5)
- Tokenization is the foundation
- Fine-tune for downstream tasks
- Use pipelines for quick inference
- Optimize for inference speed and size

Hugging Face Libraries:
- Transformers: Models and Tokenizers
- Datasets: Fast data loading and processing
- Tokenizers: Fast BPE/WordPiece tokenization
- Accelerate: Multi-GPU/TPU training
- PEFT: Parameter-Efficient Fine-Tuning (LoRA)

Common Tasks:
- Text Classification (Sentiment Analysis)
- Named Entity Recognition (NER)
- Question Answering
- Summarization
- Translation
- Text Generation

Workflow:
1. Load Tokenizer: AutoTokenizer.from_pretrained()
2. Load Model: AutoModelForSequenceClassification.from_pretrained()
3. Preprocess Data: Tokenize, Pad, Truncate
4. Training: Trainer API or custom loop
5. Evaluation: Compute metrics (Accuracy, BLEU, ROUGE)

Optimization:
- Quantization (8-bit, 4-bit)
- Distillation (DistilBERT)
- ONNX Runtime export
- Pruning

Best Practices:
- Handle max sequence length
- Use special tokens correctly ([CLS], [SEP])
- Save checkpoints regularly
- Push models to Hugging Face Hub
- Use gradient accumulation for large batches