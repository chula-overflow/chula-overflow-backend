use crate::proto::{auth_server::Auth, AuthRequest, AuthResponse, RevokeRequest, RevokeResponse};

use mongodb::bson::DateTime;
use tonic::{async_trait, Request, Response, Status};

use super::model::{Revoke, Session};

use super::super::service::Service;

#[async_trait]
impl Auth for Service<Session> {
    async fn login(&self, request: Request<AuthRequest>) -> Result<Response<AuthResponse>, Status> {
        let email = request.get_ref().email.to_owned();
        let token = uuid::Uuid::new_v4().to_string();
        let create_at = DateTime::now();

        let session = Session {
            email,
            token,
            create_at,
        };

        let result = self.repository.login(&session).await;

        let token = session.token;

        match result {
            Ok(_) => Ok(Response::new(AuthResponse { token })),
            Err(e) => {
                eprintln!("Err: {}", e);
                Err(Status::internal(e.to_string()))
            }
        }
    }

    async fn revoke(
        &self,
        request: Request<RevokeRequest>,
    ) -> Result<Response<RevokeResponse>, Status> {
        let token = request.get_ref().token.to_owned();

        let result = self.repository.revoke(&Revoke { token }).await;

        match result {
            Ok(Some(_)) => Ok(Response::new(RevokeResponse {})),
            Ok(None) => Err(Status::not_found("Token not found")),
            Err(e) => {
                eprintln!("Err: {}", e);
                Err(Status::internal(e.to_string()))
            }
        }
    }
}
