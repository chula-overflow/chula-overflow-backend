use std::net::AddrParseError;

#[derive(thiserror::Error, Debug)]
pub enum Error {
    #[error("database error: {0}")]
    DbError(#[from] mongodb::error::Error),
    #[error("io error")]
    IOError(#[from] std::io::Error),
    #[error("address parsing error")]
    AddrParseError(#[from] AddrParseError),
    #[error("var error")]
    EnvError(#[from] std::env::VarError),
    #[error("uuid parsing error")]
    UUIDParseError(#[from] mongodb::bson::uuid::Error),
    #[error("grpc server error")]
    GRpcServerError(#[from] tonic::transport::Error),
    #[error("unknown error")]
    Unknown(#[from] Box<dyn std::error::Error>),
}
