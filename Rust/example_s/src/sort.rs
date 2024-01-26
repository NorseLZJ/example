#[derive(Debug, PartialEq, Eq)]
struct Status {
    stat: i32,
}

fn main() {
    let mut vec = vec![
        Status { stat: 1 },
        Status { stat: -1 },
        Status { stat: 2 },
        Status { stat: 9 },
        Status { stat: 0 },
        Status { stat: 8 },
        Status { stat: 1 },
    ];

    vec.sort_by(|a, b| {
        if a.stat == -1 {
            std::cmp::Ordering::Greater
        } else if b.stat == -1 {
            std::cmp::Ordering::Less
        } else {
            a.stat.cmp(&b.stat)
        }
    });

    println!("{:?}", vec);
}
