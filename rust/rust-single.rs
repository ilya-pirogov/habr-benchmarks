use clap::clap_app;

fn main() {
    let matches = clap_app!(divs => (@arg maxN: +required)).get_matches();
    let n = matches.value_of("maxN").unwrap();
    let n: i64 = n.parse().unwrap();
    let mut arr = vec![0; (n+1) as usize];
    for k1 in 1 ..= (n as f64).sqrt() as i64 {
        for k2 in k1 ..= n / k1 {
            let val = if k1 != k2 { k1 + k2 } else { k1 };
            arr[(k1*k2) as usize] += val;
        }
    }
    println!("{}", arr.last().unwrap());
}
