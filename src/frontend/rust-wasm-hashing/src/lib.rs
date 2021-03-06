extern crate sha2;
extern crate wasm_bindgen;

use std::cell::Cell;
use std::string::String;

use sha2::Digest;
use sha2::Sha256;
use wasm_bindgen::prelude::*;

// When the `wee_alloc` feature is enabled, use `wee_alloc` as the global
// allocator.
#[cfg(feature = "wee_alloc")]
#[global_allocator]
static ALLOC: wee_alloc::WeeAlloc = wee_alloc::WeeAlloc::INIT;

#[wasm_bindgen]
pub struct Sha256hasher {
    hasher: Cell<Sha256>,
}

#[wasm_bindgen]
impl Sha256hasher {
    #[wasm_bindgen(constructor)]
    pub fn new() -> Sha256hasher {
        Sha256hasher {
            hasher: Cell::new(Sha256::default()),
        }
    }

    pub fn update(&mut self, input_bytes: &[u8]) {
        let hasher = self.hasher.get_mut();
        hasher.input(input_bytes)
    }

    pub fn hex_digest(&mut self) -> String {
        let hasher = self.hasher.take();
        let output = hasher.result();
        self.hasher = Cell::new(Sha256::default());
        return format!("{:x}", output);
    }
}
