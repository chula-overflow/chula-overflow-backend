use super::repository::Repository;

pub struct Service<T> {
    pub repository: Repository<T>,
}

impl<T> Service<T> {
    pub fn new(repository: Repository<T>) -> Service<T> {
        Service { repository }
    }
}
