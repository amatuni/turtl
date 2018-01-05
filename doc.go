/*
Package turtl implements an embedable register based virtual machine,
and an accompanying programming language.

The turtl VM is written to target platforms with a word size of at least 64 bits.
While this may seem limiting, it's easier to optimize for a single platform
and it's 2018 folks...my cell phone has multiple 64 bit cores.
*/
package turtl
