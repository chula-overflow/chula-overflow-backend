use mongodb::bson::RawDocument;
use service_core::{Error, Repository, Result};
use tonic::async_trait;

use super::super::user::model::User;
use super::model::{Revoke, Session};
use mongodb::{bson::doc, results::InsertOneResult};

#[async_trait]
pub trait SessionRepository {
    async fn login(&self, login_data: &Session) -> Result<InsertOneResult>;
    async fn revoke(&self, revoke_data: &Revoke) -> Result<Option<Session>>;
    async fn me(&self, token: &str) -> Result<User>;
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
    async fn me(&self, token: &str) -> Result<User> {
        let query = vec![
            doc! {
                "$match": doc! {
                    "token": token
                }
            },
            doc! {
                "$lookup": doc! {
                    "from": "user",
                    "localField": "user_id",
                    "foreignField": "_id",
                    "as": "user"
                }
            },
        ];

        // driver bug
        let mut cursor = self.collection.aggregate(query, None).await?;
        let aggregated: &RawDocument;

        if cursor.advance().await? {
            aggregated = cursor.current();
        } else {
            return Err(Error::NotFound("token not found"));
        }

        // functional time
        let user = aggregated
            .get_array("user")
            .map_err(|_| "'user' not found")
            .and_then(|user_arr| {
                user_arr
                    .get(0)
                    // not sure which one is real not found
                    .map_err(|_| "user not found")
                    .and_then(|x| x.ok_or("user not found"))
            })
            .and_then(|raw_bson| {
                raw_bson
                    .as_document()
                    .ok_or("type error: cannot convert to raw document")
            })
            .and_then(|raw_doc| {
                let email = raw_doc
                    .get_str("email")
                    .map_err(|_| "key error: cannot find 'email' in document");
                let create_at = raw_doc
                    .get_datetime("create_at")
                    .map_err(|_| "key error: cannot find 'create_at' in document");
                let object_id = raw_doc
                    .get_object_id("_id")
                    .map_err(|_| "key error: cannot find '_id' in document");

                if let (Ok(email), Ok(create_at), Ok(object_id)) = (email, create_at, object_id) {
                    Ok(User {
                        email: email.to_string(),
                        create_at,
                        id: object_id,
                    })
                } else {
                    Err("key error")
                }
            });

        if let Err(x) = user {
            if x == "user not found" {
                return Err(Error::NotFound("user not found"));
            } else {
                return Err(Error::Unknown(Box::from(x)));
            }
        }

        Ok(user.unwrap())
    }
}
