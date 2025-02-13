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
#[command(
    about = "Run the CLI version of the tool",
    arg_required_else_help = true,
    version
)]
pub struct CliCommand {
    /// path to the input image
    #[arg(short, long)]
    pub source_path: OsString,
    /// path to the output image
    #[arg(short, long)]
    pub result_path: Option<OsString>,
    /// path to the wanted font
    #[arg(long, default_value = "/usr/share/fonts/OpenSans-BoldItalic.ttf")]
    pub font_path: Option<OsString>,
    /// specify the processing scale
    #[arg(short, long, default_value = "8")]
    pub scale: Option<usize>,
    /// print the result to the terminal
    pub print: Option<bool>,
    /// color of the ascii art in hex
    #[arg(short, long)]
    pub color: Option<String>,
}
