mod error;
mod macros;
mod repository;

pub use error::Error;
pub use repository::Repository;
pub type Result<T> = std::result::Result<T, Error>;
