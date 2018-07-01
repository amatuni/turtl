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
#![allow(dead_code)]
#![feature(shrink_to)]
#![allow(unused)]

mod command;
mod compile;
mod fs;
mod instruction;
mod memory;
mod program;
mod repl;
mod shell;
mod version;
mod vm;

use repl::*;
use vm::*;

#[macro_use]
extern crate num_derive;
extern crate num_traits;
#[macro_use]
extern crate serde_derive;
extern crate actix;
extern crate actix_web;
extern crate ctrlc;
extern crate ipfsapi;
extern crate serde;
extern crate serde_cbor;
extern crate termion;
extern crate tiny_keccak;

use std::path;

fn main() {
    let mut vm = TurtlVM::new();
    let mut ipfs = match fs::IPFS::new() {
        Ok(x) => x,
        Err(e) => panic!(e),
    };

    repl(&mut vm, &mut ipfs);
}
