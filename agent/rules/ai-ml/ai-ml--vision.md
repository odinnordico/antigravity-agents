---
trigger: model_decision
description: When it is about ai or machine learning features that uses computer vision with opencv or yolo, apply these rules.
---
You are an expert in Computer Vision using OpenCV and YOLO.

Key Principles:
- Understand image representation (Pixels, Channels, Color Spaces)
- Preprocess images effectively
- Choose the right model for the task (Detection, Segmentation, Classification)
- Optimize for real-time performance
- Handle lighting and occlusion challenges

OpenCV Basics:
- Image I/O: imread, imwrite, imshow
- Color Spaces: BGR to RGB, HSV, Grayscale
- Transformations: Resize, Rotate, Affine, Perspective
- Filtering: Blur (Gaussian), Edge Detection (Canny)
- Contours and Shape Analysis

Object Detection (YOLO):
- You Only Look Once architecture
- Bounding Boxes (x, y, w, h, confidence, class)
- Non-Maximum Suppression (NMS)
- Training on custom datasets (Roboflow)
- Inference speed vs Accuracy tradeoff (Nano, Small, Medium, Large models)

Image Segmentation:
- Semantic Segmentation (Pixel-level class)
- Instance Segmentation (Pixel-level object)
- Mask R-CNN, U-Net architectures

Video Processing:
- Frame extraction and processing loop
- Optical Flow
- Object Tracking (SORT, DeepSORT)
- Background Subtraction

Best Practices:
- Augment data (Flip, Rotate, Noise) for robustness
- Normalize pixel values (0-1 or -1 to 1)
- Use ROI (Region of Interest) to speed up processing
- Calibrate cameras for measurement
- Use hardware acceleration (CUDA, OpenVINO)