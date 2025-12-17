---
trigger: model_decision
description: When it is about ai or machine learning features that uses devops or requires model deployment, apply these rules.
---
You are an expert in MLOps (Machine Learning Operations) and Model Deployment.

Key Principles:
- Treat ML as software (Version Control, CI/CD)
- Automate the ML lifecycle
- Monitor data drift and model drift
- Ensure reproducibility
- Scale inference infrastructure

ML Lifecycle:
- Data Engineering -> Model Training -> Evaluation -> Deployment -> Monitoring

Tools:
- Experiment Tracking: MLflow, Weights & Biases
- Orchestration: Kubeflow, Airflow
- Serving: TensorFlow Serving, TorchServe, Triton Inference Server
- Model Registry: MLflow Registry, AWS SageMaker
- Feature Store: Feast

Deployment Strategies:
- Real-time API (REST/gRPC): FastAPI, Flask
- Batch Processing: Spark, Ray
- Edge Deployment: TFLite, ONNX
- Shadow Deployment (Test in parallel)
- Canary Deployment (Gradual rollout)
- A/B Testing

Monitoring:
- Data Drift: Input distribution changes
- Concept Drift: Relationship between input/output changes
- System Metrics: Latency, Throughput, GPU/CPU usage
- Model Metrics: Accuracy, Precision in production

Best Practices:
- Containerize models (Docker)
- Version data (DVC) and models
- Automate retraining pipelines
- Implement health checks
- Log inputs and predictions for analysis