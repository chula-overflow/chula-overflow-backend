use crate::config::Config;
use mongodb::{options::ClientOptions, Client, Database};
use service_core::Result;

pub async fn get_db_conn(config: &Config) -> Result<Database> {
    let client_options: ClientOptions =
        ClientOptions::parse(format!("mongodb://{}", config.database.addr)).await?;

    let client = Client::with_options(client_options)?;

    let database = client.database(&config.database.db_name);

    Ok(database)
}
