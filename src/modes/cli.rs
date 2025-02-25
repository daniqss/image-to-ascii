use super::Mode;
use crate::Ascii;
use clap::builder::ValueParser;
use clap::Parser;
use image::ImageReader;
use std::ffi::OsString;
use std::rc::Rc;

#[derive(Parser, Debug)]
#[command(
    about = "run a local tool that converts images to ascii art",
    arg_required_else_help = true
)]
pub struct CliMode {
    /// path to the input image
    #[arg(index = 1)]
    pub source_path: OsString,
    /// path to the output image
    #[arg(short, long)]
    pub result_path: Option<OsString>,
    /// specify the processing scale
    #[arg(short, long, default_value = "8")]
    pub scale: usize,
    /// density string to create the ascii art
    #[arg(short, long, default_value = " .;coPO#@ ")]
    pub density: String,
    /// print the result to the terminal
    #[arg(short, long)]
    pub print: Option<bool>,
    /// color of the ascii art in hex format, if its not set it will calculate the color from the source image
    #[arg(short, long, value_parser = ValueParser::new(parse_color))]
    pub color: Option<(u8, u8, u8)>,
}

fn parse_color(s: &str) -> Result<(u8, u8, u8), std::num::ParseIntError> {
    let s = s.trim_start_matches('#');
    let r = u8::from_str_radix(&s[0..2], 16)?;
    let g = u8::from_str_radix(&s[2..4], 16)?;
    let b = u8::from_str_radix(&s[4..6], 16)?;
    Ok((r, g, b))
}

impl Mode for CliMode {
    fn run(&self) -> Result<(), Box<dyn std::error::Error>> {
        let img = Rc::new(ImageReader::open(&self.source_path)?.decode()?);

        let ascii = Ascii::new(img, self.scale, self.density.clone(), self.color);
        println!("{:?}", ascii);

        Ok(())
    }
}
