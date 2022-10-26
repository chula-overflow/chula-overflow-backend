use mongodb::bson::{oid::ObjectId, DateTime};
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
pub struct Session {
    pub _id: ObjectId, // unused
    pub user_id: ObjectId,
    pub token: String,
    pub create_at: DateTime,
}
