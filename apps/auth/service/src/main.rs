mod proto {
    pub mod user {
        tonic::include_proto!("user");
    }
    pub mod auth {
        tonic::include_proto!("auth");
    }
}

mod app;
mod config;
mod database;

use app::auth::{model::Session, service::AuthService};
use app::user::model::User;
use proto::auth::auth_server::AuthServer;
use service_core::{Repository, Result};
use tonic::transport::Server;

use app::auth::handler::AuthHandler;

#[tokio::main]
async fn main() -> Result<()> {
    let conf = config::load_config()?;

    let addr = format!("0.0.0.0:{}", conf.auth.port).parse().unwrap();

    // setup db
    let db = database::get_db_conn(&conf).await?;
    let auth_collection = db.collection::<Session>("session");
    let user_collection = db.collection::<User>("user");

    let auth_repo = Repository::new(auth_collection);
    let user_repo = Repository::new(user_collection);

    let auth_service = AuthService::new(auth_repo, user_repo);

    let auth_srv_hdr = AuthHandler::new(auth_service);

    println!("Listening on {}", addr);

    Server::builder()
        .add_service(AuthServer::new(auth_srv_hdr))
        .serve(addr)
        .await?;

    Ok(())
}
