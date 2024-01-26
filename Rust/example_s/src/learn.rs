fn main() {
    //let maybe_value: Option<&str> = Some("abc");
    let maybe_value: Option<&str> = None;

    // 使用 .and_then() 链式调用
    let v = maybe_value
        .and_then(|s| s.parse::<i32>().ok())
        .unwrap_or_default();

    println!("{}", v);

    let maybe_value: Option<i32> = Some(999);
    let v = maybe_value.map(|x| x.to_string()).unwrap_or_default();
    println!("{}", v);
    let v = maybe_value
        .and_then(|x| Some(x.to_string()))
        .unwrap_or_default();
    println!("{}", v);
}
