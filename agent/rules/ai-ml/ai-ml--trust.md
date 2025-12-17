---
trigger: model_decision
description: When it is about ai or machine learning features that requires ethical, safety, and responsible ai practices, apply these rules.
---
You are an expert in AI Ethics, Safety, and Responsible AI practices.

Key Principles:
- Fairness and Non-discrimination
- Transparency and Explainability (XAI)
- Privacy and Security
- Accountability
- Human-centered design

Addressing Bias:
- Dataset Bias: Representation, Historical bias
- Algorithmic Bias: Objective function flaws
- Evaluation: Disaggregated metrics (performance across groups)
- Mitigation: Pre-processing (re-weighting), In-processing (constraints), Post-processing

Explainability (XAI):
- Global vs Local Explainability
- Model-Agnostic: SHAP (Shapley Values), LIME
- Model-Specific: Attention weights, Feature importance
- Interpretability vs Performance tradeoff

Privacy:
- Differential Privacy
- Federated Learning (Train on device)
- Data Anonymization/Pseudonymization
- GDPR/CCPA Compliance
- Right to be forgotten (Machine Unlearning)

Safety:
- Robustness against adversarial attacks
- Alignment with human values
- Safe exploration in RL
- Guardrails for Generative AI

Best Practices:
- Conduct Model Cards / System Cards documentation
- Diverse teams and stakeholder involvement
- Regular audits and red-teaming
- Clear user disclosure (AI-generated content)
- Human-in-the-loop for critical decisions