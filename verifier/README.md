# Verifier

## Usage

`verify config.yaml ITERATION SLEEP_MS`

- `config.yaml` needs to follow the specific format
    - When you have `node1.example.com:443`, `node2.example.com:443`, `node3.example.com:443` to test:
    ```
    port: 443
    domain: example.com
    subdomain:
    - node1
    - node2
    - node3
    ```

- `ITERATION` is a positive integer denotes how many repeated test to run per Host-ClientHello pair
    - Recommended value: between 1~1000
- `SLEEP_MS` is a positive integer denotes how many millisecond to sleep between 2 adjacent TLS requests