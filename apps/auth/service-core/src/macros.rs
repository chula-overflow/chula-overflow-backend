#[macro_export]
macro_rules! throw_if_err {
    ($x: expr) => {
        match $x {
            Ok(v) => v,
            Err(e) => {
                eprintln!("Err: {}", e);
                return Err(Status::internal(e.to_string()));
            }
        }
    };
}
