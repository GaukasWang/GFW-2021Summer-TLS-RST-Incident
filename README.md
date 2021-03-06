# GFW-2021Summer-TLS-RST-Incident

Thoughts, discussions, and clarifications on the abnormal RST packet injection against TLS connections, which is possibly initiated by GFW, in 2021 Summer (June-July). 

The incident is now concluded. Therefore, no more status update until the behavior is observed again. 

You are welcomed to provide your insights and data, if any.

## Latest Status

11/08/2021 05:30 UTC
- The behavior is again observed on some(?) servers. We believe it is related to **China Communist Party's Sixth Plenary Session of the 19th Central Committee** (Nov 8-11)

07/28/2021 16:00 UTC
- The behavior has been inactive for several days. We are concluding the incident. 

07/23/2021 16:30 UTC
- The behavior is back on only 1 server I have control over. 
- A minor timeout issue is seen across all servers. Not AS4134-specific.
- Server-side ALPN-support test is conducted -- no sign.

07/23/2021 0:33 UTC
- Still cannot reproduce the issue. According to 3rd-party source, the reset happens on only AS4134. AS4809 isn't impacted. This guess matches our data.

07/22/2021 18:36 UTC
- This attack is no longer active according to my (only) test node. You are welcomed to provide any information about this incident.

## Background

Starting in 2021 June, the Great Firewall, operated by the Chinese government, started blocking and interfering with the widely used TLS proxies in China. Despite those which have got their port 443 blocked from Chinese users, some target proxy servers started to experience SSL/TLS handshake failure caused by RST injection. Meanwhile, the HTTPS website is **still accessible**.

## Initial guesses

- The attacker can analyze the request sent from TLS proxy clients to the server.
  - False. The RST happens during the TLS handshake where no actual proxy-like behavior has started.
- The attacker checks against TLS Fingerprints exposed in ClientHello and selectively attacks against infamous TLS fingerprints or allow only clients with famous TLS fingerprints to pass.
  - False. Forging the TLS fingerprint from a well-known browser does not improve the connectivity.

## Data (n=50)

<img src="https://raw.githubusercontent.com/Gaukas/GFW-2021Summer-TLS-Proxy-Attack/master/data/Stat.png">

- Note: Node name may not reflect the actual routing and/or server physical location. 
- **Latency (ms):** 65, 69, 93, 191, 63, 199, 131, 133
- **Packet loss:** 8%, 12%, 2%, 0%, 0%, 0%, 0%, 0%

<img src="https://raw.githubusercontent.com/Gaukas/GFW-2021Summer-TLS-Proxy-Attack/master/data/Stat0723.png">

07/23/2021: After about 1 days without seeing the handshake being reset, the attack is back on 1 server. Additional minor timeout problem appears. 

<img src="https://raw.githubusercontent.com/Gaukas/GFW-2021Summer-TLS-Proxy-Attack/master/data/SNI_verification.png">

<img src="https://raw.githubusercontent.com/Gaukas/GFW-2021Summer-TLS-Proxy-Attack/master/data/ALPN_verification.png">

## Rough conclusion

- The attacker **does not show the ability to identify TLS proxy traffics** from normal HTTPS web browsing traffics **in real-time**.
- The attacker **isn't utilizing a purely fingerprint-based discrimination**. We don't know if fingerprint matters in current state but it is not a major decision factor.
- Statistics gives **a strong signal about a backend IP reputation system**. This may indicate that the attacker is analyzing history TLS traffic model.
- Statistics gives **no sign about SNI sniffing** nor **server-side ALPN discrimination** in this attack.
- (Needs confirmation) RST happens only on AS4134, while AS4809 isn't impacted. (Note: Our only good server, `us-cn2gia-1` is the only one guaranteeing the route over AS4809.)

## Credits 

- [@ewust](https://github.com/ewust), [@jhalderm](https://github.com/jhalderm), [@jmwample](https://github.com/jmwample)
- [Refraction Networking/uTLS](https://github.com/refraction-networking/utls)
- [TLSfingerprint.io](https://tlsfingerprint.io/)
