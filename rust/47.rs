#[allow(unused_variables)]
#[allow(unused_assignments)]
fn main() {
    let mut x: i32 = 1;
    x = 7;
    let x = x;
    println!("{}", x);

    let y = 4;
    let y = "I can also be bound to text!";
    println!("{}", y);
}
