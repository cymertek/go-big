# Go-Big

A numerical library for GoLang forked off the built-in math/big, with extra functionality, and designed to be portable.

## Float Additions:

1. Bytes() and SetBytes() - Ability to work with the mantissa directly as a []byte slice

2. Prune() - The opposite of Int(), in which only the decimal portion of the Float is maintained.
