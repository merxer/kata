fn main() {
    let x = 5;
    match x {
        3|5|6 => { println!("first"); },
        10...16 => { println!("second"); },
        _ => { println!("default");}
    }
}
