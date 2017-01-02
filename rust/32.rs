use std::mem;

fn main() {
    let a: u8 = 123;
    println!("a = {}", a);
    
    let mut b: i8 = 0; //mutable
    println!("b = {}", b);
    b = 42;
    println!("b = {}", b);

    let mut c = 123456789;
    println!("c = {}, size = {}", c,  mem::size_of_val(&c));
    c = -1; 
    println!("c = {}, size = {}", c,  mem::size_of_val(&c));

    let z: isize = 123; // isize/ usize
    let size_of_z = mem::size_of_val(&z);
    println!("z = {}, takes up {} bytes, {}-bit OS",
            z, size_of_z, size_of_z * 8);

    let d: char = 'x';
    println!("d = {}, size = {} bytes", d, mem::size_of_val(&d));

    let e: f64 = 2.5;
    println!("e = {}, size = {} bytes", e, mem::size_of_val(&e));

    let g = false;
    println!("g = {}, size = {} bytes", g, mem::size_of_val(&g));

    let f: bool = 4 > 0;
    println!("f = {}, size = {} bytes", f, mem::size_of_val(&f));

    println!("{}", i8::min_value());
}
