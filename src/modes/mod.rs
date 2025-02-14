pub trait Mode {
    fn run(&self) -> Result<(), Box<dyn std::error::Error>>;
}

mod cli;
pub use cli::CliMode;
mod server;
pub use server::ServerMode;
