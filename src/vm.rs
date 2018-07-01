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

use program::*;

pub const NUM_GEN_REGS: usize = 12;

pub struct TurtlVM {
    /// General purpose VM registers
    register: [u64; NUM_GEN_REGS],
    /// Instruction pointer
    ip: usize,
    /// flags register
    flags: u32,
    /// Currently loaded program
    program: Program,
}

impl TurtlVM {
    pub fn new() -> Self {
        TurtlVM {
            register: [0u64; 12],
            ip: 0,
            flags: 0b1111_0000_1010_1111_1010_1111_0000_1111,
            program: Program::new(),
        }
    }

    pub fn load_program(prog: Program) {}

    /// run the VM forward by one cycle
    pub fn tick(&mut self) {
        let op = self.program.bytecode[self.ip];
    }
}
