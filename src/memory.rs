// Copyright 2018 Andrei Amatuni
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

/// A global memory store with random read and write access
pub struct Memory {
    buffer: Vec<u8>,
}

impl Memory {
    /// Create a new memory buffer given a size in bytes
    pub fn new(size: usize) -> Self {
        Memory {
            buffer: vec![0; size],
        }
    }

    /// Write bytes to the buffer at a given position
    pub fn write(&mut self, data: &[u8], position: usize) {
        if position + data.len() < self.buffer.len() {
            self.buffer[position..position + data.len()].clone_from_slice(data);
        }
    }

    // pub fn read(&self, start: usize, end: usize)

    // pub fn read()
}
