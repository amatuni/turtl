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

use std::io::{stdin, stdout, Stdin, Write};

use command::*;
use fs::*;
use version::*;
use vm::*;

use termion::{clear, color, cursor, style};

pub enum ShellError {
    InvalidCommand,
    Syntax,
}

const INIT_X_OFFSET: u16 = 14;
const INIT_Y_OFFSET: u16 = 15;

pub struct Context<'a> {
    /// ID associates to this shell context
    pub id: u64,

    /// name of the shell context
    pub name: String,

    /// history of all past input to the shell
    pub history: Vec<String>,

    /// the current line buffer. this is
    pub line_buffer: String,

    /// the minimum X CLI offset (a function of prompt length)
    pub min_x: u16,

    /// current X value in the CLI
    pub x: u16,

    /// current Y value in the CLI
    pub y: u16,
    // pub index: u64,
    // pub count: i64,
    /// current position relative to self.history
    pub current_pos: usize,

    /// a reference to the IPFS
    pub fs: &'a IPFS,

    pub program: Vec<u8>,
}

impl<'a> Context<'a> {
    pub fn new(fs: &'a IPFS) -> Self {
        Context {
            id: 0,
            name: String::new(),
            history: Vec::new(),
            line_buffer: String::new(),
            min_x: INIT_X_OFFSET,
            x: INIT_X_OFFSET,
            y: 1,
            // index: 0,
            // count: 0,
            current_pos: 0,
            fs: fs,
            program: Vec::new(),
        }
    }

    pub fn reset(&mut self) {
        self.id = 0;
        self.name = String::new();
        self.min_x = INIT_X_OFFSET;
        self.current_pos = 0;
        self.line_buffer.truncate(0);
        self.history.truncate(0);
    }

    pub fn input_char(&mut self, c: char) {
        if self.input_pos() as usize >= self.line_buffer.len() {
            self.line_buffer.push(c);
        } else {
            let idx = self.input_pos() as usize;
            self.line_buffer.insert(idx, c);
        }
        self.x += 1;
    }

    pub fn shift_history_back(&mut self) {
        if self.current_pos > 0 {
            self.current_pos -= 1;
            self.line_buffer.truncate(0);
            self.line_buffer
                .push_str(self.history[self.current_pos].as_ref());
            // self.current_pos -= 1;
            self.reset_x();
            self.x += self.line_buffer.len() as u16;
        }
    }

    pub fn shift_history_forward(&mut self) {
        if self.history.len() > 0 && self.current_pos < self.history.len() - 1 {
            self.current_pos += 1;
            self.line_buffer.truncate(0);
            self.line_buffer
                .push_str(self.history[self.current_pos].as_ref());
            self.reset_x();
            self.x += self.line_buffer.len() as u16;
        }
    }

    pub fn last_input(&self) -> &String {
        &self.history[self.history.len() - 1]
    }

    pub fn reset_x(&mut self) {
        // count number of digits in the
        // index placeholder, so that we
        // know how far to offset x in the
        // prompt
        let mut count = 0;
        let mut n = self.current_pos;
        while n != 0 {
            n /= 10;
            count += 1;
        }
        self.min_x = INIT_X_OFFSET + if self.current_pos > 0 { count - 1 } else { 0 };
        self.x = self.min_x;
    }

    pub fn input_pos(&self) -> u16 {
        self.x - self.min_x
    }
}

pub fn info_graphic() -> String {
    return format!(
        "

    .=.  ____  .=.
    \\ .-'    '-. /       turtl  v{}  - {} bit
    /.'\\_/\\_/'.\\.-p.     {}
--=|: -<_><_>- :|   >
    \\'./ \\/ \\.'/'-b'
    / '-.____.-' \\
    '='        '='
        
    ",
        VERSION,
        INTSIZE * 8,
        &String::from(GIT_HASH)[0..13]
    );
}
