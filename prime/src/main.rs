use num::Float;

static COUNT: i64 = 100000000;

fn main() {
    for i in 1..COUNT {
        let result: bool = is_prime(i);
        if result {
            println!("{} is prime", i);
        } else {
            println!("{} is not prime", i);
        }
    }
}

pub fn is_prime(num: i64) -> bool {
    if num == 1 {
        return false;
    }
    if num == 2 {
        return true;
    }
    if num % 2 == 0 {
        return false;
    }
    let root: i64 = Float::sqrt(num as f64) as i64;
    for i in (3..root + 1).step_by(2) {
        if num % i == 0 {
            return false;
        }
    }
    return true;
}
