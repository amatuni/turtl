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

use actix_web::client;
use ipfsapi;
use std::process::{Child, Command, Stdio};

pub enum Error {
    SubprocFail,
}

pub struct IPFS {
    /// The handle to the IPFS daemon child process
    child: Child,

    /// the HTTP api
    pub api: ipfsapi::IpfsApi,
}

impl IPFS {
    pub fn new() -> Result<Self, Error> {
        let child = match Command::new("ipfs")
            .arg("daemon")
            .stdout(Stdio::piped())
            .spawn()
        {
            Ok(c) => c,
            Err(e) => return Err(Error::SubprocFail),
        };
        let api = ipfsapi::IpfsApi::new("127.0.0.1", 5001);
        Ok(IPFS {
            child: child,
            api: api,
        })
    }

    pub fn kill_process(&mut self) {
        match self.child.kill() {
            Ok(()) => {}
            Err(e) => {}
        }
    }

    pub fn shutdown(&self) {
        match self.api.shutdown() {
            Ok(x) => {}
            Err(e) => {}
        }
    }
}
