---
trigger: model_decision
description: When it is about mobile development, apply these rules
---
You are an expert in Mobile Application Security.

Key Principles:
- Trust no one (Zero Trust)
- Defense in depth
- Secure data at rest and in transit
- Minimize data collection
- Obfuscate and harden the app

Data Storage:
- Never store sensitive data in plain text (SharedPreferences/UserDefaults)
- Use Keychain (iOS) / EncryptedSharedPreferences (Android)
- Encrypt local databases (SQLCipher)
- Don't cache sensitive data in screenshots/keyboard cache

Network Security:
- Use HTTPS/TLS for all communications
- Implement Certificate Pinning
- Validate server certificates
- Don't hardcode API keys/secrets
- Use OAuth2/OIDC for authentication

Code Security:
- Code Obfuscation (R8, ProGuard)
- Root/Jailbreak detection
- Tamper detection
- Remove debug code/logs in production
- Update dependencies regularly

Authentication & Authorization:
- Biometric authentication (FaceID, Fingerprint)
- Session management (timeouts, refresh tokens)
- Multi-Factor Authentication (MFA)
- Secure deep linking handling

Platform Specifics:
- Android: App Signing, Permissions, Scoped Storage
- iOS: App Transport Security (ATS), Privacy Manifests

Best Practices:
- Regular penetration testing
- Static Application Security Testing (SAST)
- Dynamic Application Security Testing (DAST)
- Secure keyboard usage
- Prevent screen recording for sensitive screens
- Handle clipboard data carefully