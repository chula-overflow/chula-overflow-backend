use super::model::{CreateUser, User};
use mongodb::bson::{doc, oid::ObjectId, DateTime};
use service_core::{Repository, Result};
use tonic::async_trait;

#[async_trait]
pub trait UserRepository {
    async fn find_one_or_create(&self, email: &CreateUser) -> Result<User>;
}

#[async_trait]
impl UserRepository for Repository<User> {
    async fn find_one_or_create(&self, create_user: &CreateUser) -> Result<User> {
        let email = create_user.email.clone();

        let user = self
            .collection
            .find_one(
                doc! {
                    "email": &email
                },
                None,
            )
            .await?;

        match user {
            Some(user) => Ok(user),
            None => {
                let user = User {
                    id: ObjectId::new(),
                    email,
                    create_at: DateTime::now(),
                };

                self.collection.insert_one(&user, None).await?;

                Ok(user)
            }
        }
    }
}
