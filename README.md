# turtl

[![Build Status (Travis)](https://travis-ci.org/andreiamatuni/turtl.svg?branch=master)](https://travis-ci.org/andreiamatuni/turtl)


A p2p register based virtual machine.


### install

```
$ git clone https://github.com/andreiamatuni/turtl.git
$ cd turtl
$ cargo install
```

### start

```
$ turtl shell
```


### FAQ

- Is this a programming language?
    - No (...well, kind of). This is the thing that sits underneath a programming language. turtl is a bytecode interpreter. You can write your own programming language to target the turtl instruction set and deploy programs to the turtl network.
-  What about security?
    - YOLO!!!!
- ...but actually...what about security?
    - The VM is sandboxed from the host environment. It doesn't have access to host CPU/memory/filesystem/IO
- How does IO work then?
    - All input and output is handled through IPFS file objects and pubsub. turtl has dedicated instructions to read and write files to IPFS and publish/subscribe to channels on IPFS pubsub.
- Are there blockchains involved?
    - No.
- Are there any programming languages that target turtl at present?
    - Not that I know of. If you're comfortable with writing assembly code, you can write programs directly in turtl assembly and compile/run/deploy them. More info on the instruction set [here](src/program.rs) and [here](src/instruction.rs). You can find examples of human readable assembly files (ending in .turtl) [here](test/turtl_code). These will be compiled into executable .turtlc files with the @compile command (in the turtl shell).
- What is the point of this? Why not just hook a wasm interpreter to IPFS?
    - see this doc describing the [motivation](docs/motivation.md)
- Is there a pithy (somewhat obnoxious) one liner that encapsulates the ethos of the project?
    - ```it's turtls all the way down``` 
- Does this VM execute code super fast?
    - again...the name of the project is turtl...