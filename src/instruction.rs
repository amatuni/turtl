// Copyright 2018 Andrei Amatuni
// This file is part of turtl.
//
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

use std::collections::HashMap;

use num_traits::cast::*;

pub enum Error {
    NotFound,
}

/// The turtl instruction set.
///
///
/// Note: operations which access the filesystem can
/// be **very** slow (e.g. LDFS, STFS, LDFSM, STFSM, etc...)
#[derive(FromPrimitive, ToPrimitive, Copy, Clone)]
pub enum Instruction {
    /// Move from one register (Rb) to another (Ra):
    ///
    /// MOVR Ra Rb
    MOVR,
    /// Move an immediate constant $C into a register (Ra):
    ///
    /// MOVI Ra $C
    MOVI,
    /// Load data from memory (given its address in Rb)
    /// into register Ra:
    ///
    /// LODR Ra [Rb]
    LODR,
    /// Load data from memory (given its address specified by
    /// immediate value %C) into register Ra:
    ///
    /// LODR Ra $C
    LODI,
    /// Store element in register to memory
    /// (given destination address in Rb):
    ///
    /// STRI Ra [Rb]
    STRR,
    /// Store element in register to memory
    /// (given destination address as immediate
    /// constant $C):
    ///
    /// STRI Ra $C
    STRI,
    /// Add values in two registers (Ra + Rb) together,
    /// storing them in a third (Rc):
    ///
    /// ADDR Rc Ra Rb
    ADDR,
    /// Add immediate value $C to value in register Rb,
    /// storing the result in register Ra:
    ///
    /// ADDR Ra Rb $C
    ADDI,
    /// Subtract values in two registers and store
    /// result in a third register:
    ///
    /// SUB Rc Ra Rb
    SUB,
    /// Multiply values in two registers and store
    /// result in a third register:
    ///
    /// MUL Rc Ra Rb
    MUL,
    /// Divide values in two registers and store
    /// result in a third register:
    ///
    /// DIV Rc Ra Rb
    DIV,
    /// Move instruction pointer to a new position
    /// determined by constant $C
    ///
    /// JMPI $C
    JMPI,
    /// Move instruction pointer to a new position
    /// determined by value in register Ra
    ///
    /// JMPI Ra
    JMPR,

    /// Logical equal operation
    EQ,
    /// Logical AND operation
    AND,
    /// Logical less-than operation
    LT,
    /// Logical NOT operation
    NOT,

    /// IPFS operations:
    ///
    /// Load data from filesystem to register.
    ///
    /// LDFS Ra [Rb]
    LDFS,
    /// Load data from filesystem to memory
    LDFSM,
    /// Load data from memory to pubsub channel
    M2PUB,
    /// Load data from pubsub to memory
    PUB2M,
}

/// A map between text representations of opcodes
/// and their binary form (Instruction enum converts to int).
/// This is used to compile human readable text of turtl
/// bytecode to an executable binary representation
pub fn txt_to_instr(op: &str) -> Result<Instruction, Error> {
    match op {
        "MOVR" => Ok(Instruction::MOVR),
        "MOVI" => Ok(Instruction::MOVI),
        "LODR" => Ok(Instruction::LODR),
        "LODI" => Ok(Instruction::LODI),
        "STRR" => Ok(Instruction::STRR),
        "STRI" => Ok(Instruction::STRI),
        "ADDR" => Ok(Instruction::ADDR),
        "ADDI" => Ok(Instruction::ADDI),
        "SUB" => Ok(Instruction::SUB),
        "MUL" => Ok(Instruction::MUL),
        "DIV" => Ok(Instruction::DIV),
        "JMPI" => Ok(Instruction::JMPI),
        "JMPR" => Ok(Instruction::JMPR),
        "EQ" => Ok(Instruction::EQ),
        "AND" => Ok(Instruction::AND),
        "LT" => Ok(Instruction::LT),
        "NOT" => Ok(Instruction::NOT),
        _ => Err(Error::NotFound),
    }
}

/// Given an instruction, return the number
/// of operands that it expects
pub fn argnum(op: Instruction) -> u8 {
    match op {
        Instruction::MOVR => 2,
        Instruction::MOVI => 2,
        Instruction::JMPI => 1,
        Instruction::JMPR => 1,
        Instruction::ADDR => 3,
        Instruction::LODR => 2,
        Instruction::LODI => 2,
        Instruction::STRR => 2,
        Instruction::STRI => 2,
        _ => 2,
    }
}
