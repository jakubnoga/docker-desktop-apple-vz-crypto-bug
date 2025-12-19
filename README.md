# Bug reproduction environment

<img width="178" height="103" alt="image" src="https://github.com/user-attachments/assets/9ddea329-9946-4963-93c0-db5e5498f572" />
<img width="976" height="645" alt="image" src="https://github.com/user-attachments/assets/eb26ca0a-2b8b-4a1a-9902-a09e64dbc4d6" />

# Running

```
docker build -t crypto-bug .
docker run crypto-bug
```

Expect output:

```
=== Apple Virtualization Framework Bug ===
XChaCha20Poly1305 decryption fails at >321 bytes

Under threshold (319 bytes): PASSED (319 bytes)
At threshold (321 bytes): FAILED (321 bytes)
```

Then switch off Rosetta, and expect output:

```
=== Apple Virtualization Framework Bug ===
XChaCha20Poly1305 decryption fails at >321 bytes

Under threshold (319 bytes): PASSED (319 bytes)
At threshold (321 bytes): PASSED (321 bytes)
```
