use super::repository::SessionRepository;
use mongodb::bson::{oid::ObjectId, DateTime};

use super::super::user::{
    model::{CreateUser, User},
    repository::UserRepository,
};
use super::model::Session;

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

    pub async fn revoke(&self, token: &str) -> Result<()> {
        let res = self.session_repository.revoke(token).await;

        match res {
            Ok(Some(_)) => Ok(()),
            Ok(None) => Err(Error::NotFound("token not found")),
            Err(e) => Err(e),
        }
    }

    pub async fn me(&self, token: &str) -> Result<User> {
        let user = self.session_repository.me(token).await?;

        Ok(user)
    }

    pub async fn validate(&self, token: &str) -> Result<String> {
        let res = self.session_repository.validate(token).await;

        match res {
            Ok(Some(x)) => Ok(x.token),
            Ok(None) => Err(Error::NotFound("token not found")),
            Err(e) => Err(e),
        }
    }
}
