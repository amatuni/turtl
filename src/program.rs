// Copyright 2018 Andrei Amatuni
//
// This file is part of turtl.

// turtl is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// turtl is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with turtl.  If not, see <http://www.gnu.org/licenses/>.

use std::fs;
use std::io;
use std::path::Path;

use serde;
use serde_cbor;

use tiny_keccak;

/// A Program is a collection of executable instructions packaged
/// with data that might be referenced during execution. Instructions
/// are encoded as 64 bit unsigned integers. Programs are serialized
/// and deserialized to CBOR format. A file with .turtlc extension is
/// just a CBOR serialized Program struct.
///
/// The first 6 bits of each bytecode instruction is the instruction
/// code (allowing for up to 64 possible instructions). This leaves
/// 58 bits for the subsequent operands. The number of operands varies
/// with the type of instruction.
///
/// Operands can be references to a register (i.e. their number 1-12),
/// an immediate value (just basic constants, e.g. 12345), or a memory
/// address (e.g. the byte value at 12345 bytes into the memory buffer)
///
/// Binary encoding of the turtl instruction set:
///
///     [ -------------------- 64 bit integer ---------------------]
///                                 ==
///     [- 6 bit instruction -] [ ---- 58 bits for operands -------]
///
///
///
///
///     INSTRUCTION - [ operand 1 (#bits), operand 2 (#bits), etc..]
///     
///     MOVR - [ register (4), --------------------- register (29) ]
///     MOVI - [ register (4), -------------------- immediate (54) ]
///     JMPI - [ ---------------- immediate (58) ----------------- ]
///     JMPR - [ ---------------- register (58) ------------------ ]
///     LODR - [register (4), ---------------------- register (54) ]
///     LODI - [register (4), --------------------- immediate (54) ]
///     STRR - [register (4), ---------------------- register (54) ]
///     ADDR - [register (19), -- register (19), -- register (20)  ]
///     ADDI - [register (4), --------------------- immediate (54) ]

#[derive(Serialize, Deserialize, Debug)]
pub struct Program {
    /// The ID assigned to this program at run time
    /// by the VM
    #[serde(skip_serializing)]
    pub runtime_id: u64,

    /// The SHA3-256 hash digest of all the bytecode
    pub hash: [u8; 32],

    /// The bytecode defining the program
    pub bytecode: Vec<u64>,
    // pub data: Vec<u8>,
}

impl Program {
    pub fn new() -> Self {
        Program {
            runtime_id: 0,
            hash: [0u8; 32],
            bytecode: Vec::new(),
        }
    }

    pub fn save(&self, path: &Path) -> io::Result<()> {
        let data = serde_cbor::to_vec(self).unwrap();
        fs::write(path, data)?;
        Ok(())
    }

    pub fn check_hash(&self) -> bool {
        // let hash = tiny_keccak::sha3_256(self.bytecode.as_slice());
        // hash == self.hash
        true
    }
}

pub fn load(path: &Path) -> Result<Program, Error> {
    let data = match fs::read(path) {
        Ok(d) => d,
        Err(e) => return Err(Error::NotFound),
    };
    let prog: Program = match serde_cbor::from_slice(data.as_slice()) {
        Ok(p) => p,
        Err(e) => return Err(Error::Malformed),
    };

    if !prog.check_hash() {
        Err(Error::WrongSum)
    } else {
        Ok(prog)
    }
}

pub enum Error {
    NotFound,
    Malformed,
    WrongSum,
}
