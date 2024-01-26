use sqlx::mysql::MySqlPool;
//use std::{io::Bytes, str::Bytes};

#[derive(Debug, sqlx::FromRow, sqlx::Type)]
pub struct User {
    pub user_id: i32,
    pub phone_number: String,
}

impl User {
    pub async fn get_all_user(pool: &MySqlPool, sql_query: &str) -> Result<Vec<User>, sqlx::Error> {
        let users = sqlx::query_as::<_, User>(sql_query).fetch_all(pool).await?;
        Ok(users)
    }

    #[allow(dead_code)]
    pub async fn get_user_by_id(
        pool: &MySqlPool,
        uid: i32,
        phone_number: &str,
    ) -> Result<User, sqlx::Error> {
        let query_sql = format!(
            "select * from users where user_id = {} and phone_number = {}",
            uid, phone_number
        );
        let one = sqlx::query_as::<_, User>(&query_sql)
            .fetch_one(pool)
            .await?;
        Ok(one)
    }
}
