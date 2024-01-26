use std::time::Duration;
use tokio::time::interval;

#[tokio::main]
async fn main() {
    let ticker1 = interval(Duration::from_secs(1));
    let ticker2 = interval(Duration::from_secs(2));
    let ticker3 = interval(Duration::from_secs(3));

    let task1 = tokio::spawn(hello_task(ticker1, "hello 1"));
    let task2 = tokio::spawn(hello_task(ticker2, "hello 2"));
    let task3 = tokio::spawn(hello_task(ticker3, "hello 3"));

    tokio::try_join!(task1, task2, task3).unwrap();
}

async fn hello_task(mut ticker: tokio::time::Interval, message: &'static str) {
    loop {
        tokio::select! {
            _ = ticker.tick() => {
                println!("{}", message);
            }
            _ = tokio::time::sleep(Duration::from_millis(100)) => {
                // Do other tasks if needed
            }
        }
    }
}
