# GFW-2021Summer-TLS-Attack

Thoughts, discussions, and clarifications on the attack against TLS possibly initiated by GFW in 2021 Summer (June-July). 

This is an ongoing incident, thus all information provided here is subjected to change.

## Latest Status

07/22/2021 0:33 UTC
- Still cannot reproduce the issue. According to 3rd-party source, the reset happens on only AS4134. AS4809 isn't impacted. This guess matches our data.

07/22/2021 18:36 UTC
- This attack is no longer active according to my (only) test node. You are welcomed to provide any information about this incident.

## Background

Starting in 2021 June, the Great Firewall, operated by the Chinese government, started blocking and interfering with the widely used TLS proxies in China. Despite those which have got their port 443 blocked from Chinese users, some target proxy servers started to experience SSL/TLS handshake failure caused by RST injection. Meanwhile, the HTTPS website is **still accessible**.

## Initial guesses

- GFW can analyze the request sent from TLS proxy clients to the server.
  - False. The RST happens during the TLS handshake where no actual proxy-like behavior has started.
- GFW checks against TLS Fingerprints exposed in ClientHello and selectively attacks against infamous TLS fingerprints or allow only clients with famous TLS fingerprints to pass.
  - False. Forging the TLS fingerprint from a well-known browser does not improve the connectivity.

## Data (n=50)

<img src="https://raw.githubusercontent.com/Gaukas/GFW-2021Summer-TLS-Proxy-Attack/master/data/Stat.png">

- Note: Node name may not reflect the actual routing and/or server physical location. 
- **Latency (ms):** 65, 69, 93, 191, 63, 199, 131, 133
- **Packet loss:** 8%, 12%, 2%, 0%, 0%, 0%, 0%, 0%

<img src="https://raw.githubusercontent.com/Gaukas/GFW-2021Summer-TLS-Proxy-Attack/master/data/SNI_verification.png">

## Rough conclusion

- GFW **does not show the ability to identify TLS proxy traffics** from normal HTTPS web browsing traffics **in real-time**.
- GFW **isn't utilizing a purely fingerprint-based discrimination**. We don't know if fingerprint matters in current state but it is not a major decision factor.
- Statistics gives **a strong signal about a backend IP reputation system**. This may indicate that GFW is analyzing the TLS traffic and there exists difference between proxy traffic and real web browsing traffic.
- Statistics gives **no sign about SNI sniffing** in this attack.
- (Needs confirmation) RST happens only on AS4134, while AS4809 isn't impacted. (Note: Our only good server, `us-cn2gia-1` is the only one routing over AS4809.)

## Credits 

- [@ewust](https://github.com/ewust), [@jhalderm](https://github.com/jhalderm), [@jmwample](https://github.com/jmwample)
- [Refraction Networking/uTLS](https://github.com/refraction-networking/utls)
- [TLSfingerprint.io](https://tlsfingerprint.io/)
