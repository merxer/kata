fn main() {
    let x = 4;
    match x {
        0 => { println!("unmatch"); },
        4 => { println!("match"); },
        _ => { println!("default");}
    }
}
