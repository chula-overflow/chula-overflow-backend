use serde::Deserialize;

use service_core::Result;

use std::env::var;

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

pub fn load_config() -> Result<Config> {
    let config = Config {
        gateway: Endpoint {
            addr: var("GATEWAY_URL")?,
        },
        auth: Endpoint {
            addr: var("AUTH_URL")?,
        },
        database: DbConfig {
            addr: var("MONGODB_URL")?,
            db_name: var("MONGODB_DB_NAME")?,
        },
        deployment: var("DEPLOYMENT")?,
    };

    Ok(config)
}
