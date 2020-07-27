
/**
 * Ron Nathaniel
 * July 26 2020
 */
 
main = (limit = 1000) => {
  let total = 0
  for (let i = 0; i < limit; ++i) {
    if (i % 3 == 0 || i % 5 == 0) {
      total += i;
    }
  }
  return total;
}

const t = main()
console.log(t)
