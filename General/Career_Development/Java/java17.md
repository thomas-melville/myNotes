# java 17

## Sealed Classes

No longer preview

no changes from java 16

## Internal changes

### Floating point semantics

simplify the jvms handling of struct float point numbers.

### macOS rendering pipeline

use the Apple Metal API as an alternative to the existing

## Ports

macOS/Arch64 port

macOS is transitioning from x64 to AArch64

## deprecations and removals

deprecate Applet API for removal

remove RMI activation

remove experimental AOT and JIT Compiler

deprecate Security Manager for removal

## Foreign Function & Memory API

incubator

evolution of two existing incubation API

replace the JNI interface

## Vector API

express vector communications

## Context specific deserializtion filters

configure context-specific and dynamically selected deserialization filters.
protection against deserializtion of untrusted data
On ObjectInputStreams

## Strong encapsulation

no longer possible to relax it, --illegal-access no longer has an affect

## Pattern matching for switch

preview feature

```

switch (o) {
  case Integer i -> ...
  case Long l -> ...
  ...
}

```

null can also be a case

can combine in case to check parts of object

```

switch (o) {
  case Integer i && ( i < 100) -> ..
}
```

## new apis

### Random number generators

new interface RandomGenerator
  new uniform API for all types of Random Generators

four specialized interfaces
  SplittableRandomGenerator
    spawn new RG from existing
  JumpableRandomGenerator
    jump ahead a moderate number of draws
  LeapableRandomGenerator
    jump ahead a large number of draws
  ArbitrarilyJumpableRandomGenerator
    can specify arbitrary jump distance
