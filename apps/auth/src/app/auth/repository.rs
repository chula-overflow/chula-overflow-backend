use super::super::repository::Repository;

use super::model::{Revoke, Session};
use crate::error::AuthError;
use mongodb::{bson::doc, results::InsertOneResult};

impl Repository<Session> {
    pub async fn login(&self, login_data: &Session) -> Result<InsertOneResult, AuthError> {
        Ok(self.collection.insert_one(login_data, None).await?)
    }

    pub async fn revoke(&self, revoke_data: &Revoke) -> Result<Option<Session>, AuthError> {
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
