use jsonwebtoken::{decode, encode, DecodingKey, EncodingKey, Header, Validation};
use serde::{Deserialize, Serialize};
use std::time::{Duration, SystemTime};

/// Our claims struct, it needs to derive `Serialize` and/or `Deserialize`
#[derive(Debug, Serialize, Deserialize)]
struct Claims {
    user_name: String,
    phone: String,
    user_id: u64,
    exp: u64,
}

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let my_claims = Claims {
        user_name: String::from("NorseLZJ"),
        phone: String::from("1233"),
        user_id: 989888,
        exp: one_day_sec(),
    };

    let secret = "5Vs3gtRo44kHWWnnTOAx";
    let token = encode(
        &Header::default(),
        &my_claims,
        &EncodingKey::from_secret(secret.as_ref()),
    )
    .unwrap_or_default();
    println!("{}", token);

    // `token` is a struct with 2 fields: `header` and `claims` where `claims` is your own struct.
    let token = decode::<Claims>(
        &token,
        &DecodingKey::from_secret(secret.as_ref()),
        &Validation::default(),
    )?;
    println!("{:?}", token);
    Ok(())
}

fn one_day_sec() -> u64 {
    let now = SystemTime::now();
    let one_day = Duration::from_secs(24 * 60 * 60); // 24 hours in seconds
    let one_day_later = now + one_day;
    let one_day_later_secs = match one_day_later.duration_since(SystemTime::UNIX_EPOCH) {
        Ok(duration) => duration.as_secs(),
        Err(_) => panic!("SystemTime before UNIX EPOCH!"),
    };
    one_day_later_secs
}
