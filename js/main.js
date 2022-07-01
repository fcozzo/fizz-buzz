for (let i = 1; i <= 100; i++) {
  console.log(fizzBuzz(i))
}

function fizzBuzz(number) {
  const isFizz = (number % 3) == 0 
  const isBuzz = (number % 5) == 0

  let result = ''

  if (isFizz) {
    result = 'Fizz'
  } 

  if (isBuzz) {
    result = result ? result + ' Buzz' : 'Buzz'
  } 

  return result ? result : number
}