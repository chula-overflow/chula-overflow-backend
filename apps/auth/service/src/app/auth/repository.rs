use service_core::{Repository, Result};
use tonic::async_trait;

use super::model::{Revoke, Session};
use mongodb::{bson::doc, results::InsertOneResult};

#[async_trait]
pub trait SessionRepository {
    async fn login(&self, login_data: &Session) -> Result<InsertOneResult>;
    async fn revoke(&self, revoke_data: &Revoke) -> Result<Option<Session>>;
}

#[async_trait]
impl SessionRepository for Repository<Session> {
    async fn login(&self, login_data: &Session) -> Result<InsertOneResult> {
        Ok(self.collection.insert_one(login_data, None).await?)
    }

    async fn revoke(&self, revoke_data: &Revoke) -> Result<Option<Session>> {
        let token = &revoke_data.token;

        Ok(self
            .collection
            .find_one_and_delete(
                doc! {
                    "token": token
                },
                None,
            )
            .await?)
    }
}
