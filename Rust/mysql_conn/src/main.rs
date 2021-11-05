use mysql::prelude::*;
use mysql::*;

#[derive(Debug, PartialEq, Eq)]
struct Payment {
    customer_id: i32,
    amount: i32,
    account_name: Option<String>,
}

impl Payment {
    fn new(params: (i32, i32, Option<String>)) -> Payment {
        Payment {
            customer_id: params.0,
            amount: params.1,
            account_name: params.2,
        }
    }
}

fn main() {
    println!("mysql conn!");
    let url = "mysql://root:123456@127.0.0.1:3308/rust";
    let opts = Opts::from_url(url).unwrap();
    //println!("opts{:#?}", opts);
    let pool = Pool::new(opts).unwrap();

    let mut conn = pool.get_conn().unwrap();

    // Let's create a table for payments.
    //TEMPORARY 临时表
    conn.query_drop(
        r"CREATE TEMPORARY TABLE payment (
        customer_id int not null,
        amount int not null,
        account_name text
    )",
    )
    .unwrap();

    let payments = vec![
        Payment::new((1, 2, None)),
        Payment::new((3, 4, Some("foo".into()))),
        Payment::new((5, 6, None)),
        Payment::new((6, 8, None)),
        Payment::new((9, 10, Some("bar".into()))),
    ];

    conn.exec_batch(
        r"INSERT INTO payment (customer_id, amount, account_name)
      VALUES (:customer_id, :amount, :account_name)",
        payments.iter().map(|p| {
            println!("p{:#?}", &p);
            params! {
                "customer_id" => p.customer_id,
                "amount" => p.amount,
                "account_name" => &p.account_name,
            }
        }),
    )
    .unwrap();

    let selected_payments = conn
        .query_map(
            "SELECT customer_id, amount, account_name from payment",
            |(customer_id, amount, account_name)| Payment {
                customer_id,
                amount,
                account_name,
            },
        )
        .unwrap();

    assert_eq!(payments, selected_payments);
    println!("Yay!");
}
