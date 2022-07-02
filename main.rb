require 'pry'

COUNT = 10

def is_prime_regex(num)
  chars = "x" * num
  return /^(?!(xx+)\1+$|^x$)/.match?(chars)
end

def is_prime_ruby(num)
  return false if num == 1
  return true if num == 2 || num == 3
  return false if num % 2 == 0

  root = Math.sqrt(num).floor
  (3..(root + 1)).step(2).each do |p|
    if num % p == 0
      return false
    end
  end

  return true
end

(1..COUNT).each do |i|
  is_prime = is_prime_regex(i)
  message = is_prime ? "#{i} is prime\n" : "#{i} is not prime\n"
  print(message)
end
