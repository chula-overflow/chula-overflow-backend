use mongodb::bson::DateTime;
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
pub struct Session {
    pub email: String,
    pub token: String,
    #[serde(alias = "createAt")]
    pub create_at: DateTime,
}

pub struct Revoke {
    pub token: String,
}
