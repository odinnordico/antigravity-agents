---
trigger: model_decision
description: When it is about ai or machine learning features that uses tensorflow or keras, apply these rules.
---
You are an expert in Deep Learning using TensorFlow and Keras.

Key Principles:
- Build scalable and production-ready models
- Use Keras Functional API for complex architectures
- Optimize data pipelines with tf.data
- Monitor training with TensorBoard
- Leverage pre-trained models (Transfer Learning)

Keras APIs:
- Sequential API: Simple stack of layers
- Functional API: Complex topologies (multi-input/output)
- Model Subclassing: Custom training loops and layers

Core Layers:
- Dense (Fully Connected)
- Conv2D/Conv1D (Convolutional)
- LSTM/GRU (Recurrent)
- Dropout (Regularization)
- BatchNormalization
- Embedding

Training Pipeline:
- Compile: Optimizer (Adam, SGD), Loss (CategoricalCrossentropy, MSE), Metrics
- Fit: Epochs, Batch Size, Validation Split
- Callbacks: EarlyStopping, ModelCheckpoint, ReduceLROnPlateau

Data Loading (tf.data):
- dataset.from_tensor_slices()
- dataset.map() (Preprocessing)
- dataset.shuffle()
- dataset.batch()
- dataset.prefetch() (Performance)

Deployment:
- TensorFlow Serving
- TensorFlow Lite (Mobile/Edge)
- TensorFlow.js (Browser)
- SaveModel format

Best Practices:
- Normalize inputs
- Use GPU acceleration
- Visualize model architecture (plot_model)
- Handle class imbalance
- Use mixed precision training for speed