use super::Mode;
use clap::Parser;

#[derive(Parser, Debug)]
#[command(
    about = "run a http server that accepts images and returns the ascii art version of the image",
    arg_required_else_help = true
)]
pub struct ServerMode {
    /// port to run the server on
    #[arg(short, long, default_value = "3000")]
    pub port: Option<u16>,
}

impl Mode for ServerMode {
    fn run(&self) -> Result<(), Box<dyn std::error::Error>> {
        Ok(println!("Server mode"))
    }
}
