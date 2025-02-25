mod args;
mod ascii;
mod modes;
use args::{ImageToAsciiArgs, ImageToAsciiCommand as Command};
use ascii::Ascii;
use clap::Parser;
use modes::*;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    match ImageToAsciiArgs::parse().command {
        Command::Cli(cli) => cli.run(),
        Command::Server(server) => server.run(),
    }
}
