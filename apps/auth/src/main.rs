mod proto {
    tonic::include_proto!("auth");
}

mod app;
mod config;
mod database;
mod error;

use std::net::SocketAddr;

use app::auth::model::Session;
use app::{repository::Repository, service::Service};
use error::AuthError;
use proto::auth_server::AuthServer;
use tonic::transport::Server;
#[allow(dead_code, unused_variables)]
#[tokio::main]
async fn main() -> Result<(), AuthError> {
    let conf = config::load_config()?;

    // setup db
    let db = database::get_db_conn(&conf).await?;

    let collection = db.collection::<Session>("session");

    // config server
    let addr: SocketAddr = conf.auth.addr.parse()?;

    let auth_repo = Repository::new(collection);
    let auth_service = Service::new(auth_repo);

    println!("Listening on {}", addr);

    Server::builder()
        .add_service(AuthServer::new(auth_service))
        .serve(addr)
        .await?;

    Ok(())
}
