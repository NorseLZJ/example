import pymysql
import random
import concurrent.futures


def insert_data(uid):
    conn = pymysql.connect(
        host="127.0.0.1",
        user="root",
        password="123456",
        db="test",
    )
    cursor = conn.cursor()

    username = f"User{uid}"
    address = f"Address{uid}"
    iphone = f"{random.randint(10000000000, 99999999999)}"
    age = random.randint(18, 60)
    salary = round(random.uniform(1000, 10000), 2)
    job = f"Job{uid % 10}"
    level = uid % 5

    sql = f"INSERT INTO User (Uid, UserName, Address, Iphone, Age, Salary, Job, Level) VALUES ({uid}, '{username}', '{address}', '{iphone}', {age}, {salary}, '{job}', {level})"

    cursor.execute(sql)
    conn.commit()
    cursor.close()
    conn.close()


if __name__ == "__main__":
    with concurrent.futures.ThreadPoolExecutor(max_workers=8) as executor:
        executor.map(insert_data, range(1, 1000001))
