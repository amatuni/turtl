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

use std::collections::hash_map::Entry;
use std::collections::HashMap;
use std::io::{stdin, stdout, Write};
use std::process;

use termion::event::Key;
use termion::input::TermRead;
use termion::raw::IntoRawMode;
use termion::{clear, color, cursor, style};

use repl;
use shell;

macro_rules! cmdmap {
    ($( $key: expr => $val: expr ),*) => {{
         let mut map: HashMap<&'static str, CommandFunc> = ::std::collections::HashMap::new();
         $( map.insert($key, $val); )*
         map
    }}
}

pub enum Error {
    InvalidCommand,
    CommandExecFail,
}

pub struct Driver {
    cmd_map: HashMap<&'static str, CommandFunc>,
}

impl Driver {
    pub fn new() -> Self {
        let cmd_map = cmdmap!(
        "@help" => help, 
        "@load" => load,
        "@reset" => reset, 
        "@compile" => compile,
        "@run" => run, 
        "@peers" => peers,
        "@save" => save, 
        "@name" => name);

        Driver { cmd_map: cmd_map }
    }

    pub fn validate(&self, input: String, ctx: &mut shell::Context) -> repl::REPLState {
        let args: Vec<&str> = input.trim().split(" ").collect();

        if args.len() > 0 {
            match self.cmd_map.get(args[0]) {
                Some(func) => {
                    func(args, ctx);
                    repl::REPLState::Continue
                }
                None => {
                    invalid_command(args, ctx);
                    repl::REPLState::Continue
                }
            }
        } else {
            repl::REPLState::Continue
        }
    }
}

type CommandFunc = fn(Vec<&str>, &mut shell::Context) -> Result<(), Error>;

pub fn help(args: Vec<&str>, ctx: &mut shell::Context) -> Result<(), Error> {
    println!(
        "{}{}{}",
        cursor::Goto(1, ctx.y as u16),
        clear::CurrentLine,
        HELP_TEXT.replace("\n", "\n\r")
    );
    ctx.y += 17;
    Ok(())
}

const HELP_TEXT: &'static str = r#"
Available Commands:

help [command]      show usage information
quit                quit the shell
load <program>      load a program
compile <program>   generate executable binary from text
run  <program>      run a program
list                list running programs
peers               list connected peers
save [path]         save current shell session to a file
reset               reset the shell to the initial state
name [name]         set the session name

Shell commands should be prefixed with 
an @ symbol, for example: @help
"#;

pub fn reset(args: Vec<&str>, ctx: &mut shell::Context) -> Result<(), Error> {
    ctx.reset();
    Ok(())
}

pub fn load(args: Vec<&str>, ctx: &mut shell::Context) -> Result<(), Error> {
    let path = args[1];
    let mut data: Vec<u8> = match ctx.fs.api.cat(path) {
        Ok(d) => d.collect::<Vec<u8>>(),
        Err(e) => return Err(Error::CommandExecFail),
    };

    ctx.program.truncate(0);
    ctx.program.append(&mut data);
    Ok(())
}

pub fn list(args: Vec<&str>, ctx: &mut shell::Context) -> Result<(), Error> {
    ctx.reset();
    Ok(())
}

pub fn compile(args: Vec<&str>, ctx: &mut shell::Context) -> Result<(), Error> {
    ctx.reset();
    Ok(())
}

pub fn run(args: Vec<&str>, ctx: &mut shell::Context) -> Result<(), Error> {
    ctx.reset();
    Ok(())
}

pub fn peers(args: Vec<&str>, ctx: &mut shell::Context) -> Result<(), Error> {
    ctx.reset();
    Ok(())
}

pub fn save(args: Vec<&str>, ctx: &mut shell::Context) -> Result<(), Error> {
    ctx.reset();
    Ok(())
}

pub fn name(args: Vec<&str>, ctx: &mut shell::Context) -> Result<(), Error> {
    ctx.reset();
    Ok(())
}

pub fn invalid_command(args: Vec<&str>, ctx: &mut shell::Context) -> Result<(), Error> {
    println!(
        "{}{}\n\r\"{}\" is not a valid command\n\r",
        cursor::Goto(1, ctx.y as u16),
        clear::CurrentLine,
        args[0]
    );
    ctx.y += 3;
    Ok(())
}
