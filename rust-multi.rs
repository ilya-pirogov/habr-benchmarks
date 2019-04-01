use clap::clap_app;
use rayon::prelude::*;
use std::sync::atomic::{AtomicIsize, Ordering};

fn main() {
    let matches = clap_app!(divs => (@arg maxN: +required)).get_matches();
    let n = matches.value_of("maxN").unwrap();
    let n: i64 = n.parse().unwrap();
    let arr: Vec<_> = (0..=n).map(|_| AtomicIsize::new(0)).collect();

    let n_sqrt = (n as f64).sqrt() as i64;
    (1..n_sqrt + 1).into_par_iter().for_each(|k1| {
        (k1 .. 1 + n / k1).into_par_iter().for_each(|k2| {
            let val = if k1 != k2 { k1 + k2 } else { k1 };
            arr[(k1 * k2) as usize].fetch_add(val as isize, Ordering::Relaxed);
        });
    });
    println!("{}", arr.last().unwrap().load(Ordering::SeqCst));
}
