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

use command;
use fs;
use shell;
use vm;

use std::borrow::Borrow;
use std::io::{stdin, stdout, Write};

use termion::event::Key;
use termion::input::TermRead;
use termion::raw::IntoRawMode;
use termion::{clear, color, cursor, style};

pub enum REPLState {
    Continue,
    Quit,
}

pub fn repl(vm: &mut vm::TurtlVM, fs: &mut fs::IPFS) {
    let stdin = stdin();
    let mut stdout = stdout().into_raw_mode().unwrap();

    let mut ctx = shell::Context::new(fs);
    let mut command_driver = command::Driver::new();

    print_welcome(&mut ctx, &mut stdout);
    prompt(&mut ctx, &mut stdout);

    for c in stdin.keys() {
        match c.unwrap() {
            Key::Char(c) => {
                if c as u64 == 10 {
                    reposition_cursor(&ctx, &mut stdout);
                    newline(&mut ctx, &mut stdout);

                    // only process if there has been input
                    if ctx.line_buffer.len() > 0 {
                        match process_entry(&mut ctx, &command_driver, &mut stdout) {
                            REPLState::Quit => {
                                quit(&ctx, &mut stdout);
                                break;
                            }
                            REPLState::Continue => {}
                        }
                    }

                    prompt(&mut ctx, &mut stdout);
                } else {
                    reposition_cursor(&ctx, &mut stdout);
                    ctx.input_char(c);
                    write_line(&mut ctx, &mut stdout);
                    reposition_cursor(&ctx, &mut stdout);
                }
            }
            Key::Alt(c) => {}
            Key::Ctrl(c) => {
                if c == 'c' {
                    quit(&ctx, &mut stdout);
                    break;
                }
            }
            Key::Esc => {}
            Key::Left => {
                if ctx.x > ctx.min_x {
                    ctx.x -= 1;
                    reposition_cursor(&ctx, &mut stdout);
                }
            }
            Key::Right => {
                if ctx.input_pos() < ctx.line_buffer.len() as u16 {
                    ctx.x += 1;
                    reposition_cursor(&ctx, &mut stdout);
                }
            }
            Key::Up => {
                ctx.shift_history_back();
                write_line(&mut ctx, &mut stdout);
            }
            Key::Down => {
                ctx.shift_history_forward();
                write_line(&mut ctx, &mut stdout);
            }
            Key::Backspace => {
                if ctx.x > ctx.min_x {
                    ctx.x -= 1;
                    let idx = ctx.input_pos() as usize;
                    ctx.line_buffer.remove(idx);
                    write_line(&mut ctx, &mut stdout)
                }
            }
            _ => {}
        }
        stdout.flush().unwrap();
    }
}

fn process_entry(ctx: &mut shell::Context, cmd: &command::Driver, out: &mut Write) -> REPLState {
    ctx.history.push(ctx.line_buffer.clone());
    ctx.current_pos = ctx.history.len();
    let repl_state = eval(ctx, cmd);
    ctx.reset_x();
    ctx.line_buffer.truncate(0);
    repl_state
}

fn eval(ctx: &mut shell::Context, cmd: &command::Driver) -> REPLState {
    if ctx.last_input().starts_with("@quit") {
        return REPLState::Quit;
    } else if ctx.last_input().starts_with("@") {
        cmd.validate(ctx.line_buffer.clone(), ctx);
        println!(
            "\n\r{:?}\n\r",
            String::from_utf8(ctx.program.clone()).unwrap()
        )
    }
    REPLState::Continue
}

fn prompt(ctx: &mut shell::Context, out: &mut Write) {
    write!(
        out,
        "{}{}[{}{}{}] {}{}turtl{} :> ",
        cursor::Goto(1, ctx.y as u16),
        color::Fg(color::Red),
        style::Reset,
        ctx.current_pos,
        color::Fg(color::Red),
        style::Bold,
        color::Fg(color::Green),
        style::Reset,
    ).unwrap();
    out.flush().unwrap();
}

fn print_welcome(ctx: &mut shell::Context, out: &mut Write) {
    write!(
        out,
        "\n\r{}{}{}\n\r{}",
        clear::All,
        cursor::Goto(1, 0),
        shell::info_graphic().replace("\n", "\n\r"),
        "Welcome to turtl! Type @help to see available commands.\n\r"
    ).unwrap();
    ctx.y += 14
}

fn reposition_cursor(ctx: &shell::Context, out: &mut Write) {
    write!(out, "{}", cursor::Goto(ctx.x, ctx.y)).unwrap();
}

fn write_line(ctx: &mut shell::Context, out: &mut Write) {
    write!(out, "{}", clear::CurrentLine).unwrap();
    prompt(ctx, out);
    write!(out, "{}{}", cursor::Goto(ctx.min_x, ctx.y), ctx.line_buffer).unwrap();
}

fn quit(ctx: &shell::Context, out: &mut Write) {
    write!(
        out,
        "{}{}{}{}bye bye :)\n\r\n\r",
        cursor::Goto(1, ctx.y as u16 + 1),
        clear::CurrentLine,
        cursor::Goto(1, ctx.y as u16 + 2),
        clear::CurrentLine,
    );
    ctx.fs.shutdown();
}

fn newline(ctx: &mut shell::Context, out: &mut Write) {
    write!(out, "{}\n\r", cursor::Goto(ctx.x, ctx.y)).unwrap();
    ctx.y += 1;
}
