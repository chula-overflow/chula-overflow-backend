use service_macros::service;

#[test]
fn simple_test() {
    struct _B;
    struct _C;
    service!(_AuthService use _B, _C);
}
