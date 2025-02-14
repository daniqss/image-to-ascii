use crate::modes::{CliMode, ServerMode};
use clap::{Parser, Subcommand};

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
    Cli(CliMode),
    Server(ServerMode),
}
