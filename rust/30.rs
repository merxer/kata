fn plus_one(i: i32) -> i32 {
    i + 1
}


fn main() {
    let f = plus_one;
    let six = f(5);
    println!("{}", six);
}
