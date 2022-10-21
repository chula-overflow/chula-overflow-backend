use serde::Deserialize;

use crate::error::AuthError;
use std::fs::File;

#[derive(Deserialize)]
pub struct Config {
    pub gateway: Endpoint,
    pub auth: Endpoint,
    pub database: DbConfig,
    pub deployment: String,
}

#[derive(Deserialize)]
pub struct Endpoint {
    pub addr: String,
}

#[derive(Deserialize)]
pub struct DbConfig {
    pub addr: String,
    pub db_name: String,
}

pub fn load_config() -> Result<Config, AuthError> {
    let file: File = File::open("../config/config.yaml")?;

    let config: Config = serde_yaml::from_reader(file)?;

    Ok(config)
}
