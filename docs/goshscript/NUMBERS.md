# Numbers

Goshell does not differentiate between integers and floating point numbers.

Numbers can be written as decimal, binary (`0b` or `0B`), hexadecimal (`0x`, `0X`) numbers.
It is even possible to specify numbers with any radix (even floating point radices) using the form `0rnn:xxx`:

 - `15` becomes 15.0
 - `3.14` becomes 3.14
 - `0b1001` becomes 9.0
 - `0b11.1` becomes 3.5
 - `0r3:100` becomes 9.0
 - `0r1.5:11` becomes 2.5
