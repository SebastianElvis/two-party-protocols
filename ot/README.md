# Oblivious Transfer

An oblivious transfer is a protocol where Alice transfers a secret bit m to to Bob "obviously". That is, 

- Bob knows whether he gets the bit, but Alice doesn’t know whether it was transferred or not.

## Protocols

### The Original Protocol from Rabin [Rab81]

- Alice picks 2 big primes $p, q$, and compute the modulus $N = pq$
- Alice encrypts the message $m$, the ciphertext is $C = E(m, p, q)$ 
  - E() is an arbitrary encryption method that if you know $p, q$ you can decrypt, else you cannot. 
- Alice sends $C, N$ to Bob
- Bob picks $a \in Z_{N}^{*}$ and sends $w = a^{2} \mod N$ to Alice
- Alice computes 4 square roots $x, -x, y, -y$ of $w$, picks one randomly, and sends it back to Bob
- If Bob gets the root which is not $\pm a$, he can factor $N$ and recover $m$, else he cannot.

### The Simplest OT [CO15]

TODO

## References

- [Rab81] M. Rabin. How to exchange secrets by oblivious transfer. Technical Report TR-81, Harvard Aiken Computation Laboratory, 1981. [URL](https://eprint.iacr.org/2005/187.pdf)
- [CO15] Tung Chou and Claudio Orlandi. The Simplest Protocol for Oblivious Transfer. LATINCRYPT'15. [URL](https://eprint.iacr.org/2015/267.pdf)