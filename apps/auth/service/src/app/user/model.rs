use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
pub struct User {
    pub email: String,
    pub id: String, // same as object id
}
