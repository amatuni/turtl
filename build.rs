// Copyright 2018 Andrei Amatuni
// This file is part of turtl.

// tirtl is free software: you can redistribute it and/or modify
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
use std::mem;
use std::process::Command;

fn git_hash() -> String {
    let output = Command::new("git")
        .arg("rev-parse")
        .arg("HEAD")
        .output()
        .expect("failed to execute process");
    let hash_str = String::from_utf8(output.stdout).unwrap();
    return hash_str.replace("\n", "");
}

fn build_version_file(path: &str) {
    let githash = git_hash();
    let semver = env!("CARGO_PKG_VERSION");
    let intsize = mem::size_of::<isize>();
    let file_data = format!(
        "
pub const GIT_HASH: &'static str = \"{}\";
pub const VERSION: &'static str = \"{}\";
pub const INTSIZE: i32 = {};
    ",
        githash, semver, intsize
    );

    fs::write(path, file_data).expect("Unable to write file");
}

fn main() {
    build_version_file("src/version.rs")
}
