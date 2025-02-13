mod args;
use args::{ImageToAsciiArgs, ImageToAsciiCommand as Command};
use clap::Parser;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    match ImageToAsciiArgs::parse().command {
        Command::Cli(cli) => Ok(println!("{:?}", cli)),
        Command::Server => Ok(println!("Server")),
    }
}
