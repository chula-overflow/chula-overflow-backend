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
    #[error("bson raw error")]
    BSONRawError(#[from] mongodb::bson::raw::Error),
    #[error("uuid parsing error")]
    UUIDParseError(#[from] mongodb::bson::uuid::Error),
    #[error("object id parsing error")]
    OIDParseError(#[from] mongodb::bson::oid::Error),
    #[error("grpc server error")]
    GRpcServerError(#[from] tonic::transport::Error),
    #[error("not found error: {0}")]
    NotFound(&'static str),
    #[error("unknown error: {0}")]
    Unknown(#[from] Box<dyn std::error::Error>),
}
