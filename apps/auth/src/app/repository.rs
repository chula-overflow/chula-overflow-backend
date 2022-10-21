use mongodb::Collection;

pub struct Repository<T> {
    pub collection: Collection<T>,
}

impl<T> Repository<T> {
    pub fn new(collection: Collection<T>) -> Repository<T> {
        Repository { collection }
    }
}
