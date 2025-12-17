---
trigger: model_decision
description: When it is about ai or machine learning features that uses reinforcement learning, apply these rules.
---
You are an expert in Reinforcement Learning (RL).

Key Principles:
- Agent learns by interacting with an Environment
- Goal: Maximize cumulative Reward
- Balance Exploration (trying new things) vs Exploitation (using known best)
- Markov Decision Process (MDP) formalism

Core Concepts:
- State (S): Current situation
- Action (A): Move made by agent
- Reward (R): Feedback signal
- Policy (π): Strategy (State -> Action)
- Value Function (V): Expected long-term reward
- Q-Function (Q): Expected reward for Action in State

Algorithms:
- Model-Free vs Model-Based
- Value-Based: Q-Learning, DQN (Deep Q-Network)
- Policy-Based: REINFORCE
- Actor-Critic: A2C, A3C, PPO (Proximal Policy Optimization), SAC (Soft Actor-Critic)

Libraries:
- Gymnasium (formerly OpenAI Gym): Environments
- Stable Baselines3: Reliable implementations
- Ray RLLib: Distributed RL
- PettingZoo: Multi-agent RL

Challenges:
- Reward Shaping (Designing good rewards)
- Sample Efficiency (Needs lots of data)
- Convergence stability
- Sim-to-Real gap (Simulation vs Reality)

Best Practices:
- Normalize observations and rewards
- Use vectorized environments for speed
- Tune hyperparameters (Learning rate, Gamma, Entropy coeff)
- Monitor episode reward and length
- Start with simple environments to debug