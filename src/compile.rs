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

use instruction;
use instruction::*;
use program::*;
use vm;

pub enum Error {
    InvalidOpcode,
    NotEnoughArgs,
    Syntax,
    InvalidOperand,
    RegisterNumOutOfBounds,
    Overflow,
}

/// Compile the text representation of turtl assembly code
/// into an executable binary
pub fn compile(input: Vec<String>) -> Result<Program, Error> {
    let mut prog = Program::new();
    for mut line in input {
        // get rid of comment
        let code_str = match line.trim().find(";") {
            Some(i) => {
                if i == 0 {
                    continue;
                }
                line.truncate(i);
                line
            }
            None => line,
        };

        // encode into binary opcode
        let code = match encode(code_str) {
            Ok(c) => c,
            Err(e) => return Err(e),
        };
        prog.bytecode.push(code);
    }
    Ok(prog)
}

fn encode(code: String) -> Result<u64, Error> {
    let args: Vec<&str> = code.trim().split(" ").collect();
    let op = match txt_to_instr(args[0]) {
        Ok(c) => c,
        Err(e) => return Err(Error::InvalidOpcode),
    };

    // first 6 bits are the instruction
    let mut result = encode_instruction(op)?;
    // the rest of the bits are instruction dependent
    result |= encode_operands(op, result, args)?;
    println!("after:  {:#066b}\n", result);

    Ok(result)
}

fn encode_instruction(op: Instruction) -> Result<u64, Error> {
    println!("before: {:#066b}", op as u8);
    Ok((op as u64) << 58)
}

fn encode_operands(inst: Instruction, dst: u64, args: Vec<&str>) -> Result<u64, Error> {
    match inst {
        Instruction::MOVR => {
            check_argnum(args.len(), inst)?;
            enc_movr(dst, args.as_slice())
        }
        Instruction::MOVI => {
            check_argnum(args.len(), inst)?;
            enc_movi(dst, args.as_slice())
        }
        Instruction::JMPR => {
            check_argnum(args.len(), inst)?;
            enc_jmpr(dst, args.as_slice())
        }
        Instruction::JMPI => {
            check_argnum(args.len(), inst)?;
            enc_jmpi(dst, args.as_slice())
        }
        Instruction::LODR => {
            check_argnum(args.len(), inst)?;
            enc_lodr(dst, args.as_slice())
        }
        Instruction::LODI => {
            check_argnum(args.len(), inst)?;
            enc_lodi(dst, args.as_slice())
        }
        Instruction::STRR => {
            check_argnum(args.len(), inst)?;
            enc_lodr(dst, args.as_slice())
        }
        Instruction::STRI => {
            check_argnum(args.len(), inst)?;
            enc_lodi(dst, args.as_slice())
        }
        Instruction::ADDR => {
            check_argnum(args.len(), inst)?;
            enc_addr(dst, args.as_slice())
        }

        _ => Err(Error::InvalidOpcode),
    }
}

fn check_argnum(arglen: usize, op: Instruction) -> Result<bool, Error> {
    if (instruction::argnum(op) as usize) != arglen {
        return Ok(true);
    }
    Err(Error::NotEnoughArgs)
}

/// Instruction specific encode functions expect that their inputs
/// have been pre-checked to make sure that their argument count is
/// correct.

#[inline(always)]
fn enc_movr(mut dst: u64, args: &[&str]) -> Result<u64, Error> {
    enc_register_register(dst, args)
}

#[inline(always)]
fn enc_movi(mut dst: u64, args: &[&str]) -> Result<u64, Error> {
    enc_register_immediate(dst, args)
}

#[inline(always)]
fn enc_jmpr(mut dst: u64, args: &[&str]) -> Result<u64, Error> {
    enc_single_register(dst, args)
}

#[inline(always)]
fn enc_jmpi(mut dst: u64, args: &[&str]) -> Result<u64, Error> {
    enc_single_immediate(dst, args)
}

#[inline(always)]
fn enc_lodr(mut dst: u64, args: &[&str]) -> Result<u64, Error> {
    enc_register_register(dst, args)
}

#[inline(always)]
fn enc_lodi(mut dst: u64, args: &[&str]) -> Result<u64, Error> {
    enc_register_immediate(dst, args)
}

fn enc_strr(mut dst: u64, args: &[&str]) -> Result<u64, Error> {
    enc_register_register(dst, args)
}

#[inline(always)]
fn enc_stri(mut dst: u64, args: &[&str]) -> Result<u64, Error> {
    enc_register_immediate(dst, args)
}

#[inline(always)]
fn enc_addr(mut dst: u64, args: &[&str]) -> Result<u64, Error> {
    enc_reg_reg_reg(dst, args)
}

/// Encode a 3 operand register+register+register instruction section
#[inline(always)]
fn enc_reg_reg_reg(mut dst: u64, args: &[&str]) -> Result<u64, Error> {
    let mut reg1 = parse_register_operand(args[1])?;
    let mut reg2 = parse_register_operand(args[2])?;
    let mut reg3 = parse_register_operand(args[3])?;

    reg1 <<= 39;
    reg2 <<= 20;

    dst |= reg1;
    dst |= reg2;
    dst |= reg3;
    Ok(dst)
}

/// Encode a 2 operand register+register instruction section
#[inline(always)]
fn enc_register_register(mut dst: u64, args: &[&str]) -> Result<u64, Error> {
    let mut reg1 = parse_register_operand(args[1])?;
    let mut reg2 = parse_register_operand(args[2])?;
    reg1 <<= 54;
    dst |= reg1;
    dst |= reg2;
    Ok(dst)
}

/// Encode a 2 operand register+immediate instruction section
#[inline(always)]
fn enc_register_immediate(mut dst: u64, args: &[&str]) -> Result<u64, Error> {
    let mut reg1 = parse_register_operand(args[1])?;
    reg1 <<= 54;
    dst |= reg1;

    let mut value = parse_immediate_operand(args[2], 54)?;
    dst |= value;

    Ok(dst)
}

/// Encode a 1 operand single register instruction section
#[inline(always)]
fn enc_single_register(mut dst: u64, args: &[&str]) -> Result<u64, Error> {
    let mut reg1 = parse_register_operand(args[1])?;
    dst |= reg1;
    Ok(dst)
}

/// Encode a 1 operand single immediate value instruction section
#[inline(always)]
fn enc_single_immediate(mut dst: u64, args: &[&str]) -> Result<u64, Error> {
    let mut value = parse_immediate_operand(args[1], 58)?;
    dst |= value;
    Ok(dst)
}

/// Parse an operand that specifies a register
///
/// checks to make sure the syntax is correct
/// and that the register number is within bounds
fn parse_register_operand(arg: &str) -> Result<u64, Error> {
    if !arg.starts_with("%r") {
        return Err(Error::Syntax);
    }

    let mut reg1 = match arg[2..].parse::<u64>() {
        Ok(i) => i,
        Err(e) => return Err(Error::InvalidOperand),
    };

    if (reg1 as usize) > vm::NUM_GEN_REGS {
        return Err(Error::RegisterNumOutOfBounds);
    }

    Ok(reg1)
}

/// Parse an immediate value
///
/// checks to make sure syntax is correct and
/// that the value is within the bounds of the
/// specified bit depth
fn parse_immediate_operand(arg: &str, bit_depth: u64) -> Result<u64, Error> {
    if !arg.starts_with("$") {
        return Err(Error::Syntax);
    }

    let mut value = match arg[1..].parse::<u64>() {
        Ok(i) => i,
        Err(e) => return Err(Error::InvalidOperand),
    };

    if value > bit_depth {
        return Err(Error::Overflow);
    }

    Ok(value)
}

#[cfg(test)]
mod tests {
    #[test]
    pub fn basic_compile() {
        use super::compile;
        let input: Vec<String> = ["MOVI %r1 $13   ; hello", "MOVI %r2 $16", "ADDR %r3 %r1 %r2"]
            .iter()
            .map(|&x| x.into())
            .collect();

        let out = compile(input);
    }
}
