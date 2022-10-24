use service_core::Error as SrvErr;
use tonic::{async_trait, Request, Response, Status};

use crate::proto::auth::{
    auth_server::Auth, AuthRequest, AuthResponse, MeRequest, MeResponse, RevokeRequest,
    RevokeResponse,
};

use super::service::AuthService;

pub struct AuthHandler {
    pub service: AuthService,
}

#[async_trait]
impl Auth for AuthHandler {
    async fn login(&self, request: Request<AuthRequest>) -> Result<Response<AuthResponse>, Status> {
        let email = &request.get_ref().email.to_owned();

        log::debug!("[Login] email: {}", email);

        let result = self.service.login(email).await;

        match result {
            Ok(token) => Ok(Response::new(AuthResponse { token })),
            Err(e) => {
                log::error!("[Login] {}", e);
                Err(Status::internal(e.to_string()))
            }
        }
    }

    async fn revoke(
        &self,
        request: Request<RevokeRequest>,
    ) -> Result<Response<RevokeResponse>, Status> {
        let token = &request.get_ref().token;

        log::debug!("[Revoke] token: {}", token);

        let result = self.service.revoke(token).await;

        match result {
            Ok(_) => Ok(Response::new(RevokeResponse {})),
            Err(SrvErr::NotFound(x)) => {
                log::info!("[Revoke] request not found has occured {}", token);
                Err(Status::not_found(x))
            }
            Err(e) => {
                log::error!("[Revoke] {}", e);
                Err(Status::internal(e.to_string()))
            }
        }
    }

    async fn me(&self, request: Request<MeRequest>) -> Result<Response<MeResponse>, Status> {
        let token = &request.get_ref().token;

        log::debug!("[Me] token: {}", token);

        let result = self.service.me(token).await;

        match result {
            Ok(user) => Ok(Response::new(MeResponse {
                user: Some(user.into()),
            })),
            Err(SrvErr::NotFound(x)) => {
                log::info!("[Me] request not found has occured {}", token);
                Err(Status::not_found(x))
            }
            Err(e) => {
                log::error!("[Me] {}", e);
                Err(Status::internal(e.to_string()))
            }
        }
    }
}

impl AuthHandler {
    pub fn new(service: AuthService) -> Self {
        Self { service }
    }
}
