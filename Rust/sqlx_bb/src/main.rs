use sqlx::mysql::MySqlPool;
use std::env;
use tokio;

mod db;
use db::user::*;

#[tokio::main]
async fn main() -> Result<(), sqlx::Error> {
    let pool = MySqlPool::connect(&env::var("DATABASE_URL").unwrap()).await?;

    let users = User::get_all_user(&pool, "select * from users where phone_number != '' ").await?;
    println!("user value :{:?}", users.len());
    let users = User::get_all_user(&pool, "select * from users").await?;
    // for user in users {
    //     println!("ID: {}, Phone Number: {}", user.user_id, user.phone_number);
    // }
    println!("user value :{:?}", users.len());
    let user = User::get_user_by_id(&pool, 163978, "14444444444").await?;
    println!("{:?}", user);
    Ok(())
}
