use core::time;
use sqlx::mysql::MySqlPool;
use std::{env, thread::sleep};
use tokio;

mod db;
use db::user::*;
mod pb;

mod robot;
use robot::Robot;

#[tokio::main]
async fn main() -> Result<(), sqlx::Error> {
    let pool = MySqlPool::connect(&env::var("DATABASE_URL").unwrap()).await?;
    let users = User::get_all_user(&pool, "select * from users where phone_number != '' ").await?;
    println!("user value :{:?}", users.len());
    for user in users {
        println!("ID: {}, Phone Number: {}", user.user_id, user.phone_number);
        let mut robot = Robot::new(123, "123456789");
        robot.init_connect(3);
        break;
    }
    // println!("user value :{:?}", users.len());
    // let user = User::get_user_by_id(&pool, 163978, "14444444444").await?;
    // println!("{:?}", user);
    let mut count = 1;
    loop {
        sleep(time::Duration::from_secs(1));
        count = count + 1;
        if count >= 999999999 {
            break;
        }
    }
    Ok(())
}
