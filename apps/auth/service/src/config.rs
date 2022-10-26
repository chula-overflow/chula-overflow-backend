use serde::Deserialize;

use service_core::Result;

use std::env::var;

#[derive(Deserialize, Debug)]
pub struct Config {
    pub auth: Endpoint,
    pub database: DbConfig,
    pub deployment: String,
}

#[derive(Deserialize, Debug)]
pub struct Endpoint {
    pub host: String,
    pub port: String,
}

#[derive(Deserialize, Debug)]
pub struct DbConfig {
    pub uri: String,
    pub db_name: String,
}

pub fn load_config() -> Result<Config> {
    let config = Config {
        auth: Endpoint {
            host: var("AUTH_HOST")?,
            port: var("AUTH_PORT")?,
        },
        database: DbConfig {
            uri: var("MONGODB_URI")?,
            db_name: var("MONGODB_DB_NAME")?,
        },
        deployment: var("DEPLOYMENT")?,
    };

    log::debug!("loaded config: {:?}", config);

    Ok(config)
}
