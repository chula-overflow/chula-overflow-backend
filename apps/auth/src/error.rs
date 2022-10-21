use std::net::AddrParseError;

use thiserror::Error;

#[derive(Error, Debug)]
pub enum AuthError {
    #[error("Database error: {0}")]
    DbError(#[from] mongodb::error::Error),
    #[error("IO error")]
    IOError(#[from] std::io::Error),
    #[error("Address parsing error")]
    AddrParseError(#[from] AddrParseError),
    #[error("YAML parsing error")]
    YAMLParseError(#[from] serde_yaml::Error),
    #[error("UUID parsing error")]
    UUIDParseError(#[from] mongodb::bson::uuid::Error),
    #[error("GRpc server error")]
    GRpcServerError(#[from] tonic::transport::Error),
    #[error("Unknown error")]
    Unknown(#[from] Box<dyn std::error::Error>),
}
