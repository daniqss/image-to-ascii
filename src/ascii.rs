use std::rc::Rc;

use image::DynamicImage;

pub struct Ascii {
    img: Rc<DynamicImage>,
    scale: usize,
    density: String,
    color: Option<(u8, u8, u8)>,
}

impl Ascii {
    pub fn new(
        img: Rc<DynamicImage>,
        scale: usize,
        density: String,
        color: Option<(u8, u8, u8)>,
    ) -> Self {
        Self {
            img: img.clone(),
            scale,
            density,
            color,
        }
    }
}

impl std::fmt::Debug for Ascii {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        f.debug_struct("Ascii")
            .field("img", &self.img)
            .field("scale", &self.scale)
            .field("density", &self.density)
            .field("color", &self.color)
            .finish()
    }
}
