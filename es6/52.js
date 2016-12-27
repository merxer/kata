let character = {
  name: 'Bruno',
  pseudonim: 'Batman',
  metadata: {
    age: 34,
    gender: 'male'
  },
  batarang: ['gass pellet', 'bat-mobile']
}

let { pseudonim: alias } = character
console.log(alias)

let { metadata: { gender } } = character
console.log(gender)
