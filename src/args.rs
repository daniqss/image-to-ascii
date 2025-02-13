use clap::{Parser, Subcommand};
use std::ffi::OsString;

#[derive(Parser, Debug)]
#[command(
    about = "Tool to create ascii art images from other images",
    arg_required_else_help = true,
    version
)]
pub struct ImageToAsciiArgs {
    #[command(subcommand)]
    pub command: ImageToAsciiCommand,
}

#[derive(Subcommand, Debug)]
pub enum ImageToAsciiCommand {
    Cli(CliCommand),
    Server,
}

#[derive(Parser, Debug)]
#[command(about = "Run the CLI version of the tool", version)]
pub struct CliCommand {
    pub path: Option<OsString>,
    pub font_path: Option<OsString>,
    pub scale: Option<usize>,
    pub print: Option<bool>,
    pub colored: Option<bool>,
}
