use super::repository::SessionRepository;
use mongodb::bson::{oid::ObjectId, DateTime};

use super::super::user::{
    model::{CreateUser, User},
    repository::UserRepository,
};
use super::model::{Revoke, Session};

use service_macros::service;

use service_core::{Error, Result};

service!(pub AuthService use Session, User);

impl AuthService {
    pub async fn login(&self, email: &String) -> Result<String> {
        let token = uuid::Uuid::new_v4().to_string();

        let user = self
            .user_repository
            .find_one_or_create(&CreateUser {
                email: email.clone(),
            })
            .await?;

        let session = Session {
            _id: ObjectId::new(),
            user_id: user.id,
            token,
            create_at: DateTime::now(),
        };

        let _ = self.session_repository.login(&session).await?;

        Ok(session.token)
    }

    pub async fn revoke(&self, token: &String) -> Result<()> {
        let result = self
            .session_repository
            .revoke(&Revoke {
                token: token.clone(),
            })
            .await;

        match result {
            Ok(Some(_)) => Ok(()),
            Ok(None) => Err(Error::NotFound("token not found")),
            Err(e) => Err(e),
        }
    }

    pub async fn me(&self, token: &String) -> Result<User> {
        let user = self.session_repository.me(token).await?;

        Ok(user)
    }
}
