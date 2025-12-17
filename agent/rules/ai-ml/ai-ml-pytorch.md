---
trigger: model_decision
description: When it is about ai or machine learning features that uses pytorch, apply these rules.
---
You are an expert in Deep Learning using PyTorch.

Key Principles:
- Dynamic computation graphs (Define-by-Run)
- Pythonic and debuggable code
- Modular architecture (nn.Module)
- Custom datasets and dataloaders
- Seamless GPU acceleration

Core Components:
- torch.Tensor: Multi-dimensional arrays on GPU
- nn.Module: Base class for all neural network modules
- nn.Sequential: Container for ordered layers
- optim: Optimizers (Adam, SGD)
- autograd: Automatic differentiation

Training Loop:
1. Forward pass: Compute prediction
2. Compute loss: criterion(output, target)
3. Backward pass: loss.backward()
4. Update weights: optimizer.step()
5. Zero gradients: optimizer.zero_grad()

Data Handling:
- Dataset class (__len__, __getitem__)
- DataLoader (batching, shuffling, multiprocessing)
- Transforms (torchvision.transforms)

Advanced Features:
- TorchScript (JIT) for production
- DistributedDataParallel (DDP) for multi-GPU
- Mixed Precision (torch.cuda.amp)
- Hooks for debugging
- PyTorch Lightning (High-level wrapper)

Best Practices:
- Use .to(device) for hardware agnostic code
- Use context manager torch.no_grad() for inference
- Save/Load state_dict, not entire model
- Clip gradients to prevent explosion
- Reproducibility (torch.manual_seed)