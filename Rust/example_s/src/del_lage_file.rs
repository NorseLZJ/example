use rayon::prelude::*;
use std::env;
use std::fs;
use std::path::Path;
use std::time::{Duration, Instant, SystemTime};

fn is_large_file(entry: &fs::DirEntry) -> bool {
    if let Ok(metadata) = entry.metadata() {
        metadata.is_file() && metadata.len() > 1_000_000_000 // 1GB
    } else {
        false
    }
}

fn process_directory(directory_path: &Path, current_time: SystemTime) {
    if let Ok(entries) = fs::read_dir(directory_path) {
        for entry in entries {
            if let Ok(entry) = entry {
                let path = entry.path();

                if path.is_dir() {
                    process_directory(&path, current_time);
                }

                if let Ok(modified_time) = entry.metadata().map(|m| m.modified()) {
                    if let Ok(modified_time) = modified_time {
                        let duration = current_time
                            .duration_since(modified_time)
                            .unwrap_or(Duration::from_secs(0));

                        if duration.as_secs() > 10 * 24 * 60 * 60 && is_large_file(&entry) {
                            println!("{:?} is a large file in an old directory.", path);
                        }
                    }
                }
            }
        }
    }
}

fn main() {
    let args: Vec<String> = env::args().collect();
    let parallel = args.len() > 1 && args[1] == "1";

    let dirs = vec![
        r"C:\Program Files\",
        r"C:\Program Files (x86)\",
        r"C:\ProgramData\",
        r"C:\Users\Administrator\",
    ];
    let current_time = SystemTime::now();
    let start_time = Instant::now();

    if parallel {
        dirs.par_iter()
            .for_each(|d| process_directory(Path::new(&d), current_time));
    } else {
        for d in dirs.iter() {
            process_directory(Path::new(d), current_time);
        }
    }
    let elapsed_time = start_time.elapsed();
    println!("Total execution time: {:?}", elapsed_time);
}
