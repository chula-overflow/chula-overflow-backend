use crate::proto::user::User as ProtoUser;
use mongodb::bson::{oid::ObjectId, DateTime};
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
pub struct User {
    #[serde(rename = "_id")]
    pub id: ObjectId,
    pub email: String,
    pub create_at: DateTime,
}

pub struct CreateUser {
    pub email: String,
}

impl Into<ProtoUser> for User {
    fn into(self) -> ProtoUser {
        ProtoUser { email: self.email }
    }
}
